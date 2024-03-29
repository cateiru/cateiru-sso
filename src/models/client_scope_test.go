// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testClientScopes(t *testing.T) {
	t.Parallel()

	query := ClientScopes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClientScopesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
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

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientScopesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClientScopes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientScopesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientScopeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientScopesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClientScopeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClientScope exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClientScopeExists to return true, but got false.")
	}
}

func testClientScopesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clientScopeFound, err := FindClientScope(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if clientScopeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClientScopesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClientScopes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClientScopesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClientScopes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClientScopesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clientScopeOne := &ClientScope{}
	clientScopeTwo := &ClientScope{}
	if err = randomize.Struct(seed, clientScopeOne, clientScopeDBTypes, false, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}
	if err = randomize.Struct(seed, clientScopeTwo, clientScopeDBTypes, false, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientScopeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientScopeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientScopes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClientScopesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clientScopeOne := &ClientScope{}
	clientScopeTwo := &ClientScope{}
	if err = randomize.Struct(seed, clientScopeOne, clientScopeDBTypes, false, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}
	if err = randomize.Struct(seed, clientScopeTwo, clientScopeDBTypes, false, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientScopeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientScopeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clientScopeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func clientScopeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientScope) error {
	*o = ClientScope{}
	return nil
}

func testClientScopesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClientScope{}
	o := &ClientScope{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clientScopeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientScope object: %s", err)
	}

	AddClientScopeHook(boil.BeforeInsertHook, clientScopeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clientScopeBeforeInsertHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.AfterInsertHook, clientScopeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clientScopeAfterInsertHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.AfterSelectHook, clientScopeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clientScopeAfterSelectHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.BeforeUpdateHook, clientScopeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clientScopeBeforeUpdateHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.AfterUpdateHook, clientScopeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clientScopeAfterUpdateHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.BeforeDeleteHook, clientScopeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clientScopeBeforeDeleteHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.AfterDeleteHook, clientScopeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clientScopeAfterDeleteHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.BeforeUpsertHook, clientScopeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clientScopeBeforeUpsertHooks = []ClientScopeHook{}

	AddClientScopeHook(boil.AfterUpsertHook, clientScopeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clientScopeAfterUpsertHooks = []ClientScopeHook{}
}

func testClientScopesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientScopesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clientScopeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientScopeToOneClientUsingClient(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ClientScope
	var foreign Client

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clientScopeDBTypes, false, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ClientID = foreign.ClientID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Client().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ClientID != foreign.ClientID {
		t.Errorf("want: %v, got %v", foreign.ClientID, check.ClientID)
	}

	ranAfterSelectHook := false
	AddClientHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Client) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ClientScopeSlice{&local}
	if err = local.L.LoadClient(ctx, tx, false, (*[]*ClientScope)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Client = nil
	if err = local.L.LoadClient(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testClientScopeToOneSetOpClientUsingClient(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClientScope
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientScopeDBTypes, false, strmangle.SetComplement(clientScopePrimaryKeyColumns, clientScopeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Client{&b, &c} {
		err = a.SetClient(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Client != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClientScopes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClientID != x.ClientID {
			t.Error("foreign key was wrong value", a.ClientID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClientID))
		reflect.Indirect(reflect.ValueOf(&a.ClientID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClientID != x.ClientID {
			t.Error("foreign key was wrong value", a.ClientID, x.ClientID)
		}
	}
}

func testClientScopesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
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

func testClientScopesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientScopeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClientScopesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientScopes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clientScopeDBTypes = map[string]string{`ID`: `int`, `ClientID`: `varchar`, `Scope`: `varchar`, `CreatedAt`: `datetime`}
	_                  = bytes.MinRead
)

func testClientScopesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clientScopeAllColumns) == len(clientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClientScopesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clientScopeAllColumns) == len(clientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientScope{}
	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientScopeDBTypes, true, clientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clientScopeAllColumns, clientScopePrimaryKeyColumns) {
		fields = clientScopeAllColumns
	} else {
		fields = strmangle.SetComplement(
			clientScopeAllColumns,
			clientScopePrimaryKeyColumns,
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

	slice := ClientScopeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClientScopesUpsert(t *testing.T) {
	t.Parallel()

	if len(clientScopeAllColumns) == len(clientScopePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLClientScopeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClientScope{}
	if err = randomize.Struct(seed, &o, clientScopeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientScope: %s", err)
	}

	count, err := ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clientScopeDBTypes, false, clientScopePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientScope struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientScope: %s", err)
	}

	count, err = ClientScopes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
