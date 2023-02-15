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

func testLoginClients(t *testing.T) {
	t.Parallel()

	query := LoginClients()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLoginClientsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
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

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LoginClients().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginClientSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginClientsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LoginClientExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LoginClient exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LoginClientExists to return true, but got false.")
	}
}

func testLoginClientsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	loginClientFound, err := FindLoginClient(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if loginClientFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLoginClientsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LoginClients().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLoginClientsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LoginClients().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLoginClientsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	loginClientOne := &LoginClient{}
	loginClientTwo := &LoginClient{}
	if err = randomize.Struct(seed, loginClientOne, loginClientDBTypes, false, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}
	if err = randomize.Struct(seed, loginClientTwo, loginClientDBTypes, false, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginClientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginClientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginClients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLoginClientsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	loginClientOne := &LoginClient{}
	loginClientTwo := &LoginClient{}
	if err = randomize.Struct(seed, loginClientOne, loginClientDBTypes, false, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}
	if err = randomize.Struct(seed, loginClientTwo, loginClientDBTypes, false, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginClientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginClientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func loginClientBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func loginClientAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginClient) error {
	*o = LoginClient{}
	return nil
}

func testLoginClientsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LoginClient{}
	o := &LoginClient{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, loginClientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginClient object: %s", err)
	}

	AddLoginClientHook(boil.BeforeInsertHook, loginClientBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	loginClientBeforeInsertHooks = []LoginClientHook{}

	AddLoginClientHook(boil.AfterInsertHook, loginClientAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	loginClientAfterInsertHooks = []LoginClientHook{}

	AddLoginClientHook(boil.AfterSelectHook, loginClientAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	loginClientAfterSelectHooks = []LoginClientHook{}

	AddLoginClientHook(boil.BeforeUpdateHook, loginClientBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	loginClientBeforeUpdateHooks = []LoginClientHook{}

	AddLoginClientHook(boil.AfterUpdateHook, loginClientAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	loginClientAfterUpdateHooks = []LoginClientHook{}

	AddLoginClientHook(boil.BeforeDeleteHook, loginClientBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	loginClientBeforeDeleteHooks = []LoginClientHook{}

	AddLoginClientHook(boil.AfterDeleteHook, loginClientAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	loginClientAfterDeleteHooks = []LoginClientHook{}

	AddLoginClientHook(boil.BeforeUpsertHook, loginClientBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	loginClientBeforeUpsertHooks = []LoginClientHook{}

	AddLoginClientHook(boil.AfterUpsertHook, loginClientAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	loginClientAfterUpsertHooks = []LoginClientHook{}
}

func testLoginClientsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginClientsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(loginClientColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginClientsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
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

func testLoginClientsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginClientSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLoginClientsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginClients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	loginClientDBTypes = map[string]string{`ID`: `int`, `ClientID`: `varbinary`, `UserID`: `varbinary`, `Created`: `datetime`}
	_                  = bytes.MinRead
)

func testLoginClientsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(loginClientPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(loginClientAllColumns) == len(loginClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLoginClientsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(loginClientAllColumns) == len(loginClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginClient{}
	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginClientDBTypes, true, loginClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(loginClientAllColumns, loginClientPrimaryKeyColumns) {
		fields = loginClientAllColumns
	} else {
		fields = strmangle.SetComplement(
			loginClientAllColumns,
			loginClientPrimaryKeyColumns,
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

	slice := LoginClientSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLoginClientsUpsert(t *testing.T) {
	t.Parallel()

	if len(loginClientAllColumns) == len(loginClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLLoginClientUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LoginClient{}
	if err = randomize.Struct(seed, &o, loginClientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginClient: %s", err)
	}

	count, err := LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, loginClientDBTypes, false, loginClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginClient struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginClient: %s", err)
	}

	count, err = LoginClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}