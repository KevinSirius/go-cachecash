// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cachecashproject/go-cachecash/ledger/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RawTX is an object representing the database table.
type RawTX struct {
	Txid  models.TXID `boil:"txid" json:"txid" toml:"txid" yaml:"txid"`
	Bytes []byte      `boil:"bytes" json:"bytes" toml:"bytes" yaml:"bytes"`

	R *rawTXR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rawTXL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RawTXColumns = struct {
	Txid  string
	Bytes string
}{
	Txid:  "txid",
	Bytes: "bytes",
}

// Generated where

type whereHelpermodels_TXID struct{ field string }

func (w whereHelpermodels_TXID) EQ(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpermodels_TXID) NEQ(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpermodels_TXID) LT(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpermodels_TXID) LTE(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpermodels_TXID) GT(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpermodels_TXID) GTE(x models.TXID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var RawTXWhere = struct {
	Txid  whereHelpermodels_TXID
	Bytes whereHelper__byte
}{
	Txid:  whereHelpermodels_TXID{field: "\"raw_tx\".\"txid\""},
	Bytes: whereHelper__byte{field: "\"raw_tx\".\"bytes\""},
}

// RawTXRels is where relationship names are stored.
var RawTXRels = struct {
}{}

// rawTXR is where relationships are stored.
type rawTXR struct {
}

// NewStruct creates a new relationship struct
func (*rawTXR) NewStruct() *rawTXR {
	return &rawTXR{}
}

// rawTXL is where Load methods for each relationship are stored.
type rawTXL struct{}

var (
	rawTXAllColumns            = []string{"txid", "bytes"}
	rawTXColumnsWithoutDefault = []string{"txid", "bytes"}
	rawTXColumnsWithDefault    = []string{}
	rawTXPrimaryKeyColumns     = []string{"txid"}
)

type (
	// RawTXSlice is an alias for a slice of pointers to RawTX.
	// This should generally be used opposed to []RawTX.
	RawTXSlice []*RawTX
	// RawTXHook is the signature for custom RawTX hook methods
	RawTXHook func(context.Context, boil.ContextExecutor, *RawTX) error

	rawTXQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rawTXType                 = reflect.TypeOf(&RawTX{})
	rawTXMapping              = queries.MakeStructMapping(rawTXType)
	rawTXPrimaryKeyMapping, _ = queries.BindMapping(rawTXType, rawTXMapping, rawTXPrimaryKeyColumns)
	rawTXInsertCacheMut       sync.RWMutex
	rawTXInsertCache          = make(map[string]insertCache)
	rawTXUpdateCacheMut       sync.RWMutex
	rawTXUpdateCache          = make(map[string]updateCache)
	rawTXUpsertCacheMut       sync.RWMutex
	rawTXUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rawTXBeforeInsertHooks []RawTXHook
var rawTXBeforeUpdateHooks []RawTXHook
var rawTXBeforeDeleteHooks []RawTXHook
var rawTXBeforeUpsertHooks []RawTXHook

var rawTXAfterInsertHooks []RawTXHook
var rawTXAfterSelectHooks []RawTXHook
var rawTXAfterUpdateHooks []RawTXHook
var rawTXAfterDeleteHooks []RawTXHook
var rawTXAfterUpsertHooks []RawTXHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RawTX) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RawTX) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RawTX) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RawTX) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RawTX) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RawTX) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RawTX) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RawTX) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RawTX) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawTXAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRawTXHook registers your hook function for all future operations.
func AddRawTXHook(hookPoint boil.HookPoint, rawTXHook RawTXHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rawTXBeforeInsertHooks = append(rawTXBeforeInsertHooks, rawTXHook)
	case boil.BeforeUpdateHook:
		rawTXBeforeUpdateHooks = append(rawTXBeforeUpdateHooks, rawTXHook)
	case boil.BeforeDeleteHook:
		rawTXBeforeDeleteHooks = append(rawTXBeforeDeleteHooks, rawTXHook)
	case boil.BeforeUpsertHook:
		rawTXBeforeUpsertHooks = append(rawTXBeforeUpsertHooks, rawTXHook)
	case boil.AfterInsertHook:
		rawTXAfterInsertHooks = append(rawTXAfterInsertHooks, rawTXHook)
	case boil.AfterSelectHook:
		rawTXAfterSelectHooks = append(rawTXAfterSelectHooks, rawTXHook)
	case boil.AfterUpdateHook:
		rawTXAfterUpdateHooks = append(rawTXAfterUpdateHooks, rawTXHook)
	case boil.AfterDeleteHook:
		rawTXAfterDeleteHooks = append(rawTXAfterDeleteHooks, rawTXHook)
	case boil.AfterUpsertHook:
		rawTXAfterUpsertHooks = append(rawTXAfterUpsertHooks, rawTXHook)
	}
}

