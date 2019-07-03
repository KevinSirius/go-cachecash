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

	"github.com/cachecashproject/go-cachecash/common"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Escrow is an object representing the database table.
type Escrow struct {
	Txid           common.EscrowID `boil:"txid" json:"txid" toml:"txid" yaml:"txid"`
	InnerMasterKey []byte          `boil:"inner_master_key" json:"inner_master_key" toml:"inner_master_key" yaml:"inner_master_key"`
	OuterMasterKey []byte          `boil:"outer_master_key" json:"outer_master_key" toml:"outer_master_key" yaml:"outer_master_key"`
	Slots          uint64          `boil:"slots" json:"slots" toml:"slots" yaml:"slots"`
	PublisherAddr  string          `boil:"publisher_addr" json:"publisher_addr" toml:"publisher_addr" yaml:"publisher_addr"`

	R *escrowR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L escrowL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EscrowColumns = struct {
	Txid           string
	InnerMasterKey string
	OuterMasterKey string
	Slots          string
	PublisherAddr  string
}{
	Txid:           "txid",
	InnerMasterKey: "inner_master_key",
	OuterMasterKey: "outer_master_key",
	Slots:          "slots",
	PublisherAddr:  "publisher_addr",
}

// Generated where

type whereHelpercommon_EscrowID struct{ field string }

func (w whereHelpercommon_EscrowID) EQ(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpercommon_EscrowID) NEQ(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpercommon_EscrowID) LT(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpercommon_EscrowID) LTE(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpercommon_EscrowID) GT(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpercommon_EscrowID) GTE(x common.EscrowID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var EscrowWhere = struct {
	Txid           whereHelpercommon_EscrowID
	InnerMasterKey whereHelper__byte
	OuterMasterKey whereHelper__byte
	Slots          whereHelperuint64
	PublisherAddr  whereHelperstring
}{
	Txid:           whereHelpercommon_EscrowID{field: `txid`},
	InnerMasterKey: whereHelper__byte{field: `inner_master_key`},
	OuterMasterKey: whereHelper__byte{field: `outer_master_key`},
	Slots:          whereHelperuint64{field: `slots`},
	PublisherAddr:  whereHelperstring{field: `publisher_addr`},
}

// EscrowRels is where relationship names are stored.
var EscrowRels = struct {
}{}

// escrowR is where relationships are stored.
type escrowR struct {
}

// NewStruct creates a new relationship struct
func (*escrowR) NewStruct() *escrowR {
	return &escrowR{}
}

// escrowL is where Load methods for each relationship are stored.
type escrowL struct{}

var (
	escrowColumns               = []string{"txid", "inner_master_key", "outer_master_key", "slots", "publisher_addr"}
	escrowColumnsWithoutDefault = []string{"txid", "inner_master_key", "outer_master_key", "slots", "publisher_addr"}
	escrowColumnsWithDefault    = []string{}
	escrowPrimaryKeyColumns     = []string{"txid"}
)

type (
	// EscrowSlice is an alias for a slice of pointers to Escrow.
	// This should generally be used opposed to []Escrow.
	EscrowSlice []*Escrow
	// EscrowHook is the signature for custom Escrow hook methods
	EscrowHook func(context.Context, boil.ContextExecutor, *Escrow) error

	escrowQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	escrowType                 = reflect.TypeOf(&Escrow{})
	escrowMapping              = queries.MakeStructMapping(escrowType)
	escrowPrimaryKeyMapping, _ = queries.BindMapping(escrowType, escrowMapping, escrowPrimaryKeyColumns)
	escrowInsertCacheMut       sync.RWMutex
	escrowInsertCache          = make(map[string]insertCache)
	escrowUpdateCacheMut       sync.RWMutex
	escrowUpdateCache          = make(map[string]updateCache)
	escrowUpsertCacheMut       sync.RWMutex
	escrowUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var escrowBeforeInsertHooks []EscrowHook
var escrowBeforeUpdateHooks []EscrowHook
var escrowBeforeDeleteHooks []EscrowHook
var escrowBeforeUpsertHooks []EscrowHook

var escrowAfterInsertHooks []EscrowHook
var escrowAfterSelectHooks []EscrowHook
var escrowAfterUpdateHooks []EscrowHook
var escrowAfterDeleteHooks []EscrowHook
var escrowAfterUpsertHooks []EscrowHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Escrow) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Escrow) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Escrow) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Escrow) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Escrow) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Escrow) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Escrow) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Escrow) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Escrow) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range escrowAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEscrowHook registers your hook function for all future operations.
func AddEscrowHook(hookPoint boil.HookPoint, escrowHook EscrowHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		escrowBeforeInsertHooks = append(escrowBeforeInsertHooks, escrowHook)
	case boil.BeforeUpdateHook:
		escrowBeforeUpdateHooks = append(escrowBeforeUpdateHooks, escrowHook)
	case boil.BeforeDeleteHook:
		escrowBeforeDeleteHooks = append(escrowBeforeDeleteHooks, escrowHook)
	case boil.BeforeUpsertHook:
		escrowBeforeUpsertHooks = append(escrowBeforeUpsertHooks, escrowHook)
	case boil.AfterInsertHook:
		escrowAfterInsertHooks = append(escrowAfterInsertHooks, escrowHook)
	case boil.AfterSelectHook:
		escrowAfterSelectHooks = append(escrowAfterSelectHooks, escrowHook)
	case boil.AfterUpdateHook:
		escrowAfterUpdateHooks = append(escrowAfterUpdateHooks, escrowHook)
	case boil.AfterDeleteHook:
		escrowAfterDeleteHooks = append(escrowAfterDeleteHooks, escrowHook)
	case boil.AfterUpsertHook:
		escrowAfterUpsertHooks = append(escrowAfterUpsertHooks, escrowHook)
	}
}

// One returns a single escrow record from the query.
func (q escrowQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Escrow, error) {
	o := &Escrow{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for escrow")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Escrow records from the query.
func (q escrowQuery) All(ctx context.Context, exec boil.ContextExecutor) (EscrowSlice, error) {
	var o []*Escrow

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Escrow slice")
	}

	if len(escrowAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Escrow records in the query.
func (q escrowQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count escrow rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q escrowQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if escrow exists")
	}

	return count > 0, nil
}

// Escrows retrieves all the records using an executor.
func Escrows(mods ...qm.QueryMod) escrowQuery {
	mods = append(mods, qm.From("\"escrow\""))
	return escrowQuery{NewQuery(mods...)}
}

// FindEscrow retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEscrow(ctx context.Context, exec boil.ContextExecutor, txid common.EscrowID, selectCols ...string) (*Escrow, error) {
	escrowObj := &Escrow{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"escrow\" where \"txid\"=?", sel,
	)

	q := queries.Raw(query, txid)

	err := q.Bind(ctx, exec, escrowObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from escrow")
	}

	return escrowObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Escrow) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no escrow provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(escrowColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	escrowInsertCacheMut.RLock()
	cache, cached := escrowInsertCache[key]
	escrowInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			escrowColumns,
			escrowColumnsWithDefault,
			escrowColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(escrowType, escrowMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(escrowType, escrowMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"escrow\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"escrow\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"escrow\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, escrowPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into escrow")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Txid,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for escrow")
	}

CacheNoHooks:
	if !cached {
		escrowInsertCacheMut.Lock()
		escrowInsertCache[key] = cache
		escrowInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Escrow.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Escrow) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	escrowUpdateCacheMut.RLock()
	cache, cached := escrowUpdateCache[key]
	escrowUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			escrowColumns,
			escrowPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update escrow, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"escrow\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, escrowPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(escrowType, escrowMapping, append(wl, escrowPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update escrow row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for escrow")
	}

	if !cached {
		escrowUpdateCacheMut.Lock()
		escrowUpdateCache[key] = cache
		escrowUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q escrowQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for escrow")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for escrow")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EscrowSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), escrowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"escrow\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, escrowPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in escrow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all escrow")
	}
	return rowsAff, nil
}

// Delete deletes a single Escrow record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Escrow) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Escrow provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), escrowPrimaryKeyMapping)
	sql := "DELETE FROM \"escrow\" WHERE \"txid\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from escrow")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for escrow")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q escrowQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no escrowQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from escrow")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for escrow")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EscrowSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Escrow slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(escrowBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), escrowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"escrow\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, escrowPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from escrow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for escrow")
	}

	if len(escrowAfterDeleteHooks) != 0 {
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
func (o *Escrow) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEscrow(ctx, exec, o.Txid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EscrowSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EscrowSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), escrowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"escrow\".* FROM \"escrow\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, escrowPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EscrowSlice")
	}

	*o = slice

	return nil
}

// EscrowExists checks if the Escrow row exists.
func EscrowExists(ctx context.Context, exec boil.ContextExecutor, txid common.EscrowID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"escrow\" where \"txid\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, txid)
	}

	row := exec.QueryRowContext(ctx, sql, txid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if escrow exists")
	}

	return exists, nil
}
