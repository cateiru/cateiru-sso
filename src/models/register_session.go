// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/strmangle"
)

// RegisterSession is an object representing the database table.
type RegisterSession struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email         string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	EmailVerified bool        `boil:"email_verified" json:"email_verified" toml:"email_verified" yaml:"email_verified"`
	SendCount     uint8       `boil:"send_count" json:"send_count" toml:"send_count" yaml:"send_count"`
	VerifyCode    string      `boil:"verify_code" json:"verify_code" toml:"verify_code" yaml:"verify_code"`
	RetryCount    uint8       `boil:"retry_count" json:"retry_count" toml:"retry_count" yaml:"retry_count"`
	OrgID         null.String `boil:"org_id" json:"org_id,omitempty" toml:"org_id" yaml:"org_id,omitempty"`
	Period        time.Time   `boil:"period" json:"period" toml:"period" yaml:"period"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *registerSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L registerSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RegisterSessionColumns = struct {
	ID            string
	Email         string
	EmailVerified string
	SendCount     string
	VerifyCode    string
	RetryCount    string
	OrgID         string
	Period        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Email:         "email",
	EmailVerified: "email_verified",
	SendCount:     "send_count",
	VerifyCode:    "verify_code",
	RetryCount:    "retry_count",
	OrgID:         "org_id",
	Period:        "period",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var RegisterSessionTableColumns = struct {
	ID            string
	Email         string
	EmailVerified string
	SendCount     string
	VerifyCode    string
	RetryCount    string
	OrgID         string
	Period        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "register_session.id",
	Email:         "register_session.email",
	EmailVerified: "register_session.email_verified",
	SendCount:     "register_session.send_count",
	VerifyCode:    "register_session.verify_code",
	RetryCount:    "register_session.retry_count",
	OrgID:         "register_session.org_id",
	Period:        "register_session.period",
	CreatedAt:     "register_session.created_at",
	UpdatedAt:     "register_session.updated_at",
}

// Generated where

var RegisterSessionWhere = struct {
	ID            whereHelperstring
	Email         whereHelperstring
	EmailVerified whereHelperbool
	SendCount     whereHelperuint8
	VerifyCode    whereHelperstring
	RetryCount    whereHelperuint8
	OrgID         whereHelpernull_String
	Period        whereHelpertime_Time
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "`register_session`.`id`"},
	Email:         whereHelperstring{field: "`register_session`.`email`"},
	EmailVerified: whereHelperbool{field: "`register_session`.`email_verified`"},
	SendCount:     whereHelperuint8{field: "`register_session`.`send_count`"},
	VerifyCode:    whereHelperstring{field: "`register_session`.`verify_code`"},
	RetryCount:    whereHelperuint8{field: "`register_session`.`retry_count`"},
	OrgID:         whereHelpernull_String{field: "`register_session`.`org_id`"},
	Period:        whereHelpertime_Time{field: "`register_session`.`period`"},
	CreatedAt:     whereHelpertime_Time{field: "`register_session`.`created_at`"},
	UpdatedAt:     whereHelpertime_Time{field: "`register_session`.`updated_at`"},
}

// RegisterSessionRels is where relationship names are stored.
var RegisterSessionRels = struct {
}{}

// registerSessionR is where relationships are stored.
type registerSessionR struct {
}

// NewStruct creates a new relationship struct
func (*registerSessionR) NewStruct() *registerSessionR {
	return &registerSessionR{}
}

// registerSessionL is where Load methods for each relationship are stored.
type registerSessionL struct{}

var (
	registerSessionAllColumns            = []string{"id", "email", "email_verified", "send_count", "verify_code", "retry_count", "org_id", "period", "created_at", "updated_at"}
	registerSessionColumnsWithoutDefault = []string{"id", "email", "verify_code", "org_id"}
	registerSessionColumnsWithDefault    = []string{"email_verified", "send_count", "retry_count", "period", "created_at", "updated_at"}
	registerSessionPrimaryKeyColumns     = []string{"id"}
	registerSessionGeneratedColumns      = []string{}
)

type (
	// RegisterSessionSlice is an alias for a slice of pointers to RegisterSession.
	// This should almost always be used instead of []RegisterSession.
	RegisterSessionSlice []*RegisterSession
	// RegisterSessionHook is the signature for custom RegisterSession hook methods
	RegisterSessionHook func(context.Context, boil.ContextExecutor, *RegisterSession) error

	registerSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	registerSessionType                 = reflect.TypeOf(&RegisterSession{})
	registerSessionMapping              = queries.MakeStructMapping(registerSessionType)
	registerSessionPrimaryKeyMapping, _ = queries.BindMapping(registerSessionType, registerSessionMapping, registerSessionPrimaryKeyColumns)
	registerSessionInsertCacheMut       sync.RWMutex
	registerSessionInsertCache          = make(map[string]insertCache)
	registerSessionUpdateCacheMut       sync.RWMutex
	registerSessionUpdateCache          = make(map[string]updateCache)
	registerSessionUpsertCacheMut       sync.RWMutex
	registerSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var registerSessionAfterSelectHooks []RegisterSessionHook

var registerSessionBeforeInsertHooks []RegisterSessionHook
var registerSessionAfterInsertHooks []RegisterSessionHook

var registerSessionBeforeUpdateHooks []RegisterSessionHook
var registerSessionAfterUpdateHooks []RegisterSessionHook

var registerSessionBeforeDeleteHooks []RegisterSessionHook
var registerSessionAfterDeleteHooks []RegisterSessionHook

var registerSessionBeforeUpsertHooks []RegisterSessionHook
var registerSessionAfterUpsertHooks []RegisterSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RegisterSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RegisterSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RegisterSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RegisterSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RegisterSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RegisterSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RegisterSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RegisterSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RegisterSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRegisterSessionHook registers your hook function for all future operations.
func AddRegisterSessionHook(hookPoint boil.HookPoint, registerSessionHook RegisterSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		registerSessionAfterSelectHooks = append(registerSessionAfterSelectHooks, registerSessionHook)
	case boil.BeforeInsertHook:
		registerSessionBeforeInsertHooks = append(registerSessionBeforeInsertHooks, registerSessionHook)
	case boil.AfterInsertHook:
		registerSessionAfterInsertHooks = append(registerSessionAfterInsertHooks, registerSessionHook)
	case boil.BeforeUpdateHook:
		registerSessionBeforeUpdateHooks = append(registerSessionBeforeUpdateHooks, registerSessionHook)
	case boil.AfterUpdateHook:
		registerSessionAfterUpdateHooks = append(registerSessionAfterUpdateHooks, registerSessionHook)
	case boil.BeforeDeleteHook:
		registerSessionBeforeDeleteHooks = append(registerSessionBeforeDeleteHooks, registerSessionHook)
	case boil.AfterDeleteHook:
		registerSessionAfterDeleteHooks = append(registerSessionAfterDeleteHooks, registerSessionHook)
	case boil.BeforeUpsertHook:
		registerSessionBeforeUpsertHooks = append(registerSessionBeforeUpsertHooks, registerSessionHook)
	case boil.AfterUpsertHook:
		registerSessionAfterUpsertHooks = append(registerSessionAfterUpsertHooks, registerSessionHook)
	}
}

// One returns a single registerSession record from the query.
func (q registerSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RegisterSession, error) {
	o := &RegisterSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for register_session")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RegisterSession records from the query.
func (q registerSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (RegisterSessionSlice, error) {
	var o []*RegisterSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RegisterSession slice")
	}

	if len(registerSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RegisterSession records in the query.
func (q registerSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count register_session rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q registerSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if register_session exists")
	}

	return count > 0, nil
}

// RegisterSessions retrieves all the records using an executor.
func RegisterSessions(mods ...qm.QueryMod) registerSessionQuery {
	mods = append(mods, qm.From("`register_session`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`register_session`.*"})
	}

	return registerSessionQuery{q}
}

// FindRegisterSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRegisterSession(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*RegisterSession, error) {
	registerSessionObj := &RegisterSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `register_session` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, registerSessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from register_session")
	}

	if err = registerSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return registerSessionObj, err
	}

	return registerSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RegisterSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_session provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(registerSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	registerSessionInsertCacheMut.RLock()
	cache, cached := registerSessionInsertCache[key]
	registerSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			registerSessionAllColumns,
			registerSessionColumnsWithDefault,
			registerSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(registerSessionType, registerSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(registerSessionType, registerSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `register_session` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `register_session` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `register_session` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, registerSessionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into register_session")
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
		return errors.Wrap(err, "models: unable to populate default values for register_session")
	}

CacheNoHooks:
	if !cached {
		registerSessionInsertCacheMut.Lock()
		registerSessionInsertCache[key] = cache
		registerSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RegisterSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RegisterSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	registerSessionUpdateCacheMut.RLock()
	cache, cached := registerSessionUpdateCache[key]
	registerSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			registerSessionAllColumns,
			registerSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update register_session, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `register_session` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, registerSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(registerSessionType, registerSessionMapping, append(wl, registerSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update register_session row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for register_session")
	}

	if !cached {
		registerSessionUpdateCacheMut.Lock()
		registerSessionUpdateCache[key] = cache
		registerSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q registerSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for register_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for register_session")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RegisterSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `register_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in registerSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all registerSession")
	}
	return rowsAff, nil
}

