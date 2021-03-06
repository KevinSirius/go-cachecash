package publisher

import (
	"net"

	cachecash "github.com/cachecashproject/go-cachecash"
	"github.com/cachecashproject/go-cachecash/batchsignature"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/colocationpuzzle"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/cachecashproject/go-cachecash/publisher/models"
	"github.com/cachecashproject/go-cachecash/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ed25519"
)

type Escrow struct {
	Publisher *ContentPublisher
	Inner     models.Escrow
	Caches    []*ParticipatingCache
	// Each row stores an index into Caches; each bundle width of indexes must
	// be a group of unique indexes; currently this is hardcoded. lookups that
	// are dealing with error cases such as clients which cannot use a given
	// cache select the next cache from Caches without violating the uniqueness
	// rule.
	lookup *[]int
}

// The info object does not need to have its keys populated.
func (p *ContentPublisher) NewEscrow(info *ccmsg.EscrowInfo) (*Escrow, error) {
	if info.DrawDelay == 0 {
		return nil, errors.New("draw delay must be at least one block")
	}
	if info.ExpirationDelay == 0 {
		return nil, errors.New("expiration delay must be at least one block")
	}
	if info.StartBlock == 0 {
		return nil, errors.New("start block number must be set")
	}
	// TODO: Perform additional validation on TicketsPerBlock.
	if len(info.TicketsPerBlock) == 0 {
		return nil, errors.New("tickets-per-block may not be empty")
	}
	id, err := common.BytesToEscrowID(info.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to interpret escrow ID")
	}

	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate keypair")
	}

	// TODO: Should we set info.PublicKey and info.PublisherPublicKey?
	return &Escrow{
		Publisher: p,
		Inner: models.Escrow{
			Txid:       id,
			PublicKey:  pub,
			PrivateKey: priv,
			State:      models.EscrowStateOk,
		},
	}, nil
}

func (e *Escrow) reserveTicketNumbers(qty int) ([]uint64, error) {
	nos := make([]uint64, qty)
	for i := uint64(0); i < uint64(qty); i++ {
		nos = append(nos, i)
	}
	return nos, nil
}

// CalculateLookup calculates chunk lookup tables for maglev style chunk allocations
// This can fail if there are not enough caches.
// This routine operates off populated structs - load from database before calling
func (e *Escrow) CalculateLookup() error {
	permutations := [][]uint64{}
	for _, cache := range e.Caches {
		permutation, err := e.Publisher.getCachePermutation(string(cache.Cache.PublicKey), uint(len(e.Caches)))
		if err != nil {
			return err
		}
		permutations = append(permutations, permutation)
	}
	caches := len(permutations)
	if caches == 0 {
		return errors.New("Cannot calculate a lookup table with 0 caches")
	}
	// Each row stores an index into e.Caches
	lookup := make([]int, len(permutations[0]))
	next := make([]int, caches)
	for pos := range lookup {
		lookup[pos] = -1
	}
	n := 0
	for {
		for i := 0; i < caches; i++ {
			c := permutations[i][next[i]]
			for lookup[c] >= 0 {
				next[i]++
				c = permutations[i][next[i]]
			}
			lookup[c] = i
			next[i]++
			n++
			if n == len(lookup) {
				e.lookup = &lookup
				return nil
			}
		}
	}
}

type ParticipatingCache struct {
	Cache          models.Cache
	InnerMasterKey []byte
}

func (p ParticipatingCache) PublicKey() ed25519.PublicKey {
	return p.Cache.PublicKey
}

func (p ParticipatingCache) Inetaddr() net.IP {
	return p.Cache.Inetaddr.To4()
}

func (p ParticipatingCache) Inet6Addr() net.IP {
	return p.Cache.Inet6addr.To16()
}

func (p ParticipatingCache) Port() uint32 {
	return p.Cache.Port
}

// BundleParams is everything necessary to generate a complete TicketBundle message.
type BundleParams struct {
	Escrow            *Escrow         // XXX: Do we need this?
	ObjectID          common.ObjectID // This is a per-escrow value.
	Entries           []BundleEntryParams
	PlaintextChunks   [][]byte
	RequestSequenceNo uint64
	ClientPublicKey   ed25519.PublicKey
}

type BundleEntryParams struct {
	TicketNo uint64
	ChunkIdx uint32
	ChunkID  common.ChunkID
	Cache    ParticipatingCache
}

func NewBundleGenerator(l *logrus.Logger, signer batchsignature.BatchSigner) *BundleGenerator {
	return &BundleGenerator{
		l:      l,
		Signer: signer,
		PuzzleParams: &colocationpuzzle.Parameters{
			Rounds:      2,
			StartOffset: 0, // TODO: Not respected yet.
			StartRange:  0,
		},
	}
}

type BundleGenerator struct {
	l            *logrus.Logger
	PuzzleParams *colocationpuzzle.Parameters
	Signer       batchsignature.BatchSigner
}

