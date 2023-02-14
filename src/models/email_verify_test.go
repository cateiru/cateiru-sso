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

func testEmailVerifies(t *testing.T) {
	t.Parallel()

	query := EmailVerifies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmailVerifiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
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

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailVerifiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EmailVerifies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailVerifiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailVerifySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailVerifiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmailVerifyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EmailVerify exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmailVerifyExists to return true, but got false.")
	}
}

func testEmailVerifiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	emailVerifyFound, err := FindEmailVerify(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if emailVerifyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmailVerifiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EmailVerifies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmailVerifiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EmailVerifies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmailVerifiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	emailVerifyOne := &EmailVerify{}
	emailVerifyTwo := &EmailVerify{}
	if err = randomize.Struct(seed, emailVerifyOne, emailVerifyDBTypes, false, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}
	if err = randomize.Struct(seed, emailVerifyTwo, emailVerifyDBTypes, false, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emailVerifyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailVerifyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailVerifies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmailVerifiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	emailVerifyOne := &EmailVerify{}
	emailVerifyTwo := &EmailVerify{}
	if err = randomize.Struct(seed, emailVerifyOne, emailVerifyDBTypes, false, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}
	if err = randomize.Struct(seed, emailVerifyTwo, emailVerifyDBTypes, false, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emailVerifyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailVerifyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func emailVerifyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func emailVerifyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EmailVerify) error {
	*o = EmailVerify{}
	return nil
}

func testEmailVerifiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EmailVerify{}
	o := &EmailVerify{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EmailVerify object: %s", err)
	}

	AddEmailVerifyHook(boil.BeforeInsertHook, emailVerifyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	emailVerifyBeforeInsertHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.AfterInsertHook, emailVerifyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	emailVerifyAfterInsertHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.AfterSelectHook, emailVerifyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	emailVerifyAfterSelectHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.BeforeUpdateHook, emailVerifyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	emailVerifyBeforeUpdateHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.AfterUpdateHook, emailVerifyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	emailVerifyAfterUpdateHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.BeforeDeleteHook, emailVerifyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	emailVerifyBeforeDeleteHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.AfterDeleteHook, emailVerifyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	emailVerifyAfterDeleteHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.BeforeUpsertHook, emailVerifyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	emailVerifyBeforeUpsertHooks = []EmailVerifyHook{}

	AddEmailVerifyHook(boil.AfterUpsertHook, emailVerifyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	emailVerifyAfterUpsertHooks = []EmailVerifyHook{}
}

func testEmailVerifiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailVerifiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(emailVerifyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailVerifiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
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

func testEmailVerifiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailVerifySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmailVerifiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailVerifies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	emailVerifyDBTypes = map[string]string{`ID`: `varbinary`, `UserID`: `varbinary`, `VerifyCode`: `char`, `Period`: `datetime`, `RetryCount`: `tinyint`, `Created`: `datetime`, `Modified`: `datetime`}
	_                  = bytes.MinRead
)

func testEmailVerifiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(emailVerifyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(emailVerifyAllColumns) == len(emailVerifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmailVerifiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(emailVerifyAllColumns) == len(emailVerifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailVerify{}
	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailVerifyDBTypes, true, emailVerifyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(emailVerifyAllColumns, emailVerifyPrimaryKeyColumns) {
		fields = emailVerifyAllColumns
	} else {
		fields = strmangle.SetComplement(
			emailVerifyAllColumns,
			emailVerifyPrimaryKeyColumns,
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

	slice := EmailVerifySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmailVerifiesUpsert(t *testing.T) {
	t.Parallel()

	if len(emailVerifyAllColumns) == len(emailVerifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEmailVerifyUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EmailVerify{}
	if err = randomize.Struct(seed, &o, emailVerifyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailVerify: %s", err)
	}

	count, err := EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, emailVerifyDBTypes, false, emailVerifyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailVerify struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailVerify: %s", err)
	}

	count, err = EmailVerifies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
