package client

import (
	"context"
	"encoding/base64"

	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/cachecashproject/go-cachecash/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"golang.org/x/crypto/ed25519"
	"google.golang.org/grpc"
)

// - Assigns sequence numbers to outbound messages.
// - Routes replies by matching sequence numbers.
// - How do we handle the consumer of a reply exiting/terminating/canceling?
type cacheGrpc struct {
	l      *logrus.Logger
	pubkey []byte

	// working state --
	// used to discriminate between chunk requests
	nextSequenceNo uint64
	// In flight chunk requests for this cache
	backlogInt uint64
	// Client side assessment of the cache - we don't depend solely on the GRPC
	// connectivity state because we need to cope with running but broken caches
	status ccmsg.ContentRequest_ClientCacheStatus_Status

	// The GRPC Connection and API client
	conn       *grpc.ClientConn
	grpcClient ccmsg.ClientCacheClient
	// l2 ticket exchanges are completed
	l2Done  <-chan bool
	l2Queue chan<- l2payment
}

type cacheConnection interface {
	SubmitRequest(context.Context, chan DownloadResult, *chunkRequest)
	ExchangeTicketL2(context.Context, *ccmsg.ClientCacheRequest)
	Close(context.Context) error
	GetStatus() ccmsg.ContentRequest_ClientCacheStatus
	PublicKey() string
	PublicKeyBytes() []byte
}

type l2payment struct {
	span *trace.Span
	req  *ccmsg.ClientCacheRequest
}

type DownloadResult struct {
	resp  *chunkRequest
	cache cacheConnection
}

var _ cacheConnection = (*cacheGrpc)(nil)

func newCacheConnection(l *logrus.Logger, addr string, pubkey ed25519.PublicKey) (*cacheGrpc, error) {
	// Cache security is currently delivered solely via the cachecash crypto
	// guarantees. TLS certs will be added in future.
	conn, err := common.GRPCDial(addr, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial")
	}

	grpcClient := ccmsg.NewClientCacheClient(conn)
	l2Done := make(chan bool)
	l2Queue := make(chan l2payment)
	go sendL2Payments(l2Done, l2Queue, grpcClient, conn, l)

	return &cacheGrpc{
		l:      l,
		pubkey: pubkey,

		nextSequenceNo: 4000, // XXX: Make this easier to pick out of logs.

		conn:       conn,
		grpcClient: grpcClient,
		l2Done:     l2Done,
		l2Queue:    l2Queue,
	}, nil
}

func (cc *cacheGrpc) Close(ctx context.Context) error {
	cc.l.WithField("cache", cc.PublicKey()).Info("cacheGrpc.Close() - enter")
	// A nil request to terminate the goroutine
	cc.l2Queue <- l2payment{}
	select {
	case <-cc.l2Done:
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "cache connection close cancelled")
	}
	if err := cc.conn.Close(); err != nil {
		return errors.Wrap(err, "failed to close connection")
	}
	return nil
}

func (cc *cacheGrpc) run(ctx context.Context, clientNotify chan DownloadResult, chunkRequest *chunkRequest) {
	defer func() { cc.backlogInt-- }()
	l := cc.l.WithFields(logrus.Fields{
		"cache": cc.PublicKey(),
	})
	l.WithFields(logrus.Fields{
		"backlog": cc.backlogInt,
	}).Debug("got download request")
	err := cc.requestChunk(ctx, chunkRequest)
	chunkRequest.err = err
	l.Debug("yielding download result")
	clientNotify <- DownloadResult{
		resp:  chunkRequest,
		cache: cc,
	}
}

func (cc *cacheGrpc) SubmitRequest(ctx context.Context, clientNotify chan DownloadResult, req *chunkRequest) {
	cc.backlogInt++
	go cc.run(ctx, clientNotify, req)
}

func (cc *cacheGrpc) ExchangeTicketL2(ctx context.Context, req *ccmsg.ClientCacheRequest) {
	cc.l2Queue <- l2payment{span: trace.FromContext(ctx), req: req}
}

func (cc *cacheGrpc) requestChunk(ctx context.Context, b *chunkRequest) error {
	ctx = trace.NewContext(ctx, b.parent)
	ctx, span := trace.StartSpan(ctx, "cachecash.com/Client/requestChunk")
	defer span.End()
	// Send request ticket to cache; await data.
	reqData, err := b.bundle.BuildClientCacheRequest(b.bundle.TicketRequest[b.idx])
	if err != nil {
		return errors.Wrap(err, "failed to build client-cache request")
	}
	msgData, err := cc.grpcClient.GetChunk(ctx, reqData)
	if err != nil {
		// TODO - make this transient - reset after a period of time
		cc.status = ccmsg.ContentRequest_ClientCacheStatus_UNUSABLE
		return errors.Wrap(err, "failed to exchange request ticket with cache")
	}
	cc.l.WithFields(logrus.Fields{
		"cache":    cc.PublicKey(),
		"chunkIdx": b.bundle.TicketRequest[b.idx].ChunkIdx,
		"len":      len(msgData.Data),
	}).Info("got data response from cache")

	// Send L1 ticket to cache; await outer decryption key.
	reqL1, err := b.bundle.BuildClientCacheRequest(b.bundle.TicketL1[b.idx])
	if err != nil {
		return errors.Wrap(err, "failed to build client-cache request")
	}
	msgL1, err := cc.grpcClient.ExchangeTicketL1(ctx, reqL1)
	if err != nil {
		cc.status = ccmsg.ContentRequest_ClientCacheStatus_UNUSABLE
		return errors.Wrap(err, "failed to exchange request ticket with cache")
	}
	cc.l.WithField("cache", cc.PublicKey()).Info("got L1 response from cache")

	// Decrypt data.
	encData, err := util.EncryptChunk(
		b.bundle.TicketRequest[b.idx].ChunkIdx,
		b.bundle.Remainder.RequestSequenceNo,
		msgL1.OuterKey.Key,
		msgData.Data)
	if err != nil {
		cc.status = ccmsg.ContentRequest_ClientCacheStatus_UNUSABLE
		return errors.Wrap(err, "failed to decrypt data")
	}
	b.encData = encData

	// Done!
	return nil
}

func (cc *cacheGrpc) GetStatus() ccmsg.ContentRequest_ClientCacheStatus {
	return ccmsg.ContentRequest_ClientCacheStatus{
		BacklogDepth: cc.backlogInt,
		Status:       cc.status,
	}
}

func (cc *cacheGrpc) PublicKey() string {
	return base64.StdEncoding.EncodeToString(cc.pubkey)
}

func (cc *cacheGrpc) PublicKeyBytes() []byte {
	return cc.pubkey
}

// send l2 payments to a cache asynchronously; terminates on a nil payment object
func sendL2Payments(l2Done chan bool, l2Queue chan l2payment, grpcClient ccmsg.ClientCacheClient, conn *grpc.ClientConn, l *logrus.Logger) {
	for payment := range l2Queue {
		if payment.req == nil {
			l2Done <- true
			return
		}
		func(ctx context.Context, payment l2payment) {
			ctx = trace.NewContext(ctx, payment.span)
			_, err := grpcClient.ExchangeTicketL2(ctx, payment.req)
			if err != nil {
				l.Errorf("Failed to send L2 payment to cache %s", err)
			}
		}(context.Background(), payment)
	}
}
