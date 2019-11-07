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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RawBlock is an object representing the database table.
type RawBlock struct {
	Blockid []byte `boil:"blockid" json:"blockid" toml:"blockid" yaml:"blockid"`
	Height  int64  `boil:"height" json:"height" toml:"height" yaml:"height"`
	Bytes   []byte `boil:"bytes" json:"bytes" toml:"bytes" yaml:"bytes"`

	R *rawBlockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rawBlockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RawBlockColumns = struct {
	Blockid string
	Height  string
	Bytes   string
}{
	Blockid: "blockid",
	Height:  "height",
	Bytes:   "bytes",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var RawBlockWhere = struct {
	Blockid whereHelper__byte
	Height  whereHelperint64
	Bytes   whereHelper__byte
}{
	Blockid: whereHelper__byte{field: "\"raw_block\".\"blockid\""},
	Height:  whereHelperint64{field: "\"raw_block\".\"height\""},
	Bytes:   whereHelper__byte{field: "\"raw_block\".\"bytes\""},
}

// RawBlockRels is where relationship names are stored.
var RawBlockRels = struct {
}{}

// rawBlockR is where relationships are stored.
type rawBlockR struct {
}

// NewStruct creates a new relationship struct
func (*rawBlockR) NewStruct() *rawBlockR {
	return &rawBlockR{}
}

// rawBlockL is where Load methods for each relationship are stored.
type rawBlockL struct{}

var (
	rawBlockAllColumns            = []string{"blockid", "height", "bytes"}
	rawBlockColumnsWithoutDefault = []string{"blockid", "height", "bytes"}
	rawBlockColumnsWithDefault    = []string{}
	rawBlockPrimaryKeyColumns     = []string{"blockid"}
)

type (
	// RawBlockSlice is an alias for a slice of pointers to RawBlock.
	// This should generally be used opposed to []RawBlock.
	RawBlockSlice []*RawBlock
	// RawBlockHook is the signature for custom RawBlock hook methods
	RawBlockHook func(context.Context, boil.ContextExecutor, *RawBlock) error

	rawBlockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rawBlockType                 = reflect.TypeOf(&RawBlock{})
	rawBlockMapping              = queries.MakeStructMapping(rawBlockType)
	rawBlockPrimaryKeyMapping, _ = queries.BindMapping(rawBlockType, rawBlockMapping, rawBlockPrimaryKeyColumns)
	rawBlockInsertCacheMut       sync.RWMutex
	rawBlockInsertCache          = make(map[string]insertCache)
	rawBlockUpdateCacheMut       sync.RWMutex
	rawBlockUpdateCache          = make(map[string]updateCache)
	rawBlockUpsertCacheMut       sync.RWMutex
	rawBlockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rawBlockBeforeInsertHooks []RawBlockHook
var rawBlockBeforeUpdateHooks []RawBlockHook
var rawBlockBeforeDeleteHooks []RawBlockHook
var rawBlockBeforeUpsertHooks []RawBlockHook

var rawBlockAfterInsertHooks []RawBlockHook
var rawBlockAfterSelectHooks []RawBlockHook
var rawBlockAfterUpdateHooks []RawBlockHook
var rawBlockAfterDeleteHooks []RawBlockHook
var rawBlockAfterUpsertHooks []RawBlockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RawBlock) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RawBlock) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RawBlock) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RawBlock) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RawBlock) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RawBlock) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RawBlock) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RawBlock) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RawBlock) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawBlockAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRawBlockHook registers your hook function for all future operations.
func AddRawBlockHook(hookPoint boil.HookPoint, rawBlockHook RawBlockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rawBlockBeforeInsertHooks = append(rawBlockBeforeInsertHooks, rawBlockHook)
	case boil.BeforeUpdateHook:
		rawBlockBeforeUpdateHooks = append(rawBlockBeforeUpdateHooks, rawBlockHook)
	case boil.BeforeDeleteHook:
		rawBlockBeforeDeleteHooks = append(rawBlockBeforeDeleteHooks, rawBlockHook)
	case boil.BeforeUpsertHook:
		rawBlockBeforeUpsertHooks = append(rawBlockBeforeUpsertHooks, rawBlockHook)
	case boil.AfterInsertHook:
		rawBlockAfterInsertHooks = append(rawBlockAfterInsertHooks, rawBlockHook)
	case boil.AfterSelectHook:
		rawBlockAfterSelectHooks = append(rawBlockAfterSelectHooks, rawBlockHook)
	case boil.AfterUpdateHook:
		rawBlockAfterUpdateHooks = append(rawBlockAfterUpdateHooks, rawBlockHook)
	case boil.AfterDeleteHook:
		rawBlockAfterDeleteHooks = append(rawBlockAfterDeleteHooks, rawBlockHook)
	case boil.AfterUpsertHook:
		rawBlockAfterUpsertHooks = append(rawBlockAfterUpsertHooks, rawBlockHook)
	}
}

