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

func testRegisterOtpSessions(t *testing.T) {
	t.Parallel()

	query := RegisterOtpSessions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRegisterOtpSessionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
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

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterOtpSessionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RegisterOtpSessions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterOtpSessionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegisterOtpSessionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegisterOtpSessionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RegisterOtpSessionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RegisterOtpSession exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RegisterOtpSessionExists to return true, but got false.")
	}
}

func testRegisterOtpSessionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	registerOtpSessionFound, err := FindRegisterOtpSession(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if registerOtpSessionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRegisterOtpSessionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RegisterOtpSessions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRegisterOtpSessionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RegisterOtpSessions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRegisterOtpSessionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	registerOtpSessionOne := &RegisterOtpSession{}
	registerOtpSessionTwo := &RegisterOtpSession{}
	if err = randomize.Struct(seed, registerOtpSessionOne, registerOtpSessionDBTypes, false, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}
	if err = randomize.Struct(seed, registerOtpSessionTwo, registerOtpSessionDBTypes, false, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registerOtpSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registerOtpSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegisterOtpSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRegisterOtpSessionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	registerOtpSessionOne := &RegisterOtpSession{}
	registerOtpSessionTwo := &RegisterOtpSession{}
	if err = randomize.Struct(seed, registerOtpSessionOne, registerOtpSessionDBTypes, false, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}
	if err = randomize.Struct(seed, registerOtpSessionTwo, registerOtpSessionDBTypes, false, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registerOtpSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registerOtpSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func registerOtpSessionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func registerOtpSessionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegisterOtpSession) error {
	*o = RegisterOtpSession{}
	return nil
}

func testRegisterOtpSessionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RegisterOtpSession{}
	o := &RegisterOtpSession{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession object: %s", err)
	}

	AddRegisterOtpSessionHook(boil.BeforeInsertHook, registerOtpSessionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionBeforeInsertHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.AfterInsertHook, registerOtpSessionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionAfterInsertHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.AfterSelectHook, registerOtpSessionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionAfterSelectHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.BeforeUpdateHook, registerOtpSessionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionBeforeUpdateHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.AfterUpdateHook, registerOtpSessionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionAfterUpdateHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.BeforeDeleteHook, registerOtpSessionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionBeforeDeleteHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.AfterDeleteHook, registerOtpSessionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionAfterDeleteHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.BeforeUpsertHook, registerOtpSessionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionBeforeUpsertHooks = []RegisterOtpSessionHook{}

	AddRegisterOtpSessionHook(boil.AfterUpsertHook, registerOtpSessionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	registerOtpSessionAfterUpsertHooks = []RegisterOtpSessionHook{}
}

func testRegisterOtpSessionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegisterOtpSessionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(registerOtpSessionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegisterOtpSessionToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RegisterOtpSession
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, registerOtpSessionDBTypes, false, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
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

	slice := RegisterOtpSessionSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*RegisterOtpSession)(&slice), nil); err != nil {
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

func testRegisterOtpSessionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegisterOtpSession
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registerOtpSessionDBTypes, false, strmangle.SetComplement(registerOtpSessionPrimaryKeyColumns, registerOtpSessionColumnsWithoutDefault)...); err != nil {
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

		if x.R.RegisterOtpSessions[0] != &a {
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

func testRegisterOtpSessionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
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

func testRegisterOtpSessionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegisterOtpSessionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegisterOtpSessionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegisterOtpSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	registerOtpSessionDBTypes = map[string]string{`ID`: `varchar`, `UserID`: `varchar`, `PublicKey`: `text`, `Secret`: `text`, `Period`: `datetime`, `RetryCount`: `tinyint`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                         = bytes.MinRead
)

func testRegisterOtpSessionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(registerOtpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(registerOtpSessionAllColumns) == len(registerOtpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRegisterOtpSessionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(registerOtpSessionAllColumns) == len(registerOtpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegisterOtpSession{}
	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registerOtpSessionDBTypes, true, registerOtpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(registerOtpSessionAllColumns, registerOtpSessionPrimaryKeyColumns) {
		fields = registerOtpSessionAllColumns
	} else {
		fields = strmangle.SetComplement(
			registerOtpSessionAllColumns,
			registerOtpSessionPrimaryKeyColumns,
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

	slice := RegisterOtpSessionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRegisterOtpSessionsUpsert(t *testing.T) {
	t.Parallel()

	if len(registerOtpSessionAllColumns) == len(registerOtpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRegisterOtpSessionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RegisterOtpSession{}
	if err = randomize.Struct(seed, &o, registerOtpSessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegisterOtpSession: %s", err)
	}

	count, err := RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, registerOtpSessionDBTypes, false, registerOtpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegisterOtpSession struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegisterOtpSession: %s", err)
	}

	count, err = RegisterOtpSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
