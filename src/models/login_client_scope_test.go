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

func testLoginClientScopes(t *testing.T) {
	t.Parallel()

	query := LoginClientScopes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLoginClientScopesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
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

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientScopesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LoginClientScopes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientScopesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginClientScopeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientScopesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LoginClientScopeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LoginClientScope exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LoginClientScopeExists to return true, but got false.")
	}
}

func testLoginClientScopesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	loginClientScopeFound, err := FindLoginClientScope(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if loginClientScopeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLoginClientScopesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LoginClientScopes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLoginClientScopesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LoginClientScopes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLoginClientScopesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	loginClientScopeOne := &LoginClientScope{}
	loginClientScopeTwo := &LoginClientScope{}
	if err = randomize.Struct(seed, loginClientScopeOne, loginClientScopeDBTypes, false, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}
	if err = randomize.Struct(seed, loginClientScopeTwo, loginClientScopeDBTypes, false, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginClientScopeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginClientScopeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginClientScopes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLoginClientScopesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	loginClientScopeOne := &LoginClientScope{}
	loginClientScopeTwo := &LoginClientScope{}
	if err = randomize.Struct(seed, loginClientScopeOne, loginClientScopeDBTypes, false, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}
	if err = randomize.Struct(seed, loginClientScopeTwo, loginClientScopeDBTypes, false, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginClientScopeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginClientScopeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func loginClientScopeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func loginClientScopeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClientScope) error {
	*o = LoginClientScope{}
	return nil
}

func testLoginClientScopesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LoginClientScope{}
	o := &LoginClientScope{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginClientScope object: %s", err)
	}

	AddLoginClientScopeHook(boil.BeforeInsertHook, loginClientScopeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	loginClientScopeBeforeInsertHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.AfterInsertHook, loginClientScopeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	loginClientScopeAfterInsertHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.AfterSelectHook, loginClientScopeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	loginClientScopeAfterSelectHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.BeforeUpdateHook, loginClientScopeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	loginClientScopeBeforeUpdateHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.AfterUpdateHook, loginClientScopeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	loginClientScopeAfterUpdateHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.BeforeDeleteHook, loginClientScopeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	loginClientScopeBeforeDeleteHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.AfterDeleteHook, loginClientScopeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	loginClientScopeAfterDeleteHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.BeforeUpsertHook, loginClientScopeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	loginClientScopeBeforeUpsertHooks = []LoginClientScopeHook{}

	AddLoginClientScopeHook(boil.AfterUpsertHook, loginClientScopeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	loginClientScopeAfterUpsertHooks = []LoginClientScopeHook{}
}

func testLoginClientScopesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginClientScopesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(loginClientScopeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginClientScopesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
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

func testLoginClientScopesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginClientScopeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLoginClientScopesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginClientScopes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	loginClientScopeDBTypes = map[string]string{`ID`: `int`, `LoginClientID`: `int`, `Scope`: `varchar`, `Created`: `datetime`}
	_                       = bytes.MinRead
)

func testLoginClientScopesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(loginClientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(loginClientScopeAllColumns) == len(loginClientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLoginClientScopesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(loginClientScopeAllColumns) == len(loginClientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginClientScope{}
	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginClientScopeDBTypes, true, loginClientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(loginClientScopeAllColumns, loginClientScopePrimaryKeyColumns) {
		fields = loginClientScopeAllColumns
	} else {
		fields = strmangle.SetComplement(
			loginClientScopeAllColumns,
			loginClientScopePrimaryKeyColumns,
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

	slice := LoginClientScopeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLoginClientScopesUpsert(t *testing.T) {
	t.Parallel()

	if len(loginClientScopeAllColumns) == len(loginClientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLLoginClientScopeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LoginClientScope{}
	if err = randomize.Struct(seed, &o, loginClientScopeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginClientScope: %s", err)
	}

	count, err := LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, loginClientScopeDBTypes, false, loginClientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClientScope struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginClientScope: %s", err)
	}

	count, err = LoginClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}