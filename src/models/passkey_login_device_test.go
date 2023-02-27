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

func testPasskeyLoginDevices(t *testing.T) {
	t.Parallel()

	query := PasskeyLoginDevices()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPasskeyLoginDevicesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
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

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasskeyLoginDevicesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PasskeyLoginDevices().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasskeyLoginDevicesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasskeyLoginDeviceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPasskeyLoginDevicesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PasskeyLoginDeviceExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PasskeyLoginDevice exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PasskeyLoginDeviceExists to return true, but got false.")
	}
}

func testPasskeyLoginDevicesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	passkeyLoginDeviceFound, err := FindPasskeyLoginDevice(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if passkeyLoginDeviceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPasskeyLoginDevicesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PasskeyLoginDevices().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPasskeyLoginDevicesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PasskeyLoginDevices().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPasskeyLoginDevicesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	passkeyLoginDeviceOne := &PasskeyLoginDevice{}
	passkeyLoginDeviceTwo := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, passkeyLoginDeviceOne, passkeyLoginDeviceDBTypes, false, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}
	if err = randomize.Struct(seed, passkeyLoginDeviceTwo, passkeyLoginDeviceDBTypes, false, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passkeyLoginDeviceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passkeyLoginDeviceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PasskeyLoginDevices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPasskeyLoginDevicesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	passkeyLoginDeviceOne := &PasskeyLoginDevice{}
	passkeyLoginDeviceTwo := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, passkeyLoginDeviceOne, passkeyLoginDeviceDBTypes, false, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}
	if err = randomize.Struct(seed, passkeyLoginDeviceTwo, passkeyLoginDeviceDBTypes, false, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passkeyLoginDeviceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passkeyLoginDeviceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func passkeyLoginDeviceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func passkeyLoginDeviceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PasskeyLoginDevice) error {
	*o = PasskeyLoginDevice{}
	return nil
}

func testPasskeyLoginDevicesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PasskeyLoginDevice{}
	o := &PasskeyLoginDevice{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice object: %s", err)
	}

	AddPasskeyLoginDeviceHook(boil.BeforeInsertHook, passkeyLoginDeviceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceBeforeInsertHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.AfterInsertHook, passkeyLoginDeviceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceAfterInsertHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.AfterSelectHook, passkeyLoginDeviceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceAfterSelectHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.BeforeUpdateHook, passkeyLoginDeviceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceBeforeUpdateHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.AfterUpdateHook, passkeyLoginDeviceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceAfterUpdateHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.BeforeDeleteHook, passkeyLoginDeviceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceBeforeDeleteHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.AfterDeleteHook, passkeyLoginDeviceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceAfterDeleteHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.BeforeUpsertHook, passkeyLoginDeviceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceBeforeUpsertHooks = []PasskeyLoginDeviceHook{}

	AddPasskeyLoginDeviceHook(boil.AfterUpsertHook, passkeyLoginDeviceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	passkeyLoginDeviceAfterUpsertHooks = []PasskeyLoginDeviceHook{}
}

func testPasskeyLoginDevicesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasskeyLoginDevicesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(passkeyLoginDeviceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPasskeyLoginDevicesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
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

func testPasskeyLoginDevicesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PasskeyLoginDeviceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPasskeyLoginDevicesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PasskeyLoginDevices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	passkeyLoginDeviceDBTypes = map[string]string{`ID`: `int`, `UserID`: `varchar`, `Device`: `varchar`, `Os`: `varchar`, `Browser`: `varchar`, `IsRegisterDevice`: `tinyint`, `Created`: `datetime`}
	_                         = bytes.MinRead
)

func testPasskeyLoginDevicesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(passkeyLoginDevicePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(passkeyLoginDeviceAllColumns) == len(passkeyLoginDevicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDevicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPasskeyLoginDevicesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(passkeyLoginDeviceAllColumns) == len(passkeyLoginDevicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PasskeyLoginDevice{}
	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passkeyLoginDeviceDBTypes, true, passkeyLoginDevicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(passkeyLoginDeviceAllColumns, passkeyLoginDevicePrimaryKeyColumns) {
		fields = passkeyLoginDeviceAllColumns
	} else {
		fields = strmangle.SetComplement(
			passkeyLoginDeviceAllColumns,
			passkeyLoginDevicePrimaryKeyColumns,
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

	slice := PasskeyLoginDeviceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPasskeyLoginDevicesUpsert(t *testing.T) {
	t.Parallel()

	if len(passkeyLoginDeviceAllColumns) == len(passkeyLoginDevicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPasskeyLoginDeviceUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PasskeyLoginDevice{}
	if err = randomize.Struct(seed, &o, passkeyLoginDeviceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PasskeyLoginDevice: %s", err)
	}

	count, err := PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, passkeyLoginDeviceDBTypes, false, passkeyLoginDevicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PasskeyLoginDevice struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PasskeyLoginDevice: %s", err)
	}

	count, err = PasskeyLoginDevices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
