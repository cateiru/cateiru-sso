// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBroadcastEntries(t *testing.T) {
	t.Parallel()

	query := BroadcastEntries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBroadcastEntriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
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

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBroadcastEntriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BroadcastEntries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBroadcastEntriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BroadcastEntrySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBroadcastEntriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BroadcastEntryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BroadcastEntry exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BroadcastEntryExists to return true, but got false.")
	}
}

func testBroadcastEntriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	broadcastEntryFound, err := FindBroadcastEntry(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if broadcastEntryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBroadcastEntriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BroadcastEntries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBroadcastEntriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BroadcastEntries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBroadcastEntriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	broadcastEntryOne := &BroadcastEntry{}
	broadcastEntryTwo := &BroadcastEntry{}
	if err = randomize.Struct(seed, broadcastEntryOne, broadcastEntryDBTypes, false, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}
	if err = randomize.Struct(seed, broadcastEntryTwo, broadcastEntryDBTypes, false, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = broadcastEntryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = broadcastEntryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BroadcastEntries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBroadcastEntriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	broadcastEntryOne := &BroadcastEntry{}
	broadcastEntryTwo := &BroadcastEntry{}
	if err = randomize.Struct(seed, broadcastEntryOne, broadcastEntryDBTypes, false, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}
	if err = randomize.Struct(seed, broadcastEntryTwo, broadcastEntryDBTypes, false, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = broadcastEntryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = broadcastEntryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func broadcastEntryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func broadcastEntryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BroadcastEntry) error {
	*o = BroadcastEntry{}
	return nil
}

func testBroadcastEntriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BroadcastEntry{}
	o := &BroadcastEntry{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry object: %s", err)
	}

	AddBroadcastEntryHook(boil.BeforeInsertHook, broadcastEntryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	broadcastEntryBeforeInsertHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.AfterInsertHook, broadcastEntryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	broadcastEntryAfterInsertHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.AfterSelectHook, broadcastEntryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	broadcastEntryAfterSelectHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.BeforeUpdateHook, broadcastEntryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	broadcastEntryBeforeUpdateHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.AfterUpdateHook, broadcastEntryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	broadcastEntryAfterUpdateHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.BeforeDeleteHook, broadcastEntryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	broadcastEntryBeforeDeleteHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.AfterDeleteHook, broadcastEntryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	broadcastEntryAfterDeleteHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.BeforeUpsertHook, broadcastEntryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	broadcastEntryBeforeUpsertHooks = []BroadcastEntryHook{}

	AddBroadcastEntryHook(boil.AfterUpsertHook, broadcastEntryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	broadcastEntryAfterUpsertHooks = []BroadcastEntryHook{}
}

func testBroadcastEntriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBroadcastEntriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(broadcastEntryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBroadcastEntriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
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

func testBroadcastEntriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BroadcastEntrySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBroadcastEntriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BroadcastEntries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	broadcastEntryDBTypes = map[string]string{`ID`: `int`, `CreateUserID`: `varchar`, `Title`: `text`, `Body`: `text`, `Created`: `datetime`, `Modified`: `datetime`}
	_                     = bytes.MinRead
)

func testBroadcastEntriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(broadcastEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(broadcastEntryAllColumns) == len(broadcastEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBroadcastEntriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(broadcastEntryAllColumns) == len(broadcastEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BroadcastEntry{}
	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, broadcastEntryDBTypes, true, broadcastEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(broadcastEntryAllColumns, broadcastEntryPrimaryKeyColumns) {
		fields = broadcastEntryAllColumns
	} else {
		fields = strmangle.SetComplement(
			broadcastEntryAllColumns,
			broadcastEntryPrimaryKeyColumns,
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

	slice := BroadcastEntrySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBroadcastEntriesUpsert(t *testing.T) {
	t.Parallel()

	if len(broadcastEntryAllColumns) == len(broadcastEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBroadcastEntryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BroadcastEntry{}
	if err = randomize.Struct(seed, &o, broadcastEntryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BroadcastEntry: %s", err)
	}

	count, err := BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, broadcastEntryDBTypes, false, broadcastEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BroadcastEntry struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BroadcastEntry: %s", err)
	}

	count, err = BroadcastEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
