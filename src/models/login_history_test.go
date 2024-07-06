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

func testLoginHistories(t *testing.T) {
	t.Parallel()

	query := LoginHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLoginHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
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

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LoginHistories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLoginHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LoginHistoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LoginHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LoginHistoryExists to return true, but got false.")
	}
}

func testLoginHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	loginHistoryFound, err := FindLoginHistory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if loginHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLoginHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LoginHistories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLoginHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LoginHistories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLoginHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	loginHistoryOne := &LoginHistory{}
	loginHistoryTwo := &LoginHistory{}
	if err = randomize.Struct(seed, loginHistoryOne, loginHistoryDBTypes, false, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, loginHistoryTwo, loginHistoryDBTypes, false, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLoginHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	loginHistoryOne := &LoginHistory{}
	loginHistoryTwo := &LoginHistory{}
	if err = randomize.Struct(seed, loginHistoryOne, loginHistoryDBTypes, false, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, loginHistoryTwo, loginHistoryDBTypes, false, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = loginHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = loginHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func loginHistoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func loginHistoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LoginHistory) error {
	*o = LoginHistory{}
	return nil
}

func testLoginHistoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LoginHistory{}
	o := &LoginHistory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginHistory object: %s", err)
	}

	AddLoginHistoryHook(boil.BeforeInsertHook, loginHistoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	loginHistoryBeforeInsertHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.AfterInsertHook, loginHistoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	loginHistoryAfterInsertHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.AfterSelectHook, loginHistoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	loginHistoryAfterSelectHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.BeforeUpdateHook, loginHistoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	loginHistoryBeforeUpdateHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.AfterUpdateHook, loginHistoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	loginHistoryAfterUpdateHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.BeforeDeleteHook, loginHistoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	loginHistoryBeforeDeleteHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.AfterDeleteHook, loginHistoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	loginHistoryAfterDeleteHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.BeforeUpsertHook, loginHistoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	loginHistoryBeforeUpsertHooks = []LoginHistoryHook{}

	AddLoginHistoryHook(boil.AfterUpsertHook, loginHistoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	loginHistoryAfterUpsertHooks = []LoginHistoryHook{}
}

func testLoginHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(loginHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLoginHistoryToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LoginHistory
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, loginHistoryDBTypes, false, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := LoginHistorySlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*LoginHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testLoginHistoryToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LoginHistory
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, loginHistoryDBTypes, false, strmangle.SetComplement(loginHistoryPrimaryKeyColumns, loginHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LoginHistories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testLoginHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
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

func testLoginHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LoginHistorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLoginHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LoginHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	loginHistoryDBTypes = map[string]string{`ID`: `int`, `UserID`: `varchar`, `RefreshID`: `varbinary`, `Device`: `varchar`, `Os`: `varchar`, `Browser`: `varchar`, `IsMobile`: `tinyint`, `IP`: `varbinary`, `CreatedAt`: `datetime`}
	_                   = bytes.MinRead
)

func testLoginHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(loginHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(loginHistoryAllColumns) == len(loginHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLoginHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(loginHistoryAllColumns) == len(loginHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LoginHistory{}
	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, loginHistoryDBTypes, true, loginHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(loginHistoryAllColumns, loginHistoryPrimaryKeyColumns) {
		fields = loginHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			loginHistoryAllColumns,
			loginHistoryPrimaryKeyColumns,
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

	slice := LoginHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLoginHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(loginHistoryAllColumns) == len(loginHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLLoginHistoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LoginHistory{}
	if err = randomize.Struct(seed, &o, loginHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginHistory: %s", err)
	}

	count, err := LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, loginHistoryDBTypes, false, loginHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LoginHistory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LoginHistory: %s", err)
	}

	count, err = LoginHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
