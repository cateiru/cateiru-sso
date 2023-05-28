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

func testRegisterSessions(t *testing.T) {
	t.Parallel()

	query := RegisterSessions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRegisterSessionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
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

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterSessionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RegisterSessions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterSessionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegisterSessionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterSessionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RegisterSessionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RegisterSession exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RegisterSessionExists to return true, but got false.")
	}
}

func testRegisterSessionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	registerSessionFound, err := FindRegisterSession(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if registerSessionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRegisterSessionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RegisterSessions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRegisterSessionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RegisterSessions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRegisterSessionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	registerSessionOne := &RegisterSession{}
	registerSessionTwo := &RegisterSession{}
	if err = randomize.Struct(seed, registerSessionOne, registerSessionDBTypes, false, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}
	if err = randomize.Struct(seed, registerSessionTwo, registerSessionDBTypes, false, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registerSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registerSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegisterSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRegisterSessionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	registerSessionOne := &RegisterSession{}
	registerSessionTwo := &RegisterSession{}
	if err = randomize.Struct(seed, registerSessionOne, registerSessionDBTypes, false, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}
	if err = randomize.Struct(seed, registerSessionTwo, registerSessionDBTypes, false, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registerSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registerSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func registerSessionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func registerSessionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterSession) error {
	*o = RegisterSession{}
	return nil
}

func testRegisterSessionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RegisterSession{}
	o := &RegisterSession{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, registerSessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegisterSession object: %s", err)
	}

	AddRegisterSessionHook(boil.BeforeInsertHook, registerSessionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	registerSessionBeforeInsertHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.AfterInsertHook, registerSessionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	registerSessionAfterInsertHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.AfterSelectHook, registerSessionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	registerSessionAfterSelectHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.BeforeUpdateHook, registerSessionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	registerSessionBeforeUpdateHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.AfterUpdateHook, registerSessionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	registerSessionAfterUpdateHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.BeforeDeleteHook, registerSessionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	registerSessionBeforeDeleteHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.AfterDeleteHook, registerSessionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	registerSessionAfterDeleteHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.BeforeUpsertHook, registerSessionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	registerSessionBeforeUpsertHooks = []RegisterSessionHook{}

	AddRegisterSessionHook(boil.AfterUpsertHook, registerSessionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	registerSessionAfterUpsertHooks = []RegisterSessionHook{}
}

func testRegisterSessionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegisterSessionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(registerSessionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegisterSessionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
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

func testRegisterSessionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegisterSessionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegisterSessionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegisterSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	registerSessionDBTypes = map[string]string{`ID`: `varchar`, `Email`: `varchar`, `EmailVerified`: `tinyint`, `SendCount`: `tinyint`, `VerifyCode`: `char`, `RetryCount`: `tinyint`, `Period`: `datetime`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                      = bytes.MinRead
)

func testRegisterSessionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(registerSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(registerSessionAllColumns) == len(registerSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRegisterSessionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(registerSessionAllColumns) == len(registerSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegisterSession{}
	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registerSessionDBTypes, true, registerSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(registerSessionAllColumns, registerSessionPrimaryKeyColumns) {
		fields = registerSessionAllColumns
	} else {
		fields = strmangle.SetComplement(
			registerSessionAllColumns,
			registerSessionPrimaryKeyColumns,
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

	slice := RegisterSessionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRegisterSessionsUpsert(t *testing.T) {
	t.Parallel()

	if len(registerSessionAllColumns) == len(registerSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRegisterSessionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RegisterSession{}
	if err = randomize.Struct(seed, &o, registerSessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegisterSession: %s", err)
	}

	count, err := RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, registerSessionDBTypes, false, registerSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterSession struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegisterSession: %s", err)
	}

	count, err = RegisterSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
