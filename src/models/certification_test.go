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

func testCertifications(t *testing.T) {
	t.Parallel()

	query := Certifications()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCertificationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
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

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCertificationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Certifications().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCertificationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CertificationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCertificationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CertificationExists(ctx, tx, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if Certification exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CertificationExists to return true, but got false.")
	}
}

func testCertificationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	certificationFound, err := FindCertification(ctx, tx, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if certificationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCertificationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Certifications().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCertificationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Certifications().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCertificationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	certificationOne := &Certification{}
	certificationTwo := &Certification{}
	if err = randomize.Struct(seed, certificationOne, certificationDBTypes, false, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}
	if err = randomize.Struct(seed, certificationTwo, certificationDBTypes, false, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = certificationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = certificationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Certifications().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCertificationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	certificationOne := &Certification{}
	certificationTwo := &Certification{}
	if err = randomize.Struct(seed, certificationOne, certificationDBTypes, false, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}
	if err = randomize.Struct(seed, certificationTwo, certificationDBTypes, false, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = certificationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = certificationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func certificationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func certificationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Certification) error {
	*o = Certification{}
	return nil
}

func testCertificationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Certification{}
	o := &Certification{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, certificationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Certification object: %s", err)
	}

	AddCertificationHook(boil.BeforeInsertHook, certificationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	certificationBeforeInsertHooks = []CertificationHook{}

	AddCertificationHook(boil.AfterInsertHook, certificationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	certificationAfterInsertHooks = []CertificationHook{}

	AddCertificationHook(boil.AfterSelectHook, certificationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	certificationAfterSelectHooks = []CertificationHook{}

	AddCertificationHook(boil.BeforeUpdateHook, certificationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	certificationBeforeUpdateHooks = []CertificationHook{}

	AddCertificationHook(boil.AfterUpdateHook, certificationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	certificationAfterUpdateHooks = []CertificationHook{}

	AddCertificationHook(boil.BeforeDeleteHook, certificationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	certificationBeforeDeleteHooks = []CertificationHook{}

	AddCertificationHook(boil.AfterDeleteHook, certificationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	certificationAfterDeleteHooks = []CertificationHook{}

	AddCertificationHook(boil.BeforeUpsertHook, certificationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	certificationBeforeUpsertHooks = []CertificationHook{}

	AddCertificationHook(boil.AfterUpsertHook, certificationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	certificationAfterUpsertHooks = []CertificationHook{}
}

func testCertificationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCertificationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(certificationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCertificationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
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

func testCertificationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CertificationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCertificationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Certifications().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	certificationDBTypes = map[string]string{`UserID`: `varbinary`, `PasskeyID`: `int`, `PasswordID`: `int`, `OtpID`: `int`, `Created`: `datetime`, `Modified`: `datetime`}
	_                    = bytes.MinRead
)

func testCertificationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(certificationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(certificationAllColumns) == len(certificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCertificationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(certificationAllColumns) == len(certificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Certification{}
	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, certificationDBTypes, true, certificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(certificationAllColumns, certificationPrimaryKeyColumns) {
		fields = certificationAllColumns
	} else {
		fields = strmangle.SetComplement(
			certificationAllColumns,
			certificationPrimaryKeyColumns,
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

	slice := CertificationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCertificationsUpsert(t *testing.T) {
	t.Parallel()

	if len(certificationAllColumns) == len(certificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCertificationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Certification{}
	if err = randomize.Struct(seed, &o, certificationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Certification: %s", err)
	}

	count, err := Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, certificationDBTypes, false, certificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Certification struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Certification: %s", err)
	}

	count, err = Certifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}