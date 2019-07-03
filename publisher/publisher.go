package publisher

import (
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/cachecashproject/go-cachecash/batchsignature"
	"github.com/cachecashproject/go-cachecash/catalog"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/cachecashproject/go-cachecash/publisher/models"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/ed25519"
)

type ContentPublisher struct {
	// The ContentPublisher knows each cache's "inner master key" (aka "master key")?  This is an AES key.
	// For each cache, it also knows an IP address, a port number, and a public key.

	l  *logrus.Logger
	db *sql.DB

	signer  ed25519.PrivateKey
	catalog catalog.ContentCatalog

	escrows []*Escrow
	caches  map[string]*ParticipatingCache

	// XXX: It's obviously not great that this is necessary.
	// Maps object IDs to metadata; necessary to allow the publisher to generate cache-miss responses.
	reverseMapping map[common.ObjectID]reverseMappingEntry

	// XXX: Need cachecash.PublicKey to be an array of bytes, not a slice of bytes, or else we can't use it as a map key
	// caches map[cachecash.PublicKey]*ParticipatingCache

	PublisherAddr string
}

type CacheInfo struct {
	// ...
}

type reverseMappingEntry struct {
	path string
}

func NewContentPublisher(l *logrus.Logger, db *sql.DB, publisherAddr string, catalog catalog.ContentCatalog, signer ed25519.PrivateKey) (*ContentPublisher, error) {
	p := &ContentPublisher{
		l:              l,
		db:             db,
		signer:         signer,
		catalog:        catalog,
		caches:         make(map[string]*ParticipatingCache),
		reverseMapping: make(map[common.ObjectID]reverseMappingEntry),
		PublisherAddr:  publisherAddr,
	}

	return p, nil
}

func (p *ContentPublisher) LoadFromDatabase(ctx context.Context) (int, error) {
	escrows, err := models.Escrows().All(ctx, p.db)
	if err != nil {
		return 0, errors.Wrap(err, "failed to query Escrows")
	}

	for _, e := range escrows {
		escrow := &Escrow{
			Inner:  *e,
			Caches: []*ParticipatingCache{},
		}

		ecs, err := e.EscrowCaches().All(ctx, p.db)
		if err != nil {
			return 0, errors.Wrap(err, "failed to query EscrowsCaches")
		}

		for _, ec := range ecs {
			c, err := ec.Cache().One(ctx, p.db)
			if err != nil {
				return 0, errors.Wrap(err, "failed to query Cache")
			}

			escrow.Caches = append(escrow.Caches, &ParticipatingCache{
				Cache:          *c,
				InnerMasterKey: ec.InnerMasterKey,
			})
		}

		err = p.AddEscrow(escrow)
		if err != nil {
			return 0, errors.Wrap(err, "failed to query Cache")
		}
	}

	return len(escrows), nil
}

// XXX: Temporary
func (p *ContentPublisher) AddEscrow(escrow *Escrow) error {
	p.escrows = append(p.escrows, escrow)

	// setup a map from pubkey -> *cache
	for _, cache := range escrow.Caches {
		p.caches[string(cache.PublicKey())] = cache
	}

	return nil
}

// XXX: Temporary
func (p *ContentPublisher) getEscrowByRequest(req *ccmsg.ContentRequest) (*Escrow, error) {
	if len(p.escrows) == 0 {
		return nil, errors.New("no escrow for request")
	}
	return p.escrows[0], nil
}

/*
The process of satisfying a request

  - A request arrives for an object, identified by a _path_, which is actually an opaque series of bytes.  (In our
    implementation, they're HTTP-like paths.)  It (optionally) includes a _byte range_, which may be open-ended (which
    means "continue until the end of the object").

  - The _byte range_ is translated to a _chunk range_ depending on how the publisher would like to split the object.
    (Right now, we only support fixed-size chunks, but this is not inherent.)  The publisher may also choose how many
    chunks it would like to serve, and how many chunk-groups they will be divided into.  (The following steps are
    repeated for each chunk-group; the results are returned together in a single response to the client.)

  - The object's _path_ is used to ensure that the object exists, and that the specified chunks are in-cache and valid.
    (This may be satisfied by the content catalog's cache, or may require contacting an upstream.)  (A future
    enhancement might require that the publisher fetch only the cipher-blocks that will be used in puzzle generation,
    instead of all of the cipher-blocks in the chunks.)

  - The _path_ and _chunk range_ are mapped to a list of _chunk identifiers_.  These are arbitrarily assigned by the
    publisher.  (Our implementation uses the chunk's digest.)

  - The publisher selects a single escrow that will be used to service the request.

  - The publisher selects a set of caches that are enrolled in that escrow.  This selection should be designed to place
    the same chunks on the same caches (expanding the number in rotation as demand for the chunks grows), and to reuse
    the same caches for consecutive chunk-groups served to a single client (so that connection reuse and pipelining can
    improve performance, once implemented).

  - For each cache, the publisher chooses a logical slot index.  (For details, see documentation on the logical cache
    model.)  This slot index should be consistent between requests for the cache to serve the same chunk.

***************************************************************

XXX: Temporary notes:

Object identifier (path) -> escrow-object (escrow & ID pair; do the IDs really matter?)

    The publisher will probably want to maintain a list of existing escrow-ID pairs for each object;
    it may also, at its option, create a new pair and return that.  (That is, it can choose to serve
    the request out of an escrow that's already been used to serve the object, or it can choose to serve
    the request out of an escrow that hasn't been.)

    This should be designed so that cache rollover/reuse between escrows is possible.

The publisher must also ensure that the metadata and data required to generate the puzzle is available
in the local catalog.  (The publisher doesn't use the catalog yet; that needs to be implemented.)

The publisher will also need to decide on LCM slot IDs for each chunk it asks a cache to serve.  These can vary per
cache, per escrow.  They should also be designed to support escrow rollover.

*/

const (
	chunksPerGroup    = 4
	bundlesPerRequest = 3
)

