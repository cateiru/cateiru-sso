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

func testPasswords(t *testing.T) {
	t.Parallel()

	query := Passwords()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPasswordsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
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

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Passwords().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasswordSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasswordsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PasswordExists(ctx, tx, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if Password exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PasswordExists to return true, but got false.")
	}
}

func testPasswordsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	passwordFound, err := FindPassword(ctx, tx, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if passwordFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPasswordsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Passwords().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPasswordsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Passwords().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPasswordsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	passwordOne := &Password{}
	passwordTwo := &Password{}
	if err = randomize.Struct(seed, passwordOne, passwordDBTypes, false, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}
	if err = randomize.Struct(seed, passwordTwo, passwordDBTypes, false, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passwordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passwordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Passwords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPasswordsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	passwordOne := &Password{}
	passwordTwo := &Password{}
	if err = randomize.Struct(seed, passwordOne, passwordDBTypes, false, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}
	if err = randomize.Struct(seed, passwordTwo, passwordDBTypes, false, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passwordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passwordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func passwordBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func passwordAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Password) error {
	*o = Password{}
	return nil
}

func testPasswordsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Password{}
	o := &Password{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, passwordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Password object: %s", err)
	}

	AddPasswordHook(boil.BeforeInsertHook, passwordBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	passwordBeforeInsertHooks = []PasswordHook{}

	AddPasswordHook(boil.AfterInsertHook, passwordAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	passwordAfterInsertHooks = []PasswordHook{}

	AddPasswordHook(boil.AfterSelectHook, passwordAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	passwordAfterSelectHooks = []PasswordHook{}

	AddPasswordHook(boil.BeforeUpdateHook, passwordBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	passwordBeforeUpdateHooks = []PasswordHook{}

	AddPasswordHook(boil.AfterUpdateHook, passwordAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	passwordAfterUpdateHooks = []PasswordHook{}

	AddPasswordHook(boil.BeforeDeleteHook, passwordBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	passwordBeforeDeleteHooks = []PasswordHook{}

	AddPasswordHook(boil.AfterDeleteHook, passwordAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	passwordAfterDeleteHooks = []PasswordHook{}

	AddPasswordHook(boil.BeforeUpsertHook, passwordBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	passwordBeforeUpsertHooks = []PasswordHook{}

	AddPasswordHook(boil.AfterUpsertHook, passwordAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	passwordAfterUpsertHooks = []PasswordHook{}
}

func testPasswordsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasswordsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(passwordColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasswordToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Password
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, passwordDBTypes, false, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
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

	slice := PasswordSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Password)(&slice), nil); err != nil {
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

func testPasswordToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Password
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passwordDBTypes, false, strmangle.SetComplement(passwordPrimaryKeyColumns, passwordColumnsWithoutDefault)...); err != nil {
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

		if x.R.Password != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := PasswordExists(ctx, tx, a.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testPasswordsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
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

func testPasswordsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasswordSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPasswordsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Passwords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	passwordDBTypes = map[string]string{`UserID`: `varchar`, `Salt`: `varbinary`, `Hash`: `varbinary`, `CreatedAt`: `datetime`, `ModifiedAt`: `datetime`}
	_               = bytes.MinRead
)

func testPasswordsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(passwordPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(passwordAllColumns) == len(passwordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPasswordsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(passwordAllColumns) == len(passwordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Password{}
	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passwordDBTypes, true, passwordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(passwordAllColumns, passwordPrimaryKeyColumns) {
		fields = passwordAllColumns
	} else {
		fields = strmangle.SetComplement(
			passwordAllColumns,
			passwordPrimaryKeyColumns,
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

	slice := PasswordSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPasswordsUpsert(t *testing.T) {
	t.Parallel()

	if len(passwordAllColumns) == len(passwordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPasswordUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Password{}
	if err = randomize.Struct(seed, &o, passwordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Password: %s", err)
	}

	count, err := Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, passwordDBTypes, false, passwordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Password struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Password: %s", err)
	}

	count, err = Passwords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
