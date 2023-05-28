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

func testOtpBackups(t *testing.T) {
	t.Parallel()

	query := OtpBackups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOtpBackupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
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

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOtpBackupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OtpBackups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOtpBackupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OtpBackupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOtpBackupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OtpBackupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OtpBackup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OtpBackupExists to return true, but got false.")
	}
}

func testOtpBackupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	otpBackupFound, err := FindOtpBackup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if otpBackupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOtpBackupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OtpBackups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOtpBackupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OtpBackups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOtpBackupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	otpBackupOne := &OtpBackup{}
	otpBackupTwo := &OtpBackup{}
	if err = randomize.Struct(seed, otpBackupOne, otpBackupDBTypes, false, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}
	if err = randomize.Struct(seed, otpBackupTwo, otpBackupDBTypes, false, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = otpBackupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = otpBackupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OtpBackups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOtpBackupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	otpBackupOne := &OtpBackup{}
	otpBackupTwo := &OtpBackup{}
	if err = randomize.Struct(seed, otpBackupOne, otpBackupDBTypes, false, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}
	if err = randomize.Struct(seed, otpBackupTwo, otpBackupDBTypes, false, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = otpBackupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = otpBackupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func otpBackupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func otpBackupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OtpBackup) error {
	*o = OtpBackup{}
	return nil
}

func testOtpBackupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OtpBackup{}
	o := &OtpBackup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, otpBackupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OtpBackup object: %s", err)
	}

	AddOtpBackupHook(boil.BeforeInsertHook, otpBackupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	otpBackupBeforeInsertHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.AfterInsertHook, otpBackupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	otpBackupAfterInsertHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.AfterSelectHook, otpBackupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	otpBackupAfterSelectHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.BeforeUpdateHook, otpBackupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	otpBackupBeforeUpdateHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.AfterUpdateHook, otpBackupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	otpBackupAfterUpdateHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.BeforeDeleteHook, otpBackupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	otpBackupBeforeDeleteHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.AfterDeleteHook, otpBackupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	otpBackupAfterDeleteHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.BeforeUpsertHook, otpBackupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	otpBackupBeforeUpsertHooks = []OtpBackupHook{}

	AddOtpBackupHook(boil.AfterUpsertHook, otpBackupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	otpBackupAfterUpsertHooks = []OtpBackupHook{}
}

func testOtpBackupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOtpBackupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(otpBackupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOtpBackupToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OtpBackup
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, otpBackupDBTypes, false, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
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

	slice := OtpBackupSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OtpBackup)(&slice), nil); err != nil {
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

func testOtpBackupToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OtpBackup
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, otpBackupDBTypes, false, strmangle.SetComplement(otpBackupPrimaryKeyColumns, otpBackupColumnsWithoutDefault)...); err != nil {
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

		if x.R.OtpBackups[0] != &a {
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

func testOtpBackupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
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

func testOtpBackupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OtpBackupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOtpBackupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OtpBackups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	otpBackupDBTypes = map[string]string{`ID`: `int`, `UserID`: `varchar`, `Code`: `varchar`, `CreatedAt`: `datetime`}
	_                = bytes.MinRead
)

func testOtpBackupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(otpBackupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(otpBackupAllColumns) == len(otpBackupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOtpBackupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(otpBackupAllColumns) == len(otpBackupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OtpBackup{}
	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, otpBackupDBTypes, true, otpBackupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(otpBackupAllColumns, otpBackupPrimaryKeyColumns) {
		fields = otpBackupAllColumns
	} else {
		fields = strmangle.SetComplement(
			otpBackupAllColumns,
			otpBackupPrimaryKeyColumns,
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

	slice := OtpBackupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOtpBackupsUpsert(t *testing.T) {
	t.Parallel()

	if len(otpBackupAllColumns) == len(otpBackupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOtpBackupUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OtpBackup{}
	if err = randomize.Struct(seed, &o, otpBackupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OtpBackup: %s", err)
	}

	count, err := OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, otpBackupDBTypes, false, otpBackupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OtpBackup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OtpBackup: %s", err)
	}

	count, err = OtpBackups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
