// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserBrand is an object representing the database table.
type UserBrand struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	BrandID   string    `boil:"brand_id" json:"brand_id" toml:"brand_id" yaml:"brand_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userBrandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userBrandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserBrandColumns = struct {
	ID        string
	UserID    string
	BrandID   string
	CreatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BrandID:   "brand_id",
	CreatedAt: "created_at",
}

var UserBrandTableColumns = struct {
	ID        string
	UserID    string
	BrandID   string
	CreatedAt string
}{
	ID:        "user_brand.id",
	UserID:    "user_brand.user_id",
	BrandID:   "user_brand.brand_id",
	CreatedAt: "user_brand.created_at",
}

// Generated where

var UserBrandWhere = struct {
	ID        whereHelperuint
	UserID    whereHelperstring
	BrandID   whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperuint{field: "`user_brand`.`id`"},
	UserID:    whereHelperstring{field: "`user_brand`.`user_id`"},
	BrandID:   whereHelperstring{field: "`user_brand`.`brand_id`"},
	CreatedAt: whereHelpertime_Time{field: "`user_brand`.`created_at`"},
}

// UserBrandRels is where relationship names are stored.
var UserBrandRels = struct {
	User  string
	Brand string
}{
	User:  "User",
	Brand: "Brand",
}

// userBrandR is where relationships are stored.
type userBrandR struct {
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
	Brand *Brand `boil:"Brand" json:"Brand" toml:"Brand" yaml:"Brand"`
}

// NewStruct creates a new relationship struct
func (*userBrandR) NewStruct() *userBrandR {
	return &userBrandR{}
}

func (r *userBrandR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *userBrandR) GetBrand() *Brand {
	if r == nil {
		return nil
	}
	return r.Brand
}

// userBrandL is where Load methods for each relationship are stored.
type userBrandL struct{}

var (
	userBrandAllColumns            = []string{"id", "user_id", "brand_id", "created_at"}
	userBrandColumnsWithoutDefault = []string{"user_id", "brand_id"}
	userBrandColumnsWithDefault    = []string{"id", "created_at"}
	userBrandPrimaryKeyColumns     = []string{"id"}
	userBrandGeneratedColumns      = []string{}
)

type (
	// UserBrandSlice is an alias for a slice of pointers to UserBrand.
	// This should almost always be used instead of []UserBrand.
	UserBrandSlice []*UserBrand
	// UserBrandHook is the signature for custom UserBrand hook methods
	UserBrandHook func(context.Context, boil.ContextExecutor, *UserBrand) error

	userBrandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userBrandType                 = reflect.TypeOf(&UserBrand{})
	userBrandMapping              = queries.MakeStructMapping(userBrandType)
	userBrandPrimaryKeyMapping, _ = queries.BindMapping(userBrandType, userBrandMapping, userBrandPrimaryKeyColumns)
	userBrandInsertCacheMut       sync.RWMutex
	userBrandInsertCache          = make(map[string]insertCache)
	userBrandUpdateCacheMut       sync.RWMutex
	userBrandUpdateCache          = make(map[string]updateCache)
	userBrandUpsertCacheMut       sync.RWMutex
	userBrandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userBrandAfterSelectHooks []UserBrandHook

var userBrandBeforeInsertHooks []UserBrandHook
var userBrandAfterInsertHooks []UserBrandHook

var userBrandBeforeUpdateHooks []UserBrandHook
var userBrandAfterUpdateHooks []UserBrandHook

var userBrandBeforeDeleteHooks []UserBrandHook
var userBrandAfterDeleteHooks []UserBrandHook

var userBrandBeforeUpsertHooks []UserBrandHook
var userBrandAfterUpsertHooks []UserBrandHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserBrand) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserBrand) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserBrand) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserBrand) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserBrand) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserBrand) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserBrand) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserBrand) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserBrand) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBrandAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserBrandHook registers your hook function for all future operations.
func AddUserBrandHook(hookPoint boil.HookPoint, userBrandHook UserBrandHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userBrandAfterSelectHooks = append(userBrandAfterSelectHooks, userBrandHook)
	case boil.BeforeInsertHook:
		userBrandBeforeInsertHooks = append(userBrandBeforeInsertHooks, userBrandHook)
	case boil.AfterInsertHook:
		userBrandAfterInsertHooks = append(userBrandAfterInsertHooks, userBrandHook)
	case boil.BeforeUpdateHook:
		userBrandBeforeUpdateHooks = append(userBrandBeforeUpdateHooks, userBrandHook)
	case boil.AfterUpdateHook:
		userBrandAfterUpdateHooks = append(userBrandAfterUpdateHooks, userBrandHook)
	case boil.BeforeDeleteHook:
		userBrandBeforeDeleteHooks = append(userBrandBeforeDeleteHooks, userBrandHook)
	case boil.AfterDeleteHook:
		userBrandAfterDeleteHooks = append(userBrandAfterDeleteHooks, userBrandHook)
	case boil.BeforeUpsertHook:
		userBrandBeforeUpsertHooks = append(userBrandBeforeUpsertHooks, userBrandHook)
	case boil.AfterUpsertHook:
		userBrandAfterUpsertHooks = append(userBrandAfterUpsertHooks, userBrandHook)
	}
}

// One returns a single userBrand record from the query.
func (q userBrandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserBrand, error) {
	o := &UserBrand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_brand")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserBrand records from the query.
func (q userBrandQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserBrandSlice, error) {
	var o []*UserBrand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserBrand slice")
	}

	if len(userBrandAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserBrand records in the query.
func (q userBrandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_brand rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userBrandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_brand exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserBrand) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Brand pointed to by the foreign key.
func (o *UserBrand) Brand(mods ...qm.QueryMod) brandQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BrandID),
	}

	queryMods = append(queryMods, mods...)

	return Brands(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userBrandL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserBrand interface{}, mods queries.Applicator) error {
	var slice []*UserBrand
	var object *UserBrand

	if singular {
		var ok bool
		object, ok = maybeUserBrand.(*UserBrand)
		if !ok {
			object = new(UserBrand)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserBrand)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserBrand))
			}
		}
	} else {
		s, ok := maybeUserBrand.(*[]*UserBrand)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserBrand)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserBrand))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userBrandR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userBrandR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserBrands = append(foreign.R.UserBrands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserBrands = append(foreign.R.UserBrands, local)
				break
			}
		}
	}

	return nil
}

// LoadBrand allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userBrandL) LoadBrand(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserBrand interface{}, mods queries.Applicator) error {
	var slice []*UserBrand
	var object *UserBrand

	if singular {
		var ok bool
		object, ok = maybeUserBrand.(*UserBrand)
		if !ok {
			object = new(UserBrand)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserBrand)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserBrand))
			}
		}
	} else {
		s, ok := maybeUserBrand.(*[]*UserBrand)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserBrand)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserBrand))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userBrandR{}
		}
		args = append(args, object.BrandID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userBrandR{}
			}

			for _, a := range args {
				if a == obj.BrandID {
					continue Outer
				}
			}

			args = append(args, obj.BrandID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`brand`),
		qm.WhereIn(`brand.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Brand")
	}

	var resultSlice []*Brand
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Brand")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for brand")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for brand")
	}

	if len(brandAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Brand = foreign
		if foreign.R == nil {
			foreign.R = &brandR{}
		}
		foreign.R.UserBrands = append(foreign.R.UserBrands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BrandID == foreign.ID {
				local.R.Brand = foreign
				if foreign.R == nil {
					foreign.R = &brandR{}
				}
				foreign.R.UserBrands = append(foreign.R.UserBrands, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userBrand to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserBrands.
func (o *UserBrand) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_brand` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, userBrandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userBrandR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserBrands: UserBrandSlice{o},
		}
	} else {
		related.R.UserBrands = append(related.R.UserBrands, o)
	}

	return nil
}

// SetBrand of the userBrand to the related item.
// Sets o.R.Brand to related.
// Adds o to related.R.UserBrands.
func (o *UserBrand) SetBrand(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Brand) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_brand` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"brand_id"}),
		strmangle.WhereClause("`", "`", 0, userBrandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BrandID = related.ID
	if o.R == nil {
		o.R = &userBrandR{
			Brand: related,
		}
	} else {
		o.R.Brand = related
	}

	if related.R == nil {
		related.R = &brandR{
			UserBrands: UserBrandSlice{o},
		}
	} else {
		related.R.UserBrands = append(related.R.UserBrands, o)
	}

	return nil
}

// UserBrands retrieves all the records using an executor.
func UserBrands(mods ...qm.QueryMod) userBrandQuery {
	mods = append(mods, qm.From("`user_brand`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_brand`.*"})
	}

	return userBrandQuery{q}
}

// FindUserBrand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserBrand(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*UserBrand, error) {
	userBrandObj := &UserBrand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_brand` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userBrandObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_brand")
	}

	if err = userBrandObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userBrandObj, err
	}

	return userBrandObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserBrand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_brand provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBrandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userBrandInsertCacheMut.RLock()
	cache, cached := userBrandInsertCache[key]
	userBrandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userBrandAllColumns,
			userBrandColumnsWithDefault,
			userBrandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userBrandType, userBrandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userBrandType, userBrandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_brand` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_brand` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_brand` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userBrandPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_brand")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userBrandMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_brand")
	}

