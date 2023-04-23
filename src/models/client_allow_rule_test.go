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

func testClientAllowRules(t *testing.T) {
	t.Parallel()

	query := ClientAllowRules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClientAllowRulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
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

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientAllowRulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClientAllowRules().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientAllowRulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientAllowRuleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientAllowRulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClientAllowRuleExists(ctx, tx, o.ClientID)
	if err != nil {
		t.Errorf("Unable to check if ClientAllowRule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClientAllowRuleExists to return true, but got false.")
	}
}

func testClientAllowRulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clientAllowRuleFound, err := FindClientAllowRule(ctx, tx, o.ClientID)
	if err != nil {
		t.Error(err)
	}

	if clientAllowRuleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClientAllowRulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClientAllowRules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClientAllowRulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClientAllowRules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClientAllowRulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clientAllowRuleOne := &ClientAllowRule{}
	clientAllowRuleTwo := &ClientAllowRule{}
	if err = randomize.Struct(seed, clientAllowRuleOne, clientAllowRuleDBTypes, false, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}
	if err = randomize.Struct(seed, clientAllowRuleTwo, clientAllowRuleDBTypes, false, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientAllowRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientAllowRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientAllowRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClientAllowRulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clientAllowRuleOne := &ClientAllowRule{}
	clientAllowRuleTwo := &ClientAllowRule{}
	if err = randomize.Struct(seed, clientAllowRuleOne, clientAllowRuleDBTypes, false, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}
	if err = randomize.Struct(seed, clientAllowRuleTwo, clientAllowRuleDBTypes, false, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clientAllowRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clientAllowRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clientAllowRuleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func clientAllowRuleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClientAllowRule) error {
	*o = ClientAllowRule{}
	return nil
}

func testClientAllowRulesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClientAllowRule{}
	o := &ClientAllowRule{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule object: %s", err)
	}

	AddClientAllowRuleHook(boil.BeforeInsertHook, clientAllowRuleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleBeforeInsertHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.AfterInsertHook, clientAllowRuleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleAfterInsertHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.AfterSelectHook, clientAllowRuleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleAfterSelectHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.BeforeUpdateHook, clientAllowRuleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleBeforeUpdateHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.AfterUpdateHook, clientAllowRuleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleAfterUpdateHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.BeforeDeleteHook, clientAllowRuleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleBeforeDeleteHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.AfterDeleteHook, clientAllowRuleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleAfterDeleteHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.BeforeUpsertHook, clientAllowRuleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleBeforeUpsertHooks = []ClientAllowRuleHook{}

	AddClientAllowRuleHook(boil.AfterUpsertHook, clientAllowRuleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clientAllowRuleAfterUpsertHooks = []ClientAllowRuleHook{}
}

func testClientAllowRulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientAllowRulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clientAllowRuleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientAllowRuleToOneClientUsingClient(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ClientAllowRule
	var foreign Client

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clientAllowRuleDBTypes, false, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
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

	slice := ClientAllowRuleSlice{&local}
	if err = local.L.LoadClient(ctx, tx, false, (*[]*ClientAllowRule)(&slice), nil); err != nil {
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

func testClientAllowRuleToOneSetOpClientUsingClient(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClientAllowRule
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientAllowRuleDBTypes, false, strmangle.SetComplement(clientAllowRulePrimaryKeyColumns, clientAllowRuleColumnsWithoutDefault)...); err != nil {
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

		if x.R.ClientAllowRule != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClientID != x.ClientID {
			t.Error("foreign key was wrong value", a.ClientID)
		}

		if exists, err := ClientAllowRuleExists(ctx, tx, a.ClientID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testClientAllowRulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
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

func testClientAllowRulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClientAllowRuleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClientAllowRulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClientAllowRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clientAllowRuleDBTypes = map[string]string{`ClientID`: `varchar`, `UserID`: `varchar`, `EmailDomain`: `varchar`, `Created`: `datetime`}
	_                      = bytes.MinRead
)

func testClientAllowRulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clientAllowRulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clientAllowRuleAllColumns) == len(clientAllowRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClientAllowRulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clientAllowRuleAllColumns) == len(clientAllowRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClientAllowRule{}
	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clientAllowRuleDBTypes, true, clientAllowRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clientAllowRuleAllColumns, clientAllowRulePrimaryKeyColumns) {
		fields = clientAllowRuleAllColumns
	} else {
		fields = strmangle.SetComplement(
			clientAllowRuleAllColumns,
			clientAllowRulePrimaryKeyColumns,
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

	slice := ClientAllowRuleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClientAllowRulesUpsert(t *testing.T) {
	t.Parallel()

	if len(clientAllowRuleAllColumns) == len(clientAllowRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLClientAllowRuleUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClientAllowRule{}
	if err = randomize.Struct(seed, &o, clientAllowRuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientAllowRule: %s", err)
	}

	count, err := ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clientAllowRuleDBTypes, false, clientAllowRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClientAllowRule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClientAllowRule: %s", err)
	}

	count, err = ClientAllowRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