// XXX: Attach this function to a struct containing configuration data (like e.g. puzzle parameters), or add those
// things as arguments.
func (gen *BundleGenerator) GenerateTicketBundle(bp *BundleParams) (*ccmsg.TicketBundle, error) {
	resp := &ccmsg.TicketBundle{
		// PublisherPublicKey: cachecash.PublicKeyMessage(e.Publisher.PublicKey),
		// EscrowPublicKey:   cachecash.PublicKeyMessage(e.PublicKey),
		Remainder: &ccmsg.TicketBundleRemainder{
			RequestSequenceNo: bp.RequestSequenceNo,
			EscrowId:          bp.Escrow.Inner.Txid[:],
			ObjectId:          bp.ObjectID[:],
			// PuzzleInfo is filled in later
			ClientPublicKey: cachecash.PublicKeyMessage(bp.ClientPublicKey),
		},
	}

	if len(bp.Entries) == 0 {
		return nil, errors.New("must serve client at least one chunk")
	}

	// Generate inner keys (one per cache) using our keyed PRF.
	var innerKeys, innerIVs [][]byte
	for _, bep := range bp.Entries {
		prfInput := []byte(bp.ClientPublicKey) // XXX:
		k, err := util.KeyedPRF(prfInput, uint32(bp.RequestSequenceNo), bep.Cache.InnerMasterKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate inner key")
		}
		innerKeys = append(innerKeys, k)

		iv, err := util.KeyedPRF(util.Uint64ToLE(uint64(bep.ChunkIdx)), uint32(bp.RequestSequenceNo), k)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate inner IV")
		}
		innerIVs = append(innerIVs, iv)
	}

	chunkIndices := make([]uint32, len(bp.Entries))
	for i := 0; i < len(bp.Entries); i++ {
		bep := bp.Entries[i]
		chunkIndices[i] = bep.ChunkIdx

		// Generate a ticket-request for each cache.
		resp.TicketRequest = append(resp.TicketRequest, &ccmsg.TicketRequest{
			ChunkIdx:       uint64(bep.ChunkIdx),
			ChunkId:        bep.ChunkID[:],
			CachePublicKey: cachecash.PublicKeyMessage(bep.Cache.PublicKey()),

			// XXX: Why is 'inner_key' in this message?  Regardless, we need the submessage not to be nil, or we'll get
			// a nil deref when computing the digest.
			InnerKey: &ccmsg.BlockKey{Key: nil},
		})

		resp.CacheInfo = append(resp.CacheInfo, &ccmsg.CacheInfo{
			Addr: &ccmsg.NetworkAddress{
				Inetaddr:  bep.Cache.Inetaddr(),
				Inet6Addr: bep.Cache.Inet6Addr(),
				Port:      bep.Cache.Port(),
			},
			Pubkey: &ccmsg.PublicKey{
				PublicKey: bep.Cache.PublicKey(),
			},
		})

		// Generate a lottery ticket 1 for each cache.
		resp.TicketL1 = append(resp.TicketL1, &ccmsg.TicketL1{
			TicketNo:       bep.TicketNo,
			CachePublicKey: cachecash.PublicKeyMessage(bep.Cache.PublicKey()), // XXX: Does this need to be repeated here?
		})
	}

	// Generate a colocation puzzle for the client to solve.
	gen.l.WithFields(logrus.Fields{
		"chunkIdx": chunkIndices,
	}).Info("generating puzzle")
	puzzle, err := colocationpuzzle.Generate(*gen.PuzzleParams, bp.PlaintextChunks, innerKeys, innerIVs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate colocation puzzle")
	}
	resp.Remainder.PuzzleInfo = &ccmsg.ColocationPuzzleInfo{
		Goal:        puzzle.Goal,
		Rounds:      gen.PuzzleParams.Rounds,
		StartOffset: uint64(gen.PuzzleParams.StartOffset), // XXX: Make typing consistent!
		StartRange:  uint64(gen.PuzzleParams.StartRange),
	}
	gen.l.WithFields(logrus.Fields{
		"initialOffset": puzzle.Offset,
		// "goal":          hex.EncodeToString(puzzle.Goal),
		// "secret":        hex.EncodeToString(puzzle.Secret),
	}).Info("generated colocation puzzle")

	// Generate a lottery ticket 2 and then marshal and encrypt it using a key and IV taken from the colocation puzzle's secret.
	ticketL2 := &ccmsg.TicketL2{}
	for _, k := range innerKeys {
		ticketL2.InnerSessionKey = append(ticketL2.InnerSessionKey, &ccmsg.BlockKey{Key: k})
	}
	resp.EncryptedTicketL2, err = common.EncryptTicketL2(puzzle, ticketL2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal and encrypt ticket L2")
	}

	// Generate our batch signature (BHT).
	cd := resp.CanonicalDigest()
	sig, err := gen.Signer.BatchSign(cd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign ticket bundle")
	}
	resp.BatchSig = sig

	// Done!
	return resp, nil
}
