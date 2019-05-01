package bootstrap

import (
	"context"
	"database/sql"
	"net"
	"time"

	"github.com/cachecashproject/go-cachecash/bootstrap/models"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/ed25519"
	"google.golang.org/grpc/peer"
)

type Bootstrapd struct {
	l  *logrus.Logger
	db *sql.DB
}

func NewBootstrapd(l *logrus.Logger, db *sql.DB) (*Bootstrapd, error) {
	return &Bootstrapd{
		l:  l,
		db: db,
	}, nil
}

func (b *Bootstrapd) HandleCacheAnnounceRequest(ctx context.Context, req *ccmsg.CacheAnnounceRequest) (*ccmsg.CacheAnnounceResponse, error) {
	startupTime := time.Unix(req.StartupTime, 0)

	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("failed to get grpc peer from ctx")
	}

	var srcIP net.IP
	switch addr := peer.Addr.(type) {
	case *net.UDPAddr:
		srcIP = addr.IP
	case *net.TCPAddr:
		srcIP = addr.IP
	}

	// TODO: last_ping column

	cache := models.Cache{
		PublicKey:   ed25519.PublicKey(req.PublicKey),
		Version:     req.Version,
		FreeMemory:  req.FreeMemory,
		TotalMemory: req.TotalMemory,
		FreeDisk:    req.FreeDisk,
		TotalDisk:   req.TotalDisk,
		StartupTime: startupTime,
		ExternalIP:  srcIP,
		ContactURL:  req.ContactUrl,
		LastPing:    time.Now(),
	}

	// TODO: figure out how to do proper upserts
	/*
		err := cache.Upsert(ctx, b.db, true, []string{"public_key"}, boil.Infer(), boil.Infer())
		if err != nil {
			return nil, errors.Wrap(err, "failed to add cache to database")
		}
	*/

	// XXX: ignore duplicate key errors
	cache.Insert(ctx, b.db, boil.Infer())
	// XXX: force an update in case the insert failed due to a conflict
	cache.Update(ctx, b.db, boil.Infer())

	b.reapStaleAnnoucements(ctx)

	return &ccmsg.CacheAnnounceResponse{}, nil
}

func (b *Bootstrapd) reapStaleAnnoucements(ctx context.Context) error {
	deadline := time.Now().Add(-5 * time.Minute)
	rows, err := models.Caches(qm.Where("last_ping<?", deadline)).DeleteAll(ctx, b.db)
	if err != nil {
		return err
	}
	b.l.Debugf("Removed %d stale caches from database", rows)
	return nil
}

func (b *Bootstrapd) HandleCacheFetchRequest(ctx context.Context, req *ccmsg.CacheFetchRequest) (*ccmsg.CacheFetchResponse, error) {
	b.reapStaleAnnoucements(ctx)

	caches, err := models.Caches().All(ctx, b.db)
	if err != nil {
		return nil, err
	}

	resp := &ccmsg.CacheFetchResponse{}
	for _, c := range caches {
		resp.Caches = append(resp.Caches, &ccmsg.CacheDescription{
			PublicKey:   c.PublicKey,
			Version:     c.Version,
			FreeMemory:  c.FreeMemory,
			TotalMemory: c.TotalMemory,
			FreeDisk:    c.FreeDisk,
			TotalDisk:   c.TotalDisk,
			StartupTime: c.StartupTime.Unix(),
			ContactUrl:  c.ContactURL,
			ExternalIp:  c.ExternalIP.String(),
		})
	}

	return resp, nil
}