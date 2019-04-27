// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"golang.org/x/crypto/ed25519"
	"net"
)

// Cache is an object representing the database table.
type Cache struct {
	PublicKey   ed25519.PublicKey `boil:"public_key" json:"public_key" toml:"public_key" yaml:"public_key"`
	Version     string            `boil:"version" json:"version" toml:"version" yaml:"version"`
	FreeMemory  uint64            `boil:"free_memory" json:"free_memory" toml:"free_memory" yaml:"free_memory"`
	TotalMemory uint64            `boil:"total_memory" json:"total_memory" toml:"total_memory" yaml:"total_memory"`
	FreeDisk    uint64            `boil:"free_disk" json:"free_disk" toml:"free_disk" yaml:"free_disk"`
	TotalDisk   uint64            `boil:"total_disk" json:"total_disk" toml:"total_disk" yaml:"total_disk"`
	StartupTime time.Time         `boil:"startup_time" json:"startup_time" toml:"startup_time" yaml:"startup_time"`
	ExternalIP  net.IP            `boil:"external_ip" json:"external_ip" toml:"external_ip" yaml:"external_ip"`
	ContactURL  string            `boil:"contact_url" json:"contact_url" toml:"contact_url" yaml:"contact_url"`
	LastPing    time.Time         `boil:"last_ping" json:"last_ping" toml:"last_ping" yaml:"last_ping"`

	R *cacheR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cacheL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CacheColumns = struct {
	PublicKey   string
	Version     string
	FreeMemory  string
	TotalMemory string
	FreeDisk    string
	TotalDisk   string
	StartupTime string
	ExternalIP  string
	ContactURL  string
	LastPing    string
}{
	PublicKey:   "public_key",
	Version:     "version",
	FreeMemory:  "free_memory",
	TotalMemory: "total_memory",
	FreeDisk:    "free_disk",
	TotalDisk:   "total_disk",
	StartupTime: "startup_time",
	ExternalIP:  "external_ip",
	ContactURL:  "contact_url",
	LastPing:    "last_ping",
}

// Generated where

type whereHelpered25519_PublicKey struct{ field string }

func (w whereHelpered25519_PublicKey) EQ(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpered25519_PublicKey) NEQ(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpered25519_PublicKey) LT(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpered25519_PublicKey) LTE(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpered25519_PublicKey) GT(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpered25519_PublicKey) GTE(x ed25519.PublicKey) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernet_IP struct{ field string }

func (w whereHelpernet_IP) EQ(x net.IP) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelpernet_IP) NEQ(x net.IP) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelpernet_IP) LT(x net.IP) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelpernet_IP) LTE(x net.IP) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelpernet_IP) GT(x net.IP) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelpernet_IP) GTE(x net.IP) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CacheWhere = struct {
	PublicKey   whereHelpered25519_PublicKey
	Version     whereHelperstring
	FreeMemory  whereHelperuint64
	TotalMemory whereHelperuint64
	FreeDisk    whereHelperuint64
	TotalDisk   whereHelperuint64
	StartupTime whereHelpertime_Time
	ExternalIP  whereHelpernet_IP
	ContactURL  whereHelperstring
	LastPing    whereHelpertime_Time
}{
	PublicKey:   whereHelpered25519_PublicKey{field: `public_key`},
	Version:     whereHelperstring{field: `version`},
	FreeMemory:  whereHelperuint64{field: `free_memory`},
	TotalMemory: whereHelperuint64{field: `total_memory`},
	FreeDisk:    whereHelperuint64{field: `free_disk`},
	TotalDisk:   whereHelperuint64{field: `total_disk`},
	StartupTime: whereHelpertime_Time{field: `startup_time`},
	ExternalIP:  whereHelpernet_IP{field: `external_ip`},
	ContactURL:  whereHelperstring{field: `contact_url`},
	LastPing:    whereHelpertime_Time{field: `last_ping`},
}

// CacheRels is where relationship names are stored.
var CacheRels = struct {
}{}

// cacheR is where relationships are stored.
type cacheR struct {
}

// NewStruct creates a new relationship struct
func (*cacheR) NewStruct() *cacheR {
	return &cacheR{}
}

// cacheL is where Load methods for each relationship are stored.
type cacheL struct{}

var (
	cacheColumns               = []string{"public_key", "version", "free_memory", "total_memory", "free_disk", "total_disk", "startup_time", "external_ip", "contact_url", "last_ping"}
	cacheColumnsWithoutDefault = []string{"public_key", "version", "free_memory", "total_memory", "free_disk", "total_disk", "startup_time", "external_ip", "contact_url", "last_ping"}
	cacheColumnsWithDefault    = []string{}
	cachePrimaryKeyColumns     = []string{"public_key"}
)

