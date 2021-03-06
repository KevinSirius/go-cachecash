// +build sqlboiler_test

// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRawTxes(t *testing.T) {
	t.Parallel()

	query := RawTxes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRawTxesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawTxesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RawTxes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawTxesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RawTXSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawTxesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RawTXExists(ctx, tx, o.Txid)
	if err != nil {
		t.Errorf("Unable to check if RawTX exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RawTXExists to return true, but got false.")
	}
}

func testRawTxesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	rawTXFound, err := FindRawTX(ctx, tx, o.Txid)
	if err != nil {
		t.Error(err)
	}

	if rawTXFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRawTxesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RawTxes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRawTxesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RawTxes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRawTxesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	rawTXOne := &RawTX{}
	rawTXTwo := &RawTX{}
	if err = randomize.Struct(seed, rawTXOne, rawTXDBTypes, false, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}
	if err = randomize.Struct(seed, rawTXTwo, rawTXDBTypes, false, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rawTXOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rawTXTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RawTxes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRawTxesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	rawTXOne := &RawTX{}
	rawTXTwo := &RawTX{}
	if err = randomize.Struct(seed, rawTXOne, rawTXDBTypes, false, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}
	if err = randomize.Struct(seed, rawTXTwo, rawTXDBTypes, false, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rawTXOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rawTXTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func rawTXBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func rawTXAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RawTX) error {
	*o = RawTX{}
	return nil
}

func testRawTxesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RawTX{}
	o := &RawTX{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, rawTXDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RawTX object: %s", err)
	}

	AddRawTXHook(boil.BeforeInsertHook, rawTXBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	rawTXBeforeInsertHooks = []RawTXHook{}

	AddRawTXHook(boil.AfterInsertHook, rawTXAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	rawTXAfterInsertHooks = []RawTXHook{}

	AddRawTXHook(boil.AfterSelectHook, rawTXAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	rawTXAfterSelectHooks = []RawTXHook{}

	AddRawTXHook(boil.BeforeUpdateHook, rawTXBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	rawTXBeforeUpdateHooks = []RawTXHook{}

	AddRawTXHook(boil.AfterUpdateHook, rawTXAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	rawTXAfterUpdateHooks = []RawTXHook{}

	AddRawTXHook(boil.BeforeDeleteHook, rawTXBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	rawTXBeforeDeleteHooks = []RawTXHook{}

	AddRawTXHook(boil.AfterDeleteHook, rawTXAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	rawTXAfterDeleteHooks = []RawTXHook{}

	AddRawTXHook(boil.BeforeUpsertHook, rawTXBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	rawTXBeforeUpsertHooks = []RawTXHook{}

	AddRawTXHook(boil.AfterUpsertHook, rawTXAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	rawTXAfterUpsertHooks = []RawTXHook{}
}

func testRawTxesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRawTxesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(rawTXColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRawTxesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRawTxesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RawTXSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRawTxesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RawTxes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	rawTXDBTypes = map[string]string{`Txid`: `bytea`, `Bytes`: `bytea`}
	_            = bytes.MinRead
)

func testRawTxesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rawTXPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(rawTXAllColumns) == len(rawTXPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRawTxesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(rawTXAllColumns) == len(rawTXPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RawTX{}
	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rawTXDBTypes, true, rawTXPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(rawTXAllColumns, rawTXPrimaryKeyColumns) {
		fields = rawTXAllColumns
	} else {
		fields = strmangle.SetComplement(
			rawTXAllColumns,
			rawTXPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RawTXSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRawTxesUpsert(t *testing.T) {
	t.Parallel()

	if len(rawTXAllColumns) == len(rawTXPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RawTX{}
	if err = randomize.Struct(seed, &o, rawTXDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RawTX: %s", err)
	}

	count, err := RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, rawTXDBTypes, false, rawTXPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawTX struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RawTX: %s", err)
	}

	count, err = RawTxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