func (p *ContentPublisher) HandleContentRequest(ctx context.Context, req *ccmsg.ContentRequest) ([]*ccmsg.TicketBundle, error) {
	p.l.WithFields(logrus.Fields{
		"path":       req.Path,
		"rangeBegin": req.RangeBegin,
		"rangeEnd":   req.RangeEnd,
	}).Info("content request")

	for cache, depth := range req.BacklogDepth {
		p.l.WithFields(logrus.Fields{
			"cache":   base64.StdEncoding.EncodeToString([]byte(cache)),
			"backlog": depth,
		}).Debug("received cache backlog length")
	}

	// - The object's _path_ is used to ensure that the object exists, and that the specified chunks are in-cache and
	//   valid.  (This may be satisfied by the content catalog's cache, or may require contacting an upstream.)  (A
	//   future enhancement might require that the publisher fetch only the cipher-blocks that will be used in puzzle
	//   generation, instead of all of the cipher-blocks in the chunks.)
	p.l.Debug("pulling metadata and chunks into catalog")
	obj, err := p.catalog.GetData(ctx, &ccmsg.ContentRequest{Path: req.Path})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get metadata for requested object")
	}
	p.l.WithFields(logrus.Fields{
		"size": obj.ObjectSize(),
	}).Debug("received metadata and chunks")

	// - The _byte range_ is translated to a _chunk range_ depending on how the publisher would like to chunk the object.
	//   (Right now, we only support fixed-size chunks, but this is not inherent.)  The publisher may also choose how
	//   many chunks it would like to serve, and how many chunk-groups they will be divided into.  (The following steps
	//   are repeated for each chunk-group; the results are returned together in a single response to the client.)
	if req.RangeEnd != 0 && req.RangeEnd <= req.RangeBegin {
		// TODO: Return 4xx, since this is a bad request from the client.
		return nil, errors.New("invalid range")
	}

	// - The publisher selects a single escrow that will be used to service the request.
	p.l.Debug("selecting escrow")
	escrow, err := p.getEscrowByRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get escrow for request")
	}

	// - The publisher selects a set of caches that are enrolled in that escrow.  This selection should be designed to
	//   place the same chunks on the same caches (expanding the number in rotation as demand for the chunks grows), and
	//   to reuse the same caches for consecutive chunk-groups served to a single client (so that connection reuse and
	//   pipelining can improve performance, once implemented).
	p.l.Debug("selecting caches")
	ecs, err := models.EscrowCaches(qm.Load("Cache"), qm.Where("escrow_id = ?", escrow.Inner.ID)).All(ctx, p.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query database")
	}

	caches := []ParticipatingCache{}
	for _, participant := range ecs {
		caches = append(caches, ParticipatingCache{
			InnerMasterKey: participant.InnerMasterKey,
			Cache:          *participant.R.Cache,
		})
	}

	chunkRangeBegin := uint64(req.RangeBegin / obj.PolicyChunkSize())
	clientPublicKey := req.ClientPublicKey.PublicKey
	sequenceNo := req.SequenceNo

	numberOfBundles := bundlesPerRequest
	if req.RangeEnd != 0 {
		// TODO: if we need more than one bundle to reach RangeEnd,
		// calculate how many we need to send
		numberOfBundles = 1
	}

	bundles := []*ccmsg.TicketBundle{}
	for i := 0; i < numberOfBundles; i++ {
		bundle, err := p.generateBundle(ctx, escrow, caches, obj, req.Path, chunkRangeBegin, clientPublicKey, sequenceNo)
		if err != nil {
			return nil, err
		}
		chunkRangeBegin += uint64(len(bundle.GetTicketRequest()))
		sequenceNo++
		bundles = append(bundles, bundle)

		if chunkRangeBegin >= uint64(obj.ChunkCount()) {
			// we've reached the end of the object
			break
		}
	}

	return bundles, nil
}

