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

func testBrands(t *testing.T) {
	t.Parallel()

	query := Brands()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBrandsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
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

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBrandsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Brands().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBrandsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BrandSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBrandsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BrandExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Brand exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BrandExists to return true, but got false.")
	}
}

func testBrandsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	brandFound, err := FindBrand(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if brandFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBrandsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Brands().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBrandsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Brands().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBrandsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	brandOne := &Brand{}
	brandTwo := &Brand{}
	if err = randomize.Struct(seed, brandOne, brandDBTypes, false, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}
	if err = randomize.Struct(seed, brandTwo, brandDBTypes, false, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = brandOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = brandTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Brands().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBrandsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	brandOne := &Brand{}
	brandTwo := &Brand{}
	if err = randomize.Struct(seed, brandOne, brandDBTypes, false, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}
	if err = randomize.Struct(seed, brandTwo, brandDBTypes, false, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = brandOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = brandTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func brandBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func brandAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Brand) error {
	*o = Brand{}
	return nil
}

func testBrandsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Brand{}
	o := &Brand{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, brandDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Brand object: %s", err)
	}

	AddBrandHook(boil.BeforeInsertHook, brandBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	brandBeforeInsertHooks = []BrandHook{}

	AddBrandHook(boil.AfterInsertHook, brandAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	brandAfterInsertHooks = []BrandHook{}

	AddBrandHook(boil.AfterSelectHook, brandAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	brandAfterSelectHooks = []BrandHook{}

	AddBrandHook(boil.BeforeUpdateHook, brandBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	brandBeforeUpdateHooks = []BrandHook{}

	AddBrandHook(boil.AfterUpdateHook, brandAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	brandAfterUpdateHooks = []BrandHook{}

	AddBrandHook(boil.BeforeDeleteHook, brandBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	brandBeforeDeleteHooks = []BrandHook{}

	AddBrandHook(boil.AfterDeleteHook, brandAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	brandAfterDeleteHooks = []BrandHook{}

	AddBrandHook(boil.BeforeUpsertHook, brandBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	brandBeforeUpsertHooks = []BrandHook{}

	AddBrandHook(boil.AfterUpsertHook, brandAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	brandAfterUpsertHooks = []BrandHook{}
}

func testBrandsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBrandsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(brandColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBrandsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
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

func testBrandsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BrandSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBrandsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Brands().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	brandDBTypes = map[string]string{`ID`: `int`, `UserID`: `varchar`, `Brand`: `varchar`, `Created`: `datetime`}
	_            = bytes.MinRead
)

func testBrandsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(brandPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(brandAllColumns) == len(brandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, brandDBTypes, true, brandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBrandsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(brandAllColumns) == len(brandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Brand{}
	if err = randomize.Struct(seed, o, brandDBTypes, true, brandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, brandDBTypes, true, brandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(brandAllColumns, brandPrimaryKeyColumns) {
		fields = brandAllColumns
	} else {
		fields = strmangle.SetComplement(
			brandAllColumns,
			brandPrimaryKeyColumns,
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

	slice := BrandSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBrandsUpsert(t *testing.T) {
	t.Parallel()

	if len(brandAllColumns) == len(brandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBrandUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Brand{}
	if err = randomize.Struct(seed, &o, brandDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Brand: %s", err)
	}

	count, err := Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, brandDBTypes, false, brandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Brand struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Brand: %s", err)
	}

	count, err = Brands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
