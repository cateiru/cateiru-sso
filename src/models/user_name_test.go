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

func testUserNames(t *testing.T) {
	t.Parallel()

	query := UserNames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserNamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
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

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserNamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserNames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserNamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserNameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserNamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserNameExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserName exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserNameExists to return true, but got false.")
	}
}

func testUserNamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userNameFound, err := FindUserName(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userNameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserNamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserNames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserNamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserNames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserNamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userNameOne := &UserName{}
	userNameTwo := &UserName{}
	if err = randomize.Struct(seed, userNameOne, userNameDBTypes, false, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}
	if err = randomize.Struct(seed, userNameTwo, userNameDBTypes, false, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserNamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userNameOne := &UserName{}
	userNameTwo := &UserName{}
	if err = randomize.Struct(seed, userNameOne, userNameDBTypes, false, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}
	if err = randomize.Struct(seed, userNameTwo, userNameDBTypes, false, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userNameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func userNameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserName) error {
	*o = UserName{}
	return nil
}

func testUserNamesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserName{}
	o := &UserName{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userNameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserName object: %s", err)
	}

	AddUserNameHook(boil.BeforeInsertHook, userNameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userNameBeforeInsertHooks = []UserNameHook{}

	AddUserNameHook(boil.AfterInsertHook, userNameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userNameAfterInsertHooks = []UserNameHook{}

	AddUserNameHook(boil.AfterSelectHook, userNameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userNameAfterSelectHooks = []UserNameHook{}

	AddUserNameHook(boil.BeforeUpdateHook, userNameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userNameBeforeUpdateHooks = []UserNameHook{}

	AddUserNameHook(boil.AfterUpdateHook, userNameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userNameAfterUpdateHooks = []UserNameHook{}

	AddUserNameHook(boil.BeforeDeleteHook, userNameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userNameBeforeDeleteHooks = []UserNameHook{}

	AddUserNameHook(boil.AfterDeleteHook, userNameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userNameAfterDeleteHooks = []UserNameHook{}

	AddUserNameHook(boil.BeforeUpsertHook, userNameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userNameBeforeUpsertHooks = []UserNameHook{}

	AddUserNameHook(boil.AfterUpsertHook, userNameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userNameAfterUpsertHooks = []UserNameHook{}
}

func testUserNamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserNamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userNameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserNamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
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

func testUserNamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserNameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserNamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userNameDBTypes = map[string]string{`ID`: `bigint`, `UserName`: `varchar`, `UserID`: `varchar`, `Period`: `datetime`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_               = bytes.MinRead
)

func testUserNamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userNamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userNameAllColumns) == len(userNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserNamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userNameAllColumns) == len(userNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserName{}
	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userNameDBTypes, true, userNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userNameAllColumns, userNamePrimaryKeyColumns) {
		fields = userNameAllColumns
	} else {
		fields = strmangle.SetComplement(
			userNameAllColumns,
			userNamePrimaryKeyColumns,
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

	slice := UserNameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserNamesUpsert(t *testing.T) {
	t.Parallel()

	if len(userNameAllColumns) == len(userNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLUserNameUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserName{}
	if err = randomize.Struct(seed, &o, userNameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserName: %s", err)
	}

	count, err := UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userNameDBTypes, false, userNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserName struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserName: %s", err)
	}

	count, err = UserNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