func (p *ContentPublisher) generateBundle(ctx context.Context, escrow *Escrow, caches []ParticipatingCache, obj *catalog.ObjectMetadata, path string, chunkRangeBegin uint64, clientPublicKey ed25519.PublicKey, sequenceNo uint64) (*ccmsg.TicketBundle, error) {
	// XXX: this doesn't work with empty files
	if chunkRangeBegin >= uint64(obj.ChunkCount()) {
		return nil, errors.New("chunkRangeBegin beyond last chunk")
	}

	// TODO: Return multiple chunk-groups if appropriate.
	rangeEnd := chunkRangeBegin + chunksPerGroup
	if rangeEnd > uint64(obj.ChunkCount()) {
		rangeEnd = uint64(obj.ChunkCount())
	}

	p.l.WithFields(logrus.Fields{
		"chunkRangeBegin": chunkRangeBegin,
		"chunkRangeEnd":   rangeEnd,
	}).Info("content request")

	// - The _path_ and _chunk range_ are mapped to a list of _chunk identifiers_.  These are arbitrarily assigned by
	// the publisher.  (Our implementation uses the chunk's digest.)
	p.l.Debug("mapping chunk indices into chunk identifiers")
	chunkIndices := make([]uint64, 0, rangeEnd-chunkRangeBegin)
	chunkIDs := make([]common.ChunkID, 0, rangeEnd-chunkRangeBegin)
	for chunkIdx := chunkRangeBegin; chunkIdx < rangeEnd; chunkIdx++ {
		chunkID, err := p.getChunkID(obj, chunkIdx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get chunk ID")
		}

		chunkIDs = append(chunkIDs, chunkID)
		chunkIndices = append(chunkIndices, chunkIdx)
	}

	if len(caches) < len(chunkIndices) {
		return nil, errors.New(fmt.Sprintf("not enough caches: have %v; need %v", len(caches), len(chunkIndices)))
	}
	caches = caches[0:len(chunkIndices)]

	// - For each cache, the publisher chooses a logical slot index.  (For details, see documentation on the logical
	//   cache model.)  This slot index should be consistent between requests for the cache to serve the same chunk.

	// *********************************************************

	/*
		// XXX: If the object doesn't exist, we shouldn't reserve ticket numbers to satisfy the request!
		// XXX: This is what we need to remove to make the change to the content catalog; the `obj` here allows access to
		//      the entire contents of the object!
		obj, objID, err := escrow.GetObjectByPath(ctx, req.Path)
		if err != nil {
			return nil, errors.Wrap(err, "no object for path")
		}
	*/

	// XXX: Should be based on the upstream path, which the current implementation conflates with the request path.
	objID, err := generateObjectID(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate object ID")
	}
	p.reverseMapping[objID] = reverseMappingEntry{path: path}

	// Reserve a lottery ticket for each cache.  (Recall that lottery ticket numbers must be unique, and we are limited
	// in the number that we can issue during each blockchain block to the number that we declared in our begin-escrow
	// transaction.)
	// XXX: We need to make sure that these numbers are released to be reused if the request fails.
	p.l.Debug("reserving tickets")
	ticketNos, err := escrow.reserveTicketNumbers(len(caches))
	if err != nil {
		return nil, errors.Wrap(err, "failed to reserve ticket numbers")
	}

	p.l.Debug("building bundle parameters")
	bp := &BundleParams{
		Escrow:            escrow,
		RequestSequenceNo: sequenceNo,
		ClientPublicKey:   clientPublicKey,
		ObjectID:          objID,
	}
	for i, chunkIdx := range chunkIndices {
		// XXX: Need this to be non-zero; otherwise all of our chunks collide!
		bp.Entries = append(bp.Entries, BundleEntryParams{
			TicketNo: ticketNos[i],
			ChunkIdx: uint32(chunkIdx),
			ChunkID:  chunkIDs[i],
			Cache:    caches[i],
		})

		b, err := obj.GetChunk(uint32(chunkIdx))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get chunk")
		}
		bp.PlaintextChunks = append(bp.PlaintextChunks, b)
	}

	p.l.Debug("generating and signing bundle")
	batchSigner, err := batchsignature.NewTrivialBatchSigner(p.signer)
	if err != nil {
		return nil, err
	}
	gen := NewBundleGenerator(p.l, batchSigner)
	bundle, err := gen.GenerateTicketBundle(bp)
	if err != nil {
		return nil, err
	}

	// Attach metadata.
	// XXX: This needs to be covered by unit tests.
	bundle.Metadata = &ccmsg.ObjectMetadata{
		ChunkSize:  obj.PolicyChunkSize(),
		ObjectSize: obj.ObjectSize(),

		// TODO: don't hardcode those
		MinimumBacklogDepth:   2,
		BundleRequestInterval: 5,
	}

	p.l.Debug("done; returning bundle")
	return bundle, nil
}

func (p *ContentPublisher) assignSlot(path string, chunkIdx uint64, chunkID uint64) uint64 {
	// XXX: should depend on number of slots available to cache, etc.
	return chunkIdx
}

// TODO: XXX: Since object policy is, by definition, something that the publisher can set arbitrarily on a per-object
// basis, this should be the only place that these values are hardcoded.
func (p *ContentPublisher) objectPolicy(path string) (*catalog.ObjectPolicy, error) {
	return &catalog.ObjectPolicy{
		ChunkSize: 128 * 1024,
	}, nil
}

func (p *ContentPublisher) getChunkID(obj *catalog.ObjectMetadata, chunkIdx uint64) (common.ChunkID, error) {
	data, err := obj.GetChunk(uint32(chunkIdx))
	if err != nil {
		return common.ChunkID{}, errors.Wrap(err, "failed to get chunk data to generate ID")
	}

	var id common.ChunkID
	digest := sha512.Sum384(data)
	copy(id[:], digest[0:common.ChunkIDSize])

	p.l.WithFields(logrus.Fields{
		"chunkIdx": chunkIdx,
		"chunkID":  id,
	}).Debug("generating chunk ID")

	return id, nil
}

