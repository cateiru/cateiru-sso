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

// EmailVerifySession is an object representing the database table.
type EmailVerifySession struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	NewEmail   string    `boil:"new_email" json:"new_email" toml:"new_email" yaml:"new_email"`
	VerifyCode string    `boil:"verify_code" json:"verify_code" toml:"verify_code" yaml:"verify_code"`
	Period     time.Time `boil:"period" json:"period" toml:"period" yaml:"period"`
	RetryCount uint8     `boil:"retry_count" json:"retry_count" toml:"retry_count" yaml:"retry_count"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *emailVerifySessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L emailVerifySessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmailVerifySessionColumns = struct {
	ID         string
	UserID     string
	NewEmail   string
	VerifyCode string
	Period     string
	RetryCount string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	UserID:     "user_id",
	NewEmail:   "new_email",
	VerifyCode: "verify_code",
	Period:     "period",
	RetryCount: "retry_count",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var EmailVerifySessionTableColumns = struct {
	ID         string
	UserID     string
	NewEmail   string
	VerifyCode string
	Period     string
	RetryCount string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "email_verify_session.id",
	UserID:     "email_verify_session.user_id",
	NewEmail:   "email_verify_session.new_email",
	VerifyCode: "email_verify_session.verify_code",
	Period:     "email_verify_session.period",
	RetryCount: "email_verify_session.retry_count",
	CreatedAt:  "email_verify_session.created_at",
	UpdatedAt:  "email_verify_session.updated_at",
}

// Generated where

type whereHelperuint8 struct{ field string }

func (w whereHelperuint8) EQ(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint8) NEQ(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint8) LT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint8) LTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint8) GT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint8) GTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint8) IN(slice []uint8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint8) NIN(slice []uint8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var EmailVerifySessionWhere = struct {
	ID         whereHelperstring
	UserID     whereHelperstring
	NewEmail   whereHelperstring
	VerifyCode whereHelperstring
	Period     whereHelpertime_Time
	RetryCount whereHelperuint8
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "`email_verify_session`.`id`"},
	UserID:     whereHelperstring{field: "`email_verify_session`.`user_id`"},
	NewEmail:   whereHelperstring{field: "`email_verify_session`.`new_email`"},
	VerifyCode: whereHelperstring{field: "`email_verify_session`.`verify_code`"},
	Period:     whereHelpertime_Time{field: "`email_verify_session`.`period`"},
	RetryCount: whereHelperuint8{field: "`email_verify_session`.`retry_count`"},
	CreatedAt:  whereHelpertime_Time{field: "`email_verify_session`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`email_verify_session`.`updated_at`"},
}

// EmailVerifySessionRels is where relationship names are stored.
var EmailVerifySessionRels = struct {
	User string
}{
	User: "User",
}

// emailVerifySessionR is where relationships are stored.
type emailVerifySessionR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*emailVerifySessionR) NewStruct() *emailVerifySessionR {
	return &emailVerifySessionR{}
}

func (r *emailVerifySessionR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// emailVerifySessionL is where Load methods for each relationship are stored.
type emailVerifySessionL struct{}

var (
	emailVerifySessionAllColumns            = []string{"id", "user_id", "new_email", "verify_code", "period", "retry_count", "created_at", "updated_at"}
	emailVerifySessionColumnsWithoutDefault = []string{"id", "user_id", "new_email", "verify_code"}
	emailVerifySessionColumnsWithDefault    = []string{"period", "retry_count", "created_at", "updated_at"}
	emailVerifySessionPrimaryKeyColumns     = []string{"id"}
	emailVerifySessionGeneratedColumns      = []string{}
)

type (
	// EmailVerifySessionSlice is an alias for a slice of pointers to EmailVerifySession.
	// This should almost always be used instead of []EmailVerifySession.
	EmailVerifySessionSlice []*EmailVerifySession
	// EmailVerifySessionHook is the signature for custom EmailVerifySession hook methods
	EmailVerifySessionHook func(context.Context, boil.ContextExecutor, *EmailVerifySession) error

	emailVerifySessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	emailVerifySessionType                 = reflect.TypeOf(&EmailVerifySession{})
	emailVerifySessionMapping              = queries.MakeStructMapping(emailVerifySessionType)
	emailVerifySessionPrimaryKeyMapping, _ = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, emailVerifySessionPrimaryKeyColumns)
	emailVerifySessionInsertCacheMut       sync.RWMutex
	emailVerifySessionInsertCache          = make(map[string]insertCache)
	emailVerifySessionUpdateCacheMut       sync.RWMutex
	emailVerifySessionUpdateCache          = make(map[string]updateCache)
	emailVerifySessionUpsertCacheMut       sync.RWMutex
	emailVerifySessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var emailVerifySessionAfterSelectHooks []EmailVerifySessionHook

var emailVerifySessionBeforeInsertHooks []EmailVerifySessionHook
var emailVerifySessionAfterInsertHooks []EmailVerifySessionHook

var emailVerifySessionBeforeUpdateHooks []EmailVerifySessionHook
var emailVerifySessionAfterUpdateHooks []EmailVerifySessionHook

var emailVerifySessionBeforeDeleteHooks []EmailVerifySessionHook
var emailVerifySessionAfterDeleteHooks []EmailVerifySessionHook

var emailVerifySessionBeforeUpsertHooks []EmailVerifySessionHook
var emailVerifySessionAfterUpsertHooks []EmailVerifySessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EmailVerifySession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EmailVerifySession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EmailVerifySession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EmailVerifySession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EmailVerifySession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EmailVerifySession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EmailVerifySession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EmailVerifySession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EmailVerifySession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailVerifySessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEmailVerifySessionHook registers your hook function for all future operations.
func AddEmailVerifySessionHook(hookPoint boil.HookPoint, emailVerifySessionHook EmailVerifySessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		emailVerifySessionAfterSelectHooks = append(emailVerifySessionAfterSelectHooks, emailVerifySessionHook)
	case boil.BeforeInsertHook:
		emailVerifySessionBeforeInsertHooks = append(emailVerifySessionBeforeInsertHooks, emailVerifySessionHook)
	case boil.AfterInsertHook:
		emailVerifySessionAfterInsertHooks = append(emailVerifySessionAfterInsertHooks, emailVerifySessionHook)
	case boil.BeforeUpdateHook:
		emailVerifySessionBeforeUpdateHooks = append(emailVerifySessionBeforeUpdateHooks, emailVerifySessionHook)
	case boil.AfterUpdateHook:
		emailVerifySessionAfterUpdateHooks = append(emailVerifySessionAfterUpdateHooks, emailVerifySessionHook)
	case boil.BeforeDeleteHook:
		emailVerifySessionBeforeDeleteHooks = append(emailVerifySessionBeforeDeleteHooks, emailVerifySessionHook)
	case boil.AfterDeleteHook:
		emailVerifySessionAfterDeleteHooks = append(emailVerifySessionAfterDeleteHooks, emailVerifySessionHook)
	case boil.BeforeUpsertHook:
		emailVerifySessionBeforeUpsertHooks = append(emailVerifySessionBeforeUpsertHooks, emailVerifySessionHook)
	case boil.AfterUpsertHook:
		emailVerifySessionAfterUpsertHooks = append(emailVerifySessionAfterUpsertHooks, emailVerifySessionHook)
	}
}

// One returns a single emailVerifySession record from the query.
func (q emailVerifySessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EmailVerifySession, error) {
	o := &EmailVerifySession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for email_verify_session")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EmailVerifySession records from the query.
func (q emailVerifySessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmailVerifySessionSlice, error) {
	var o []*EmailVerifySession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EmailVerifySession slice")
	}

	if len(emailVerifySessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EmailVerifySession records in the query.
func (q emailVerifySessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count email_verify_session rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q emailVerifySessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if email_verify_session exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *EmailVerifySession) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (emailVerifySessionL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmailVerifySession interface{}, mods queries.Applicator) error {
	var slice []*EmailVerifySession
	var object *EmailVerifySession

	if singular {
		var ok bool
		object, ok = maybeEmailVerifySession.(*EmailVerifySession)
		if !ok {
			object = new(EmailVerifySession)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEmailVerifySession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEmailVerifySession))
			}
		}
	} else {
		s, ok := maybeEmailVerifySession.(*[]*EmailVerifySession)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEmailVerifySession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEmailVerifySession))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emailVerifySessionR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emailVerifySessionR{}
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
		foreign.R.EmailVerifySessions = append(foreign.R.EmailVerifySessions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.EmailVerifySessions = append(foreign.R.EmailVerifySessions, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the emailVerifySession to the related item.
// Sets o.R.User to related.
// Adds o to related.R.EmailVerifySessions.
func (o *EmailVerifySession) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `email_verify_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, emailVerifySessionPrimaryKeyColumns),
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
		o.R = &emailVerifySessionR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			EmailVerifySessions: EmailVerifySessionSlice{o},
		}
	} else {
		related.R.EmailVerifySessions = append(related.R.EmailVerifySessions, o)
	}

	return nil
}

// EmailVerifySessions retrieves all the records using an executor.
func EmailVerifySessions(mods ...qm.QueryMod) emailVerifySessionQuery {
	mods = append(mods, qm.From("`email_verify_session`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`email_verify_session`.*"})
	}

	return emailVerifySessionQuery{q}
}

// FindEmailVerifySession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmailVerifySession(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*EmailVerifySession, error) {
	emailVerifySessionObj := &EmailVerifySession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `email_verify_session` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, emailVerifySessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from email_verify_session")
	}

	if err = emailVerifySessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return emailVerifySessionObj, err
	}

	return emailVerifySessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EmailVerifySession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_verify_session provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailVerifySessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	emailVerifySessionInsertCacheMut.RLock()
	cache, cached := emailVerifySessionInsertCache[key]
	emailVerifySessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			emailVerifySessionAllColumns,
			emailVerifySessionColumnsWithDefault,
			emailVerifySessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `email_verify_session` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `email_verify_session` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `email_verify_session` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, emailVerifySessionPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into email_verify_session")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
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
		return errors.Wrap(err, "models: unable to populate default values for email_verify_session")
	}

CacheNoHooks:
	if !cached {
		emailVerifySessionInsertCacheMut.Lock()
		emailVerifySessionInsertCache[key] = cache
		emailVerifySessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EmailVerifySession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EmailVerifySession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	emailVerifySessionUpdateCacheMut.RLock()
	cache, cached := emailVerifySessionUpdateCache[key]
	emailVerifySessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			emailVerifySessionAllColumns,
			emailVerifySessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update email_verify_session, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `email_verify_session` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, emailVerifySessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, append(wl, emailVerifySessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update email_verify_session row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for email_verify_session")
	}

	if !cached {
		emailVerifySessionUpdateCacheMut.Lock()
		emailVerifySessionUpdateCache[key] = cache
		emailVerifySessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q emailVerifySessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for email_verify_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for email_verify_session")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmailVerifySessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailVerifySessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `email_verify_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailVerifySessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in emailVerifySession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all emailVerifySession")
	}
	return rowsAff, nil
}

var mySQLEmailVerifySessionUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EmailVerifySession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_verify_session provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailVerifySessionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEmailVerifySessionUniqueColumns, o)

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

	emailVerifySessionUpsertCacheMut.RLock()
	cache, cached := emailVerifySessionUpsertCache[key]
	emailVerifySessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			emailVerifySessionAllColumns,
			emailVerifySessionColumnsWithDefault,
			emailVerifySessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			emailVerifySessionAllColumns,
			emailVerifySessionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert email_verify_session, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`email_verify_session`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `email_verify_session` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for email_verify_session")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(emailVerifySessionType, emailVerifySessionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for email_verify_session")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for email_verify_session")
	}

CacheNoHooks:
	if !cached {
		emailVerifySessionUpsertCacheMut.Lock()
		emailVerifySessionUpsertCache[key] = cache
		emailVerifySessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EmailVerifySession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EmailVerifySession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmailVerifySession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), emailVerifySessionPrimaryKeyMapping)
	sql := "DELETE FROM `email_verify_session` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from email_verify_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for email_verify_session")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q emailVerifySessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no emailVerifySessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from email_verify_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_verify_session")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmailVerifySessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(emailVerifySessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailVerifySessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `email_verify_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailVerifySessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from emailVerifySession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_verify_session")
	}

	if len(emailVerifySessionAfterDeleteHooks) != 0 {
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
func (o *EmailVerifySession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmailVerifySession(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmailVerifySessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmailVerifySessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailVerifySessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `email_verify_session`.* FROM `email_verify_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailVerifySessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmailVerifySessionSlice")
	}

	*o = slice

	return nil
}

// EmailVerifySessionExists checks if the EmailVerifySession row exists.
func EmailVerifySessionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `email_verify_session` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if email_verify_session exists")
	}

	return exists, nil
}

// Exists checks if the EmailVerifySession row exists.
func (o *EmailVerifySession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EmailVerifySessionExists(ctx, exec, o.ID)
}