var mySQLRegisterSessionUniqueColumns = []string{
	"id",
	"email",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RegisterSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_session provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(registerSessionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRegisterSessionUniqueColumns, o)

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

	registerSessionUpsertCacheMut.RLock()
	cache, cached := registerSessionUpsertCache[key]
	registerSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			registerSessionAllColumns,
			registerSessionColumnsWithDefault,
			registerSessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			registerSessionAllColumns,
			registerSessionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert register_session, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`register_session`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `register_session` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(registerSessionType, registerSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(registerSessionType, registerSessionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for register_session")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(registerSessionType, registerSessionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for register_session")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for register_session")
	}

CacheNoHooks:
	if !cached {
		registerSessionUpsertCacheMut.Lock()
		registerSessionUpsertCache[key] = cache
		registerSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RegisterSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RegisterSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RegisterSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), registerSessionPrimaryKeyMapping)
	sql := "DELETE FROM `register_session` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from register_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for register_session")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q registerSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no registerSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from register_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_session")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RegisterSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(registerSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `register_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from registerSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_session")
	}

	if len(registerSessionAfterDeleteHooks) != 0 {
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
func (o *RegisterSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRegisterSession(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RegisterSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RegisterSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `register_session`.* FROM `register_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RegisterSessionSlice")
	}

	*o = slice

	return nil
}

// RegisterSessionExists checks if the RegisterSession row exists.
func RegisterSessionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `register_session` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if register_session exists")
	}

	return exists, nil
}

// Exists checks if the RegisterSession row exists.
func (o *RegisterSession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RegisterSessionExists(ctx, exec, o.ID)
}