CacheNoHooks:
	if !cached {
		userBrandInsertCacheMut.Lock()
		userBrandInsertCache[key] = cache
		userBrandInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserBrand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserBrand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userBrandUpdateCacheMut.RLock()
	cache, cached := userBrandUpdateCache[key]
	userBrandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userBrandAllColumns,
			userBrandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_brand, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_brand` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userBrandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userBrandType, userBrandMapping, append(wl, userBrandPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update user_brand row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_brand")
	}

	if !cached {
		userBrandUpdateCacheMut.Lock()
		userBrandUpdateCache[key] = cache
		userBrandUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userBrandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_brand")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_brand")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserBrandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBrandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_brand` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userBrandPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userBrand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userBrand")
	}
	return rowsAff, nil
}

var mySQLUserBrandUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserBrand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_brand provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBrandColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserBrandUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userBrandUpsertCacheMut.RLock()
	cache, cached := userBrandUpsertCache[key]
	userBrandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userBrandAllColumns,
			userBrandColumnsWithDefault,
			userBrandColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userBrandAllColumns,
			userBrandPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert user_brand, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_brand`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_brand` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userBrandType, userBrandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userBrandType, userBrandMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for user_brand")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userBrandMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userBrandType, userBrandMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for user_brand")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_brand")
	}

CacheNoHooks:
	if !cached {
		userBrandUpsertCacheMut.Lock()
		userBrandUpsertCache[key] = cache
		userBrandUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserBrand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserBrand) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserBrand provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userBrandPrimaryKeyMapping)
	sql := "DELETE FROM `user_brand` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_brand")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_brand")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userBrandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userBrandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_brand")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_brand")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserBrandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userBrandBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBrandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_brand` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userBrandPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userBrand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_brand")
	}

	if len(userBrandAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserBrand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserBrand(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserBrandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserBrandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBrandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_brand`.* FROM `user_brand` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userBrandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserBrandSlice")
	}

	*o = slice

	return nil
}

// UserBrandExists checks if the UserBrand row exists.
func UserBrandExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_brand` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_brand exists")
	}

	return exists, nil
}

// Exists checks if the UserBrand row exists.
func (o *UserBrand) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserBrandExists(ctx, exec, o.ID)
}
