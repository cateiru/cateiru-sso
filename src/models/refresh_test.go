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

func testRefreshes(t *testing.T) {
	t.Parallel()

	query := Refreshes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRefreshesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
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

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRefreshesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Refreshes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRefreshesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RefreshSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRefreshesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RefreshExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Refresh exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RefreshExists to return true, but got false.")
	}
}

func testRefreshesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	refreshFound, err := FindRefresh(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if refreshFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRefreshesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Refreshes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRefreshesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Refreshes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRefreshesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	refreshOne := &Refresh{}
	refreshTwo := &Refresh{}
	if err = randomize.Struct(seed, refreshOne, refreshDBTypes, false, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}
	if err = randomize.Struct(seed, refreshTwo, refreshDBTypes, false, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = refreshOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = refreshTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Refreshes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRefreshesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	refreshOne := &Refresh{}
	refreshTwo := &Refresh{}
	if err = randomize.Struct(seed, refreshOne, refreshDBTypes, false, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}
	if err = randomize.Struct(seed, refreshTwo, refreshDBTypes, false, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = refreshOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = refreshTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func refreshBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func refreshAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Refresh) error {
	*o = Refresh{}
	return nil
}

func testRefreshesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Refresh{}
	o := &Refresh{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, refreshDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Refresh object: %s", err)
	}

	AddRefreshHook(boil.BeforeInsertHook, refreshBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	refreshBeforeInsertHooks = []RefreshHook{}

	AddRefreshHook(boil.AfterInsertHook, refreshAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	refreshAfterInsertHooks = []RefreshHook{}

	AddRefreshHook(boil.AfterSelectHook, refreshAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	refreshAfterSelectHooks = []RefreshHook{}

	AddRefreshHook(boil.BeforeUpdateHook, refreshBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	refreshBeforeUpdateHooks = []RefreshHook{}

	AddRefreshHook(boil.AfterUpdateHook, refreshAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	refreshAfterUpdateHooks = []RefreshHook{}

	AddRefreshHook(boil.BeforeDeleteHook, refreshBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	refreshBeforeDeleteHooks = []RefreshHook{}

	AddRefreshHook(boil.AfterDeleteHook, refreshAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	refreshAfterDeleteHooks = []RefreshHook{}

	AddRefreshHook(boil.BeforeUpsertHook, refreshBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	refreshBeforeUpsertHooks = []RefreshHook{}

	AddRefreshHook(boil.AfterUpsertHook, refreshAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	refreshAfterUpsertHooks = []RefreshHook{}
}

func testRefreshesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRefreshesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(refreshColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRefreshToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Refresh
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, refreshDBTypes, false, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
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

	slice := RefreshSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Refresh)(&slice), nil); err != nil {
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

func testRefreshToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Refresh
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, refreshDBTypes, false, strmangle.SetComplement(refreshPrimaryKeyColumns, refreshColumnsWithoutDefault)...); err != nil {
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

		if x.R.Refreshes[0] != &a {
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

func testRefreshesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
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

func testRefreshesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RefreshSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRefreshesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Refreshes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	refreshDBTypes = map[string]string{`ID`: `varchar`, `UserID`: `varchar`, `HistoryID`: `varbinary`, `SessionID`: `varchar`, `Period`: `datetime`, `Created`: `datetime`, `Modified`: `datetime`}
	_              = bytes.MinRead
)

func testRefreshesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(refreshPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(refreshAllColumns) == len(refreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRefreshesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(refreshAllColumns) == len(refreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Refresh{}
	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, refreshDBTypes, true, refreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(refreshAllColumns, refreshPrimaryKeyColumns) {
		fields = refreshAllColumns
	} else {
		fields = strmangle.SetComplement(
			refreshAllColumns,
			refreshPrimaryKeyColumns,
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

	slice := RefreshSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRefreshesUpsert(t *testing.T) {
	t.Parallel()

	if len(refreshAllColumns) == len(refreshPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRefreshUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Refresh{}
	if err = randomize.Struct(seed, &o, refreshDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Refresh: %s", err)
	}

	count, err := Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, refreshDBTypes, false, refreshPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Refresh struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Refresh: %s", err)
	}

	count, err = Refreshes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