type (
	// CacheSlice is an alias for a slice of pointers to Cache.
	// This should generally be used opposed to []Cache.
	CacheSlice []*Cache
	// CacheHook is the signature for custom Cache hook methods
	CacheHook func(context.Context, boil.ContextExecutor, *Cache) error

	cacheQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cacheType                 = reflect.TypeOf(&Cache{})
	cacheMapping              = queries.MakeStructMapping(cacheType)
	cachePrimaryKeyMapping, _ = queries.BindMapping(cacheType, cacheMapping, cachePrimaryKeyColumns)
	cacheInsertCacheMut       sync.RWMutex
	cacheInsertCache          = make(map[string]insertCache)
	cacheUpdateCacheMut       sync.RWMutex
	cacheUpdateCache          = make(map[string]updateCache)
	cacheUpsertCacheMut       sync.RWMutex
	cacheUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cacheBeforeInsertHooks []CacheHook
var cacheBeforeUpdateHooks []CacheHook
var cacheBeforeDeleteHooks []CacheHook
var cacheBeforeUpsertHooks []CacheHook

var cacheAfterInsertHooks []CacheHook
var cacheAfterSelectHooks []CacheHook
var cacheAfterUpdateHooks []CacheHook
var cacheAfterDeleteHooks []CacheHook
var cacheAfterUpsertHooks []CacheHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cache) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cache) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cache) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cache) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cache) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cache) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cache) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cache) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cache) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cacheAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCacheHook registers your hook function for all future operations.
func AddCacheHook(hookPoint boil.HookPoint, cacheHook CacheHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cacheBeforeInsertHooks = append(cacheBeforeInsertHooks, cacheHook)
	case boil.BeforeUpdateHook:
		cacheBeforeUpdateHooks = append(cacheBeforeUpdateHooks, cacheHook)
	case boil.BeforeDeleteHook:
		cacheBeforeDeleteHooks = append(cacheBeforeDeleteHooks, cacheHook)
	case boil.BeforeUpsertHook:
		cacheBeforeUpsertHooks = append(cacheBeforeUpsertHooks, cacheHook)
	case boil.AfterInsertHook:
		cacheAfterInsertHooks = append(cacheAfterInsertHooks, cacheHook)
	case boil.AfterSelectHook:
		cacheAfterSelectHooks = append(cacheAfterSelectHooks, cacheHook)
	case boil.AfterUpdateHook:
		cacheAfterUpdateHooks = append(cacheAfterUpdateHooks, cacheHook)
	case boil.AfterDeleteHook:
		cacheAfterDeleteHooks = append(cacheAfterDeleteHooks, cacheHook)
	case boil.AfterUpsertHook:
		cacheAfterUpsertHooks = append(cacheAfterUpsertHooks, cacheHook)
	}
}

// One returns a single cache record from the query.
func (q cacheQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Cache, error) {
	o := &Cache{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Cache records from the query.
func (q cacheQuery) All(ctx context.Context, exec boil.ContextExecutor) (CacheSlice, error) {
	var o []*Cache

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Cache slice")
	}

	if len(cacheAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Cache records in the query.
func (q cacheQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cacheQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if caches exists")
	}

	return count > 0, nil
}

// Caches retrieves all the records using an executor.
func Caches(mods ...qm.QueryMod) cacheQuery {
	mods = append(mods, qm.From("\"caches\""))
	return cacheQuery{NewQuery(mods...)}
}

// FindCache retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCache(ctx context.Context, exec boil.ContextExecutor, publicKey ed25519.PublicKey, selectCols ...string) (*Cache, error) {
	cacheObj := &Cache{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"caches\" where \"public_key\"=?", sel,
	)

	q := queries.Raw(query, publicKey)

	err := q.Bind(ctx, exec, cacheObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from caches")
	}

	return cacheObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Cache) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no caches provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cacheColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cacheInsertCacheMut.RLock()
	cache, cached := cacheInsertCache[key]
	cacheInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cacheColumns,
			cacheColumnsWithDefault,
			cacheColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cacheType, cacheMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cacheType, cacheMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"caches\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"caches\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, cachePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into caches")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PublicKey,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for caches")
	}

CacheNoHooks:
	if !cached {
		cacheInsertCacheMut.Lock()
		cacheInsertCache[key] = cache
		cacheInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Cache.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Cache) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cacheUpdateCacheMut.RLock()
	cache, cached := cacheUpdateCache[key]
	cacheUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cacheColumns,
			cachePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, cachePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cacheType, cacheMapping, append(wl, cachePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for caches")
	}

	if !cached {
		cacheUpdateCacheMut.Lock()
		cacheUpdateCache[key] = cache
		cacheUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cacheQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CacheSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cachePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cachePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cache slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cache")
	}
	return rowsAff, nil
}

// Delete deletes a single Cache record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cache) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Cache provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cachePrimaryKeyMapping)
	sql := "DELETE FROM \"caches\" WHERE \"public_key\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cacheQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cacheQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CacheSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Cache slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(cacheBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cachePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cachePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cache slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for caches")
	}

	if len(cacheAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cache) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCache(ctx, exec, o.PublicKey)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CacheSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CacheSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cachePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"caches\".* FROM \"caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cachePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CacheSlice")
	}

	*o = slice

	return nil
}

// CacheExists checks if the Cache row exists.
func CacheExists(ctx context.Context, exec boil.ContextExecutor, publicKey ed25519.PublicKey) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"caches\" where \"public_key\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, publicKey)
	}

	row := exec.QueryRowContext(ctx, sql, publicKey)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if caches exists")
	}

	return exists, nil
}
