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

func testClientRefreshes(t *testing.T) {
	t.Parallel()

	query := ClientRefreshes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClientRefreshesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
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

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRefreshesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClientRefreshes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRefreshesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientRefreshSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRefreshesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClientRefreshExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClientRefresh exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClientRefreshExists to return true, but got false.")
	}
}

func testClientRefreshesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clientRefreshFound, err := FindClientRefresh(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if clientRefreshFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClientRefreshesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClientRefreshes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClientRefreshesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClientRefreshes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClientRefreshesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clientRefreshOne := &ClientRefresh{}
	clientRefreshTwo := &ClientRefresh{}
	if err = randomize.Struct(seed, clientRefreshOne, clientRefreshDBTypes, false, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}
	if err = randomize.Struct(seed, clientRefreshTwo, clientRefreshDBTypes, false, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientRefreshOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientRefreshTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientRefreshes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClientRefreshesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clientRefreshOne := &ClientRefresh{}
	clientRefreshTwo := &ClientRefresh{}
	if err = randomize.Struct(seed, clientRefreshOne, clientRefreshDBTypes, false, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}
	if err = randomize.Struct(seed, clientRefreshTwo, clientRefreshDBTypes, false, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientRefreshOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientRefreshTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clientRefreshBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func clientRefreshAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRefresh) error {
	*o = ClientRefresh{}
	return nil
}

func testClientRefreshesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClientRefresh{}
	o := &ClientRefresh{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientRefresh object: %s", err)
	}

	AddClientRefreshHook(boil.BeforeInsertHook, clientRefreshBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clientRefreshBeforeInsertHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.AfterInsertHook, clientRefreshAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clientRefreshAfterInsertHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.AfterSelectHook, clientRefreshAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clientRefreshAfterSelectHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.BeforeUpdateHook, clientRefreshBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clientRefreshBeforeUpdateHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.AfterUpdateHook, clientRefreshAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clientRefreshAfterUpdateHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.BeforeDeleteHook, clientRefreshBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clientRefreshBeforeDeleteHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.AfterDeleteHook, clientRefreshAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clientRefreshAfterDeleteHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.BeforeUpsertHook, clientRefreshBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clientRefreshBeforeUpsertHooks = []ClientRefreshHook{}

	AddClientRefreshHook(boil.AfterUpsertHook, clientRefreshAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clientRefreshAfterUpsertHooks = []ClientRefreshHook{}
}

func testClientRefreshesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientRefreshesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clientRefreshColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientRefreshesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
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

func testClientRefreshesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientRefreshSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClientRefreshesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientRefreshes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clientRefreshDBTypes = map[string]string{`ID`: `varchar`, `UserID`: `varchar`, `ClientID`: `varchar`, `Scopes`: `json`, `SessionID`: `varchar`, `Period`: `datetime`, `Created`: `datetime`, `Modified`: `datetime`}
	_                    = bytes.MinRead
)

func testClientRefreshesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clientRefreshPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clientRefreshAllColumns) == len(clientRefreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClientRefreshesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clientRefreshAllColumns) == len(clientRefreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientRefresh{}
	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientRefreshDBTypes, true, clientRefreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clientRefreshAllColumns, clientRefreshPrimaryKeyColumns) {
		fields = clientRefreshAllColumns
	} else {
		fields = strmangle.SetComplement(
			clientRefreshAllColumns,
			clientRefreshPrimaryKeyColumns,
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

	slice := ClientRefreshSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClientRefreshesUpsert(t *testing.T) {
	t.Parallel()

	if len(clientRefreshAllColumns) == len(clientRefreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLClientRefreshUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClientRefresh{}
	if err = randomize.Struct(seed, &o, clientRefreshDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientRefresh: %s", err)
	}

	count, err := ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clientRefreshDBTypes, false, clientRefreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRefresh struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientRefresh: %s", err)
	}

	count, err = ClientRefreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
