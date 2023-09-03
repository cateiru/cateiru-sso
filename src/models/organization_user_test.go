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

func testOrganizationUsers(t *testing.T) {
	t.Parallel()

	query := OrganizationUsers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
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

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationUsers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationUserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationUserExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OrganizationUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationUserExists to return true, but got false.")
	}
}

func testOrganizationUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationUserFound, err := FindOrganizationUser(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if organizationUserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationUsers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationUsers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationUserOne := &OrganizationUser{}
	organizationUserTwo := &OrganizationUser{}
	if err = randomize.Struct(seed, organizationUserOne, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationUserTwo, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationUserOne := &OrganizationUser{}
	organizationUserTwo := &OrganizationUser{}
	if err = randomize.Struct(seed, organizationUserOne, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationUserTwo, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationUserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func organizationUserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationUser) error {
	*o = OrganizationUser{}
	return nil
}

func testOrganizationUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationUser{}
	o := &OrganizationUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationUser object: %s", err)
	}

	AddOrganizationUserHook(boil.BeforeInsertHook, organizationUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationUserBeforeInsertHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.AfterInsertHook, organizationUserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationUserAfterInsertHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.AfterSelectHook, organizationUserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationUserAfterSelectHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.BeforeUpdateHook, organizationUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationUserBeforeUpdateHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.AfterUpdateHook, organizationUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationUserAfterUpdateHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.BeforeDeleteHook, organizationUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationUserBeforeDeleteHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.AfterDeleteHook, organizationUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationUserAfterDeleteHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.BeforeUpsertHook, organizationUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationUserBeforeUpsertHooks = []OrganizationUserHook{}

	AddOrganizationUserHook(boil.AfterUpsertHook, organizationUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationUserAfterUpsertHooks = []OrganizationUserHook{}
}

func testOrganizationUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationUserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationUserToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationUser
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrganizationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organization().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddOrganizationHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Organization) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := OrganizationUserSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*OrganizationUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organization = nil
	if err = local.L.LoadOrganization(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testOrganizationUserToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationUser
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationUserDBTypes, false, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
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

	slice := OrganizationUserSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OrganizationUser)(&slice), nil); err != nil {
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

func testOrganizationUserToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationUser
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationUserDBTypes, false, strmangle.SetComplement(organizationUserPrimaryKeyColumns, organizationUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organization{&b, &c} {
		err = a.SetOrganization(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organization != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrganizationUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganizationID))
		reflect.Indirect(reflect.ValueOf(&a.OrganizationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID, x.ID)
		}
	}
}
func testOrganizationUserToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationUser
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationUserDBTypes, false, strmangle.SetComplement(organizationUserPrimaryKeyColumns, organizationUserColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationUsers[0] != &a {
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

func testOrganizationUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
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

func testOrganizationUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationUserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationUserDBTypes = map[string]string{`ID`: `int`, `OrganizationID`: `varchar`, `UserID`: `varchar`, `Role`: `enum('owner','member','guest')`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                       = bytes.MinRead
)

func testOrganizationUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationUserAllColumns) == len(organizationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationUserAllColumns) == len(organizationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationUser{}
	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationUserDBTypes, true, organizationUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationUserAllColumns, organizationUserPrimaryKeyColumns) {
		fields = organizationUserAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationUserAllColumns,
			organizationUserPrimaryKeyColumns,
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

	slice := OrganizationUserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationUserAllColumns) == len(organizationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOrganizationUserUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationUser{}
	if err = randomize.Struct(seed, &o, organizationUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationUser: %s", err)
	}

	count, err := OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationUserDBTypes, false, organizationUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationUser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationUser: %s", err)
	}

	count, err = OrganizationUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
