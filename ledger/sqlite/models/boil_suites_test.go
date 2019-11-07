// +build sqlboiler_test

// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("RawBlocks", testRawBlocks)
	t.Run("RawTxes", testRawTxes)
}

func TestDelete(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksDelete)
	t.Run("RawTxes", testRawTxesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksQueryDeleteAll)
	t.Run("RawTxes", testRawTxesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksSliceDeleteAll)
	t.Run("RawTxes", testRawTxesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksExists)
	t.Run("RawTxes", testRawTxesExists)
}

func TestFind(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksFind)
	t.Run("RawTxes", testRawTxesFind)
}

func TestBind(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksBind)
	t.Run("RawTxes", testRawTxesBind)
}

func TestOne(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksOne)
	t.Run("RawTxes", testRawTxesOne)
}

func TestAll(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksAll)
	t.Run("RawTxes", testRawTxesAll)
}

func TestCount(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksCount)
	t.Run("RawTxes", testRawTxesCount)
}

func TestHooks(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksHooks)
	t.Run("RawTxes", testRawTxesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksInsert)
	t.Run("RawBlocks", testRawBlocksInsertWhitelist)
	t.Run("RawTxes", testRawTxesInsert)
	t.Run("RawTxes", testRawTxesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksReload)
	t.Run("RawTxes", testRawTxesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksReloadAll)
	t.Run("RawTxes", testRawTxesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksSelect)
	t.Run("RawTxes", testRawTxesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksUpdate)
	t.Run("RawTxes", testRawTxesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("RawBlocks", testRawBlocksSliceUpdateAll)
	t.Run("RawTxes", testRawTxesSliceUpdateAll)
}
