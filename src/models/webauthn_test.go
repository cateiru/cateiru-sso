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

func testWebauthns(t *testing.T) {
	t.Parallel()

	query := Webauthns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWebauthnsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
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

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebauthnsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Webauthns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebauthnsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebauthnSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebauthnsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WebauthnExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Webauthn exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WebauthnExists to return true, but got false.")
	}
}

func testWebauthnsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	webauthnFound, err := FindWebauthn(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if webauthnFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWebauthnsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Webauthns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWebauthnsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Webauthns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWebauthnsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	webauthnOne := &Webauthn{}
	webauthnTwo := &Webauthn{}
	if err = randomize.Struct(seed, webauthnOne, webauthnDBTypes, false, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}
	if err = randomize.Struct(seed, webauthnTwo, webauthnDBTypes, false, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = webauthnOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = webauthnTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Webauthns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWebauthnsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	webauthnOne := &Webauthn{}
	webauthnTwo := &Webauthn{}
	if err = randomize.Struct(seed, webauthnOne, webauthnDBTypes, false, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}
	if err = randomize.Struct(seed, webauthnTwo, webauthnDBTypes, false, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = webauthnOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = webauthnTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func webauthnBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func webauthnAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Webauthn) error {
	*o = Webauthn{}
	return nil
}

func testWebauthnsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Webauthn{}
	o := &Webauthn{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, webauthnDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Webauthn object: %s", err)
	}

	AddWebauthnHook(boil.BeforeInsertHook, webauthnBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	webauthnBeforeInsertHooks = []WebauthnHook{}

	AddWebauthnHook(boil.AfterInsertHook, webauthnAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	webauthnAfterInsertHooks = []WebauthnHook{}

	AddWebauthnHook(boil.AfterSelectHook, webauthnAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	webauthnAfterSelectHooks = []WebauthnHook{}

	AddWebauthnHook(boil.BeforeUpdateHook, webauthnBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	webauthnBeforeUpdateHooks = []WebauthnHook{}

	AddWebauthnHook(boil.AfterUpdateHook, webauthnAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	webauthnAfterUpdateHooks = []WebauthnHook{}

	AddWebauthnHook(boil.BeforeDeleteHook, webauthnBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	webauthnBeforeDeleteHooks = []WebauthnHook{}

	AddWebauthnHook(boil.AfterDeleteHook, webauthnAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	webauthnAfterDeleteHooks = []WebauthnHook{}

	AddWebauthnHook(boil.BeforeUpsertHook, webauthnBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	webauthnBeforeUpsertHooks = []WebauthnHook{}

	AddWebauthnHook(boil.AfterUpsertHook, webauthnAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	webauthnAfterUpsertHooks = []WebauthnHook{}
}

func testWebauthnsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebauthnsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(webauthnColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebauthnsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
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

func testWebauthnsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebauthnSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWebauthnsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Webauthns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	webauthnDBTypes = map[string]string{`ID`: `bigint`, `UserID`: `varchar`, `Credential`: `json`, `Device`: `varchar`, `Os`: `varchar`, `Browser`: `varchar`, `IsMobile`: `tinyint`, `IP`: `varbinary`, `Created`: `datetime`}
	_               = bytes.MinRead
)

func testWebauthnsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(webauthnPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(webauthnAllColumns) == len(webauthnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWebauthnsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(webauthnAllColumns) == len(webauthnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Webauthn{}
	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, webauthnDBTypes, true, webauthnPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(webauthnAllColumns, webauthnPrimaryKeyColumns) {
		fields = webauthnAllColumns
	} else {
		fields = strmangle.SetComplement(
			webauthnAllColumns,
			webauthnPrimaryKeyColumns,
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

	slice := WebauthnSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWebauthnsUpsert(t *testing.T) {
	t.Parallel()

	if len(webauthnAllColumns) == len(webauthnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLWebauthnUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Webauthn{}
	if err = randomize.Struct(seed, &o, webauthnDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Webauthn: %s", err)
	}

	count, err := Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, webauthnDBTypes, false, webauthnPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Webauthn struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Webauthn: %s", err)
	}

	count, err = Webauthns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}