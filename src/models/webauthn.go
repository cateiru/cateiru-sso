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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Webauthn is an object representing the database table.
type Webauthn struct {
	ID         uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Credential types.JSON  `boil:"credential" json:"credential" toml:"credential" yaml:"credential"`
	Device     null.String `boil:"device" json:"device,omitempty" toml:"device" yaml:"device,omitempty"`
	Os         null.String `boil:"os" json:"os,omitempty" toml:"os" yaml:"os,omitempty"`
	Browser    null.String `boil:"browser" json:"browser,omitempty" toml:"browser" yaml:"browser,omitempty"`
	IsMobile   null.Bool   `boil:"is_mobile" json:"is_mobile,omitempty" toml:"is_mobile" yaml:"is_mobile,omitempty"`
	IP         []byte      `boil:"ip" json:"ip" toml:"ip" yaml:"ip"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *webauthnR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L webauthnL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WebauthnColumns = struct {
	ID         string
	UserID     string
	Credential string
	Device     string
	Os         string
	Browser    string
	IsMobile   string
	IP         string
	CreatedAt  string
}{
	ID:         "id",
	UserID:     "user_id",
	Credential: "credential",
	Device:     "device",
	Os:         "os",
	Browser:    "browser",
	IsMobile:   "is_mobile",
	IP:         "ip",
	CreatedAt:  "created_at",
}

var WebauthnTableColumns = struct {
	ID         string
	UserID     string
	Credential string
	Device     string
	Os         string
	Browser    string
	IsMobile   string
	IP         string
	CreatedAt  string
}{
	ID:         "webauthn.id",
	UserID:     "webauthn.user_id",
	Credential: "webauthn.credential",
	Device:     "webauthn.device",
	Os:         "webauthn.os",
	Browser:    "webauthn.browser",
	IsMobile:   "webauthn.is_mobile",
	IP:         "webauthn.ip",
	CreatedAt:  "webauthn.created_at",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var WebauthnWhere = struct {
	ID         whereHelperuint64
	UserID     whereHelperstring
	Credential whereHelpertypes_JSON
	Device     whereHelpernull_String
	Os         whereHelpernull_String
	Browser    whereHelpernull_String
	IsMobile   whereHelpernull_Bool
	IP         whereHelper__byte
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperuint64{field: "`webauthn`.`id`"},
	UserID:     whereHelperstring{field: "`webauthn`.`user_id`"},
	Credential: whereHelpertypes_JSON{field: "`webauthn`.`credential`"},
	Device:     whereHelpernull_String{field: "`webauthn`.`device`"},
	Os:         whereHelpernull_String{field: "`webauthn`.`os`"},
	Browser:    whereHelpernull_String{field: "`webauthn`.`browser`"},
	IsMobile:   whereHelpernull_Bool{field: "`webauthn`.`is_mobile`"},
	IP:         whereHelper__byte{field: "`webauthn`.`ip`"},
	CreatedAt:  whereHelpertime_Time{field: "`webauthn`.`created_at`"},
}

// WebauthnRels is where relationship names are stored.
var WebauthnRels = struct {
	User string
}{
	User: "User",
}

// webauthnR is where relationships are stored.
type webauthnR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*webauthnR) NewStruct() *webauthnR {
	return &webauthnR{}
}

func (r *webauthnR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// webauthnL is where Load methods for each relationship are stored.
type webauthnL struct{}

var (
	webauthnAllColumns            = []string{"id", "user_id", "credential", "device", "os", "browser", "is_mobile", "ip", "created_at"}
	webauthnColumnsWithoutDefault = []string{"user_id", "credential", "device", "os", "browser", "is_mobile", "ip"}
	webauthnColumnsWithDefault    = []string{"id", "created_at"}
	webauthnPrimaryKeyColumns     = []string{"id"}
	webauthnGeneratedColumns      = []string{}
)

type (
	// WebauthnSlice is an alias for a slice of pointers to Webauthn.
	// This should almost always be used instead of []Webauthn.
	WebauthnSlice []*Webauthn
	// WebauthnHook is the signature for custom Webauthn hook methods
	WebauthnHook func(context.Context, boil.ContextExecutor, *Webauthn) error

	webauthnQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	webauthnType                 = reflect.TypeOf(&Webauthn{})
	webauthnMapping              = queries.MakeStructMapping(webauthnType)
	webauthnPrimaryKeyMapping, _ = queries.BindMapping(webauthnType, webauthnMapping, webauthnPrimaryKeyColumns)
	webauthnInsertCacheMut       sync.RWMutex
	webauthnInsertCache          = make(map[string]insertCache)
	webauthnUpdateCacheMut       sync.RWMutex
	webauthnUpdateCache          = make(map[string]updateCache)
	webauthnUpsertCacheMut       sync.RWMutex
	webauthnUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var webauthnAfterSelectHooks []WebauthnHook

var webauthnBeforeInsertHooks []WebauthnHook
var webauthnAfterInsertHooks []WebauthnHook

var webauthnBeforeUpdateHooks []WebauthnHook
var webauthnAfterUpdateHooks []WebauthnHook

var webauthnBeforeDeleteHooks []WebauthnHook
var webauthnAfterDeleteHooks []WebauthnHook

var webauthnBeforeUpsertHooks []WebauthnHook
var webauthnAfterUpsertHooks []WebauthnHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Webauthn) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Webauthn) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Webauthn) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Webauthn) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Webauthn) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Webauthn) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Webauthn) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Webauthn) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Webauthn) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range webauthnAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWebauthnHook registers your hook function for all future operations.
func AddWebauthnHook(hookPoint boil.HookPoint, webauthnHook WebauthnHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		webauthnAfterSelectHooks = append(webauthnAfterSelectHooks, webauthnHook)
	case boil.BeforeInsertHook:
		webauthnBeforeInsertHooks = append(webauthnBeforeInsertHooks, webauthnHook)
	case boil.AfterInsertHook:
		webauthnAfterInsertHooks = append(webauthnAfterInsertHooks, webauthnHook)
	case boil.BeforeUpdateHook:
		webauthnBeforeUpdateHooks = append(webauthnBeforeUpdateHooks, webauthnHook)
	case boil.AfterUpdateHook:
		webauthnAfterUpdateHooks = append(webauthnAfterUpdateHooks, webauthnHook)
	case boil.BeforeDeleteHook:
		webauthnBeforeDeleteHooks = append(webauthnBeforeDeleteHooks, webauthnHook)
	case boil.AfterDeleteHook:
		webauthnAfterDeleteHooks = append(webauthnAfterDeleteHooks, webauthnHook)
	case boil.BeforeUpsertHook:
		webauthnBeforeUpsertHooks = append(webauthnBeforeUpsertHooks, webauthnHook)
	case boil.AfterUpsertHook:
		webauthnAfterUpsertHooks = append(webauthnAfterUpsertHooks, webauthnHook)
	}
}

// One returns a single webauthn record from the query.
func (q webauthnQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Webauthn, error) {
	o := &Webauthn{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for webauthn")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Webauthn records from the query.
func (q webauthnQuery) All(ctx context.Context, exec boil.ContextExecutor) (WebauthnSlice, error) {
	var o []*Webauthn

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Webauthn slice")
	}

	if len(webauthnAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Webauthn records in the query.
func (q webauthnQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count webauthn rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q webauthnQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if webauthn exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Webauthn) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (webauthnL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWebauthn interface{}, mods queries.Applicator) error {
	var slice []*Webauthn
	var object *Webauthn

	if singular {
		var ok bool
		object, ok = maybeWebauthn.(*Webauthn)
		if !ok {
			object = new(Webauthn)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeWebauthn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeWebauthn))
			}
		}
	} else {
		s, ok := maybeWebauthn.(*[]*Webauthn)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeWebauthn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeWebauthn))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &webauthnR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &webauthnR{}
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
		foreign.R.Webauthns = append(foreign.R.Webauthns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Webauthns = append(foreign.R.Webauthns, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the webauthn to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Webauthns.
func (o *Webauthn) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `webauthn` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, webauthnPrimaryKeyColumns),
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
		o.R = &webauthnR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Webauthns: WebauthnSlice{o},
		}
	} else {
		related.R.Webauthns = append(related.R.Webauthns, o)
	}

	return nil
}

// Webauthns retrieves all the records using an executor.
func Webauthns(mods ...qm.QueryMod) webauthnQuery {
	mods = append(mods, qm.From("`webauthn`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`webauthn`.*"})
	}

	return webauthnQuery{q}
}

// FindWebauthn retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWebauthn(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*Webauthn, error) {
	webauthnObj := &Webauthn{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `webauthn` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, webauthnObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from webauthn")
	}

	if err = webauthnObj.doAfterSelectHooks(ctx, exec); err != nil {
		return webauthnObj, err
	}

	return webauthnObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Webauthn) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no webauthn provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(webauthnColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	webauthnInsertCacheMut.RLock()
	cache, cached := webauthnInsertCache[key]
	webauthnInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			webauthnAllColumns,
			webauthnColumnsWithDefault,
			webauthnColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(webauthnType, webauthnMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(webauthnType, webauthnMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `webauthn` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `webauthn` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `webauthn` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, webauthnPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into webauthn")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webauthnMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for webauthn")
	}

CacheNoHooks:
	if !cached {
		webauthnInsertCacheMut.Lock()
		webauthnInsertCache[key] = cache
		webauthnInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Webauthn.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Webauthn) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	webauthnUpdateCacheMut.RLock()
	cache, cached := webauthnUpdateCache[key]
	webauthnUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			webauthnAllColumns,
			webauthnPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update webauthn, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `webauthn` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, webauthnPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(webauthnType, webauthnMapping, append(wl, webauthnPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update webauthn row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for webauthn")
	}

	if !cached {
		webauthnUpdateCacheMut.Lock()
		webauthnUpdateCache[key] = cache
		webauthnUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q webauthnQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for webauthn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for webauthn")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WebauthnSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webauthnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `webauthn` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webauthnPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in webauthn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all webauthn")
	}
	return rowsAff, nil
}

var mySQLWebauthnUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Webauthn) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no webauthn provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(webauthnColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLWebauthnUniqueColumns, o)

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

	webauthnUpsertCacheMut.RLock()
	cache, cached := webauthnUpsertCache[key]
	webauthnUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			webauthnAllColumns,
			webauthnColumnsWithDefault,
			webauthnColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			webauthnAllColumns,
			webauthnPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert webauthn, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`webauthn`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `webauthn` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(webauthnType, webauthnMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(webauthnType, webauthnMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for webauthn")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webauthnMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(webauthnType, webauthnMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for webauthn")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for webauthn")
	}

CacheNoHooks:
	if !cached {
		webauthnUpsertCacheMut.Lock()
		webauthnUpsertCache[key] = cache
		webauthnUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Webauthn record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Webauthn) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Webauthn provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), webauthnPrimaryKeyMapping)
	sql := "DELETE FROM `webauthn` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from webauthn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for webauthn")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q webauthnQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no webauthnQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from webauthn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for webauthn")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WebauthnSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(webauthnBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webauthnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `webauthn` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webauthnPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from webauthn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for webauthn")
	}

	if len(webauthnAfterDeleteHooks) != 0 {
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
func (o *Webauthn) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWebauthn(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WebauthnSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WebauthnSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webauthnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `webauthn`.* FROM `webauthn` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webauthnPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in WebauthnSlice")
	}

	*o = slice

	return nil
}

// WebauthnExists checks if the Webauthn row exists.
func WebauthnExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `webauthn` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if webauthn exists")
	}

	return exists, nil
}

// Exists checks if the Webauthn row exists.
func (o *Webauthn) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return WebauthnExists(ctx, exec, o.ID)
}