func (p *ContentPublisher) CacheMiss(ctx context.Context, req *ccmsg.CacheMissRequest) (*ccmsg.CacheMissResponse, error) {
	// TODO: How do we identify the cache submitting the request?

	objectID, err := common.BytesToObjectID(req.ObjectId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to interpret object ID")
	}

	rme, ok := p.reverseMapping[objectID]
	if !ok {
		return nil, errors.New("no reverse mapping found for object ID")
	}
	path := rme.path

	if req.RangeEnd != 0 && req.RangeEnd <= req.RangeBegin {
		return nil, errors.New("invalid range")
	}
	// if req.RangeEnd <= number-of-chunks-in-object ... invalid range

	objMeta, err := p.catalog.GetMetadata(ctx, path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get metadata for object")
	}

	// Convert object policy, which is required to convert chunk range into byte range.
	pol, err := p.objectPolicy(path)
	if err != nil {
		return nil, errors.New("failed to get object policy")
	}

	resp := ccmsg.CacheMissResponse{
		Chunks: []*ccmsg.Chunk{},
	}

	// XXX: Shouldn't we be telling the cache what chunk IDs it should expect, and providing enough information for it
	// to verify that it's getting the right data (e.g. a digest)?

	// Select logical cache slot for each chunk.
	for i := req.RangeBegin; i < req.RangeEnd; i++ {
		chunkID := i // XXX: Not true!
		slotIdx := p.assignSlot(path, i, chunkID)
		chunk, err := p.catalog.ChunkSource(ctx, req, path, objMeta)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get chunk source")
		}

		// TODO: we shouldn't need to modify the chunk afterwards
		// chunk.ChunkId = ChunkID,
		chunk.SlotIdx = slotIdx

		resp.Chunks = append(resp.Chunks, chunk)
	}

	resp.Metadata = &ccmsg.ObjectMetadata{
		ObjectSize: objMeta.ObjectSize(),
		ChunkSize:  uint64(pol.ChunkSize),
	}

	return &resp, nil
}

func (p *ContentPublisher) AddEscrowToDatabase(ctx context.Context, escrow *Escrow) error {
	if err := p.AddEscrow(escrow); err != nil {
		return errors.Wrap(err, "failed to add escrow to publisher")
	}
	if err := escrow.Inner.Insert(ctx, p.db, boil.Infer()); err != nil {
		return errors.Wrap(err, "failed to add escrow to database")
	}

	for _, c := range escrow.Caches {
		p.l.Info("Adding cache to database: ", c)
		err := c.Cache.Upsert(ctx, p.db, true, []string{"public_key"}, boil.Whitelist("inetaddr", "port"), boil.Infer())
		if err != nil {
			return errors.Wrap(err, "failed to add cache to database")
		}

		ec := models.EscrowCache{
			EscrowID:       escrow.Inner.ID,
			CacheID:        c.Cache.ID,
			InnerMasterKey: c.InnerMasterKey,
		}
		err = ec.Upsert(ctx, p.db, false, []string{"escrow_id", "cache_id"}, boil.Whitelist("inner_master_key"), boil.Infer())
		if err != nil {
			return errors.Wrap(err, "failed to link cache to escrow")
		}
	}

	return nil
}

// XXX: This is, obviously, temporary.  We should be using object IDs that are larger than 64 bits, among other
// problems.  We also must account for the fact that the object stored at a path may change (e.g. when the mtime/etag
// are updated).
func generateObjectID(path string) (common.ObjectID, error) {
	digest := sha256.Sum256([]byte(path))
	return common.BytesToObjectID(digest[0:common.ObjectIDSize])
}