// One returns a single rawBlock record from the query.
func (q rawBlockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RawBlock, error) {
	o := &RawBlock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for raw_block")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RawBlock records from the query.
func (q rawBlockQuery) All(ctx context.Context, exec boil.ContextExecutor) (RawBlockSlice, error) {
	var o []*RawBlock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RawBlock slice")
	}

	if len(rawBlockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RawBlock records in the query.
func (q rawBlockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count raw_block rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q rawBlockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if raw_block exists")
	}

	return count > 0, nil
}

// RawBlocks retrieves all the records using an executor.
func RawBlocks(mods ...qm.QueryMod) rawBlockQuery {
	mods = append(mods, qm.From("\"raw_block\""))
	return rawBlockQuery{NewQuery(mods...)}
}

// FindRawBlock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRawBlock(ctx context.Context, exec boil.ContextExecutor, blockid []byte, selectCols ...string) (*RawBlock, error) {
	rawBlockObj := &RawBlock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"raw_block\" where \"blockid\"=$1", sel,
	)

	q := queries.Raw(query, blockid)

	err := q.Bind(ctx, exec, rawBlockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from raw_block")
	}

	return rawBlockObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RawBlock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_block provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawBlockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rawBlockInsertCacheMut.RLock()
	cache, cached := rawBlockInsertCache[key]
	rawBlockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rawBlockAllColumns,
			rawBlockColumnsWithDefault,
			rawBlockColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rawBlockType, rawBlockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rawBlockType, rawBlockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"raw_block\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"raw_block\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into raw_block")
	}

	if !cached {
		rawBlockInsertCacheMut.Lock()
		rawBlockInsertCache[key] = cache
		rawBlockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RawBlock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RawBlock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rawBlockUpdateCacheMut.RLock()
	cache, cached := rawBlockUpdateCache[key]
	rawBlockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rawBlockAllColumns,
			rawBlockPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update raw_block, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"raw_block\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rawBlockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rawBlockType, rawBlockMapping, append(wl, rawBlockPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update raw_block row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for raw_block")
	}

	if !cached {
		rawBlockUpdateCacheMut.Lock()
		rawBlockUpdateCache[key] = cache
		rawBlockUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q rawBlockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for raw_block")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for raw_block")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RawBlockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"raw_block\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rawBlockPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rawBlock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rawBlock")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RawBlock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_block provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawBlockColumnsWithDefault, o)

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

	rawBlockUpsertCacheMut.RLock()
	cache, cached := rawBlockUpsertCache[key]
	rawBlockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rawBlockAllColumns,
			rawBlockColumnsWithDefault,
			rawBlockColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rawBlockAllColumns,
			rawBlockPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert raw_block, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rawBlockPrimaryKeyColumns))
			copy(conflict, rawBlockPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"raw_block\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rawBlockType, rawBlockMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rawBlockType, rawBlockMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert raw_block")
	}

	if !cached {
		rawBlockUpsertCacheMut.Lock()
		rawBlockUpsertCache[key] = cache
		rawBlockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RawBlock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RawBlock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RawBlock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rawBlockPrimaryKeyMapping)
	sql := "DELETE FROM \"raw_block\" WHERE \"blockid\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from raw_block")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for raw_block")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rawBlockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rawBlockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from raw_block")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_block")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RawBlockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(rawBlockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"raw_block\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawBlockPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rawBlock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_block")
	}

	if len(rawBlockAfterDeleteHooks) != 0 {
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
func (o *RawBlock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRawBlock(ctx, exec, o.Blockid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RawBlockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RawBlockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"raw_block\".* FROM \"raw_block\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawBlockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RawBlockSlice")
	}

	*o = slice

	return nil
}

// RawBlockExists checks if the RawBlock row exists.
func RawBlockExists(ctx context.Context, exec boil.ContextExecutor, blockid []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"raw_block\" where \"blockid\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, blockid)
	}

	row := exec.QueryRowContext(ctx, sql, blockid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if raw_block exists")
	}

	return exists, nil
}
