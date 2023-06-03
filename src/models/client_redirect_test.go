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

func testClientRedirects(t *testing.T) {
	t.Parallel()

	query := ClientRedirects()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClientRedirectsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
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

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRedirectsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClientRedirects().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRedirectsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientRedirectSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientRedirectsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClientRedirectExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClientRedirect exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClientRedirectExists to return true, but got false.")
	}
}

func testClientRedirectsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clientRedirectFound, err := FindClientRedirect(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if clientRedirectFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClientRedirectsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClientRedirects().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClientRedirectsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClientRedirects().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClientRedirectsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clientRedirectOne := &ClientRedirect{}
	clientRedirectTwo := &ClientRedirect{}
	if err = randomize.Struct(seed, clientRedirectOne, clientRedirectDBTypes, false, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}
	if err = randomize.Struct(seed, clientRedirectTwo, clientRedirectDBTypes, false, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientRedirectOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientRedirectTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientRedirects().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClientRedirectsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clientRedirectOne := &ClientRedirect{}
	clientRedirectTwo := &ClientRedirect{}
	if err = randomize.Struct(seed, clientRedirectOne, clientRedirectDBTypes, false, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}
	if err = randomize.Struct(seed, clientRedirectTwo, clientRedirectDBTypes, false, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientRedirectOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientRedirectTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clientRedirectBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func clientRedirectAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientRedirect) error {
	*o = ClientRedirect{}
	return nil
}

func testClientRedirectsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClientRedirect{}
	o := &ClientRedirect{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientRedirect object: %s", err)
	}

	AddClientRedirectHook(boil.BeforeInsertHook, clientRedirectBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clientRedirectBeforeInsertHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.AfterInsertHook, clientRedirectAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clientRedirectAfterInsertHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.AfterSelectHook, clientRedirectAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clientRedirectAfterSelectHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.BeforeUpdateHook, clientRedirectBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clientRedirectBeforeUpdateHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.AfterUpdateHook, clientRedirectAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clientRedirectAfterUpdateHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.BeforeDeleteHook, clientRedirectBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clientRedirectBeforeDeleteHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.AfterDeleteHook, clientRedirectAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clientRedirectAfterDeleteHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.BeforeUpsertHook, clientRedirectBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clientRedirectBeforeUpsertHooks = []ClientRedirectHook{}

	AddClientRedirectHook(boil.AfterUpsertHook, clientRedirectAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clientRedirectAfterUpsertHooks = []ClientRedirectHook{}
}

func testClientRedirectsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientRedirectsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clientRedirectColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientRedirectToOneClientUsingClient(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ClientRedirect
	var foreign Client

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clientRedirectDBTypes, false, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
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

	slice := ClientRedirectSlice{&local}
	if err = local.L.LoadClient(ctx, tx, false, (*[]*ClientRedirect)(&slice), nil); err != nil {
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

func testClientRedirectToOneSetOpClientUsingClient(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClientRedirect
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientRedirectDBTypes, false, strmangle.SetComplement(clientRedirectPrimaryKeyColumns, clientRedirectColumnsWithoutDefault)...); err != nil {
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

		if x.R.ClientRedirects[0] != &a {
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

func testClientRedirectsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
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

func testClientRedirectsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientRedirectSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClientRedirectsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientRedirects().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clientRedirectDBTypes = map[string]string{`ID`: `int`, `ClientID`: `varchar`, `URL`: `text`, `CreatedAt`: `datetime`}
	_                     = bytes.MinRead
)

func testClientRedirectsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clientRedirectPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clientRedirectAllColumns) == len(clientRedirectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClientRedirectsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clientRedirectAllColumns) == len(clientRedirectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientRedirect{}
	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientRedirectDBTypes, true, clientRedirectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clientRedirectAllColumns, clientRedirectPrimaryKeyColumns) {
		fields = clientRedirectAllColumns
	} else {
		fields = strmangle.SetComplement(
			clientRedirectAllColumns,
			clientRedirectPrimaryKeyColumns,
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

	slice := ClientRedirectSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClientRedirectsUpsert(t *testing.T) {
	t.Parallel()

	if len(clientRedirectAllColumns) == len(clientRedirectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLClientRedirectUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClientRedirect{}
	if err = randomize.Struct(seed, &o, clientRedirectDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientRedirect: %s", err)
	}

	count, err := ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clientRedirectDBTypes, false, clientRedirectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientRedirect struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientRedirect: %s", err)
	}

	count, err = ClientRedirects().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