// One returns a single rawTX record from the query.
func (q rawTXQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RawTX, error) {
	o := &RawTX{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for raw_tx")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RawTX records from the query.
func (q rawTXQuery) All(ctx context.Context, exec boil.ContextExecutor) (RawTXSlice, error) {
	var o []*RawTX

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RawTX slice")
	}

	if len(rawTXAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RawTX records in the query.
func (q rawTXQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count raw_tx rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q rawTXQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if raw_tx exists")
	}

	return count > 0, nil
}

// RawTxes retrieves all the records using an executor.
func RawTxes(mods ...qm.QueryMod) rawTXQuery {
	mods = append(mods, qm.From("\"raw_tx\""))
	return rawTXQuery{NewQuery(mods...)}
}

// FindRawTX retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRawTX(ctx context.Context, exec boil.ContextExecutor, txid models.TXID, selectCols ...string) (*RawTX, error) {
	rawTXObj := &RawTX{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"raw_tx\" where \"txid\"=$1", sel,
	)

	q := queries.Raw(query, txid)

	err := q.Bind(ctx, exec, rawTXObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from raw_tx")
	}

	return rawTXObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RawTX) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_tx provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawTXColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rawTXInsertCacheMut.RLock()
	cache, cached := rawTXInsertCache[key]
	rawTXInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rawTXAllColumns,
			rawTXColumnsWithDefault,
			rawTXColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rawTXType, rawTXMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rawTXType, rawTXMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"raw_tx\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"raw_tx\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into raw_tx")
	}

	if !cached {
		rawTXInsertCacheMut.Lock()
		rawTXInsertCache[key] = cache
		rawTXInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RawTX.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RawTX) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rawTXUpdateCacheMut.RLock()
	cache, cached := rawTXUpdateCache[key]
	rawTXUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rawTXAllColumns,
			rawTXPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update raw_tx, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"raw_tx\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rawTXPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rawTXType, rawTXMapping, append(wl, rawTXPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update raw_tx row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for raw_tx")
	}

	if !cached {
		rawTXUpdateCacheMut.Lock()
		rawTXUpdateCache[key] = cache
		rawTXUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q rawTXQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for raw_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for raw_tx")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RawTXSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"raw_tx\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rawTXPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rawTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rawTX")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RawTX) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_tx provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawTXColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	rawTXUpsertCacheMut.RLock()
	cache, cached := rawTXUpsertCache[key]
	rawTXUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rawTXAllColumns,
			rawTXColumnsWithDefault,
			rawTXColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rawTXAllColumns,
			rawTXPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert raw_tx, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rawTXPrimaryKeyColumns))
			copy(conflict, rawTXPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"raw_tx\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rawTXType, rawTXMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rawTXType, rawTXMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert raw_tx")
	}

	if !cached {
		rawTXUpsertCacheMut.Lock()
		rawTXUpsertCache[key] = cache
		rawTXUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RawTX record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RawTX) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RawTX provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rawTXPrimaryKeyMapping)
	sql := "DELETE FROM \"raw_tx\" WHERE \"txid\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from raw_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for raw_tx")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rawTXQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rawTXQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from raw_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RawTXSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(rawTXBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"raw_tx\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawTXPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rawTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_tx")
	}

	if len(rawTXAfterDeleteHooks) != 0 {
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
func (o *RawTX) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRawTX(ctx, exec, o.Txid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RawTXSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RawTXSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"raw_tx\".* FROM \"raw_tx\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawTXPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RawTXSlice")
	}

	*o = slice

	return nil
}

// RawTXExists checks if the RawTX row exists.
func RawTXExists(ctx context.Context, exec boil.ContextExecutor, txid models.TXID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"raw_tx\" where \"txid\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, txid)
	}

	row := exec.QueryRowContext(ctx, sql, txid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if raw_tx exists")
	}

	return exists, nil
}
