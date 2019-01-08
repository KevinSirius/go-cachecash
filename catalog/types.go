package catalog

import (
	"context"
	"net/http"

	"github.com/kelleyk/go-cachecash/ccmsg"
)

//go:generate stringer -type=ObjectStatus

type ObjectStatus int

const (
	StatusUnknown ObjectStatus = iota
	StatusOK
	StatusNotFound
	StatusInternalError
	StatusUpstreamUnreachable
	StatusUpstreamError
)

// FetchResult describes the metadata, data, and/or errors returned by an Upstream in response to a single request.
// They are consumed by the catalog, which uses them to update its cache.
type FetchResult struct {
	// XXX: This should be changed; this struct is supposed to be protocol-agnostic.
	header http.Header
	data   []byte
	status ObjectStatus
}

type Upstream interface {
	// FetchData ensures that metadata is fresh, and also that the indicated blocks are available in the cache.  An
	// empty list of block indices is allowed; this ensures metadata freshness but does not pull any data blocks.
	//
	// forceMetadata requires that object metadata be fetched even if it would not otherwise be fetched.
	//
	// rangeEnd must be >= rangeBegin.  rangeEnd == 0 means "continue to he end of the object".
	//
	// Cases:
	// - We want to fetch metadata only.
	// - We want to fetch metadata *and* a series of blocks.
	// - We have metadata that we already believe to be be valid, so we don't necessarily need to fetch it, if that's
	//   any extra effort.  We want to fetch a series of blocks.
	//
	// Questions:
	// - Some upstreams will require CacheCash (not upstream) metadata for the object.  For example, the HTTP upstream
	//   will need to know block sizes in order to translate block indices into byte ranges.  How should this be done?
	//
	FetchData(ctx context.Context, path string, forceMetadata bool, rangeBegin, rangeEnd uint) (*FetchResult, error)

	CacheMiss(path string, rangeBegin, rangeEnd uint64) (*ccmsg.CacheMissResponse, error)
}

type ContentCatalog interface {
	GetObjectMetadata(ctx context.Context, path string) (*ObjectMetadata, error)

	CacheMiss(path string, rangeBegin, rangeEnd uint64) (*ccmsg.CacheMissResponse, error)
}

type ContentLocator interface {
	GetContentSource(ctx context.Context, req *ccmsg.CacheMissRequest) (*ccmsg.CacheMissResponse, error)
}
