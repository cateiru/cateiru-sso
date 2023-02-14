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

// RegisterOtpSession is an object representing the database table.
type RegisterOtpSession struct {
	ID         []byte    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     []byte    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	PublicKey  string    `boil:"public_key" json:"public_key" toml:"public_key" yaml:"public_key"`
	Secret     string    `boil:"secret" json:"secret" toml:"secret" yaml:"secret"`
	Period     time.Time `boil:"period" json:"period" toml:"period" yaml:"period"`
	RetryCount uint8     `boil:"retry_count" json:"retry_count" toml:"retry_count" yaml:"retry_count"`
	Created    time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`
	Modified   time.Time `boil:"modified" json:"modified" toml:"modified" yaml:"modified"`

	R *registerOtpSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L registerOtpSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RegisterOtpSessionColumns = struct {
	ID         string
	UserID     string
	PublicKey  string
	Secret     string
	Period     string
	RetryCount string
	Created    string
	Modified   string
}{
	ID:         "id",
	UserID:     "user_id",
	PublicKey:  "public_key",
	Secret:     "secret",
	Period:     "period",
	RetryCount: "retry_count",
	Created:    "created",
	Modified:   "modified",
}

var RegisterOtpSessionTableColumns = struct {
	ID         string
	UserID     string
	PublicKey  string
	Secret     string
	Period     string
	RetryCount string
	Created    string
	Modified   string
}{
	ID:         "register_otp_session.id",
	UserID:     "register_otp_session.user_id",
	PublicKey:  "register_otp_session.public_key",
	Secret:     "register_otp_session.secret",
	Period:     "register_otp_session.period",
	RetryCount: "register_otp_session.retry_count",
	Created:    "register_otp_session.created",
	Modified:   "register_otp_session.modified",
}

// Generated where

var RegisterOtpSessionWhere = struct {
	ID         whereHelper__byte
	UserID     whereHelper__byte
	PublicKey  whereHelperstring
	Secret     whereHelperstring
	Period     whereHelpertime_Time
	RetryCount whereHelperuint8
	Created    whereHelpertime_Time
	Modified   whereHelpertime_Time
}{
	ID:         whereHelper__byte{field: "`register_otp_session`.`id`"},
	UserID:     whereHelper__byte{field: "`register_otp_session`.`user_id`"},
	PublicKey:  whereHelperstring{field: "`register_otp_session`.`public_key`"},
	Secret:     whereHelperstring{field: "`register_otp_session`.`secret`"},
	Period:     whereHelpertime_Time{field: "`register_otp_session`.`period`"},
	RetryCount: whereHelperuint8{field: "`register_otp_session`.`retry_count`"},
	Created:    whereHelpertime_Time{field: "`register_otp_session`.`created`"},
	Modified:   whereHelpertime_Time{field: "`register_otp_session`.`modified`"},
}

// RegisterOtpSessionRels is where relationship names are stored.
var RegisterOtpSessionRels = struct {
}{}

// registerOtpSessionR is where relationships are stored.
type registerOtpSessionR struct {
}

// NewStruct creates a new relationship struct
func (*registerOtpSessionR) NewStruct() *registerOtpSessionR {
	return &registerOtpSessionR{}
}

// registerOtpSessionL is where Load methods for each relationship are stored.
type registerOtpSessionL struct{}

var (
	registerOtpSessionAllColumns            = []string{"id", "user_id", "public_key", "secret", "period", "retry_count", "created", "modified"}
	registerOtpSessionColumnsWithoutDefault = []string{"id", "user_id", "public_key", "secret"}
	registerOtpSessionColumnsWithDefault    = []string{"period", "retry_count", "created", "modified"}
	registerOtpSessionPrimaryKeyColumns     = []string{"id"}
	registerOtpSessionGeneratedColumns      = []string{}
)

type (
	// RegisterOtpSessionSlice is an alias for a slice of pointers to RegisterOtpSession.
	// This should almost always be used instead of []RegisterOtpSession.
	RegisterOtpSessionSlice []*RegisterOtpSession
	// RegisterOtpSessionHook is the signature for custom RegisterOtpSession hook methods
	RegisterOtpSessionHook func(context.Context, boil.ContextExecutor, *RegisterOtpSession) error

	registerOtpSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	registerOtpSessionType                 = reflect.TypeOf(&RegisterOtpSession{})
	registerOtpSessionMapping              = queries.MakeStructMapping(registerOtpSessionType)
	registerOtpSessionPrimaryKeyMapping, _ = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, registerOtpSessionPrimaryKeyColumns)
	registerOtpSessionInsertCacheMut       sync.RWMutex
	registerOtpSessionInsertCache          = make(map[string]insertCache)
	registerOtpSessionUpdateCacheMut       sync.RWMutex
	registerOtpSessionUpdateCache          = make(map[string]updateCache)
	registerOtpSessionUpsertCacheMut       sync.RWMutex
	registerOtpSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var registerOtpSessionAfterSelectHooks []RegisterOtpSessionHook

var registerOtpSessionBeforeInsertHooks []RegisterOtpSessionHook
var registerOtpSessionAfterInsertHooks []RegisterOtpSessionHook

var registerOtpSessionBeforeUpdateHooks []RegisterOtpSessionHook
var registerOtpSessionAfterUpdateHooks []RegisterOtpSessionHook

var registerOtpSessionBeforeDeleteHooks []RegisterOtpSessionHook
var registerOtpSessionAfterDeleteHooks []RegisterOtpSessionHook

var registerOtpSessionBeforeUpsertHooks []RegisterOtpSessionHook
var registerOtpSessionAfterUpsertHooks []RegisterOtpSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RegisterOtpSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RegisterOtpSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RegisterOtpSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RegisterOtpSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RegisterOtpSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RegisterOtpSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RegisterOtpSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RegisterOtpSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RegisterOtpSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerOtpSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRegisterOtpSessionHook registers your hook function for all future operations.
func AddRegisterOtpSessionHook(hookPoint boil.HookPoint, registerOtpSessionHook RegisterOtpSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		registerOtpSessionAfterSelectHooks = append(registerOtpSessionAfterSelectHooks, registerOtpSessionHook)
	case boil.BeforeInsertHook:
		registerOtpSessionBeforeInsertHooks = append(registerOtpSessionBeforeInsertHooks, registerOtpSessionHook)
	case boil.AfterInsertHook:
		registerOtpSessionAfterInsertHooks = append(registerOtpSessionAfterInsertHooks, registerOtpSessionHook)
	case boil.BeforeUpdateHook:
		registerOtpSessionBeforeUpdateHooks = append(registerOtpSessionBeforeUpdateHooks, registerOtpSessionHook)
	case boil.AfterUpdateHook:
		registerOtpSessionAfterUpdateHooks = append(registerOtpSessionAfterUpdateHooks, registerOtpSessionHook)
	case boil.BeforeDeleteHook:
		registerOtpSessionBeforeDeleteHooks = append(registerOtpSessionBeforeDeleteHooks, registerOtpSessionHook)
	case boil.AfterDeleteHook:
		registerOtpSessionAfterDeleteHooks = append(registerOtpSessionAfterDeleteHooks, registerOtpSessionHook)
	case boil.BeforeUpsertHook:
		registerOtpSessionBeforeUpsertHooks = append(registerOtpSessionBeforeUpsertHooks, registerOtpSessionHook)
	case boil.AfterUpsertHook:
		registerOtpSessionAfterUpsertHooks = append(registerOtpSessionAfterUpsertHooks, registerOtpSessionHook)
	}
}

// One returns a single registerOtpSession record from the query.
func (q registerOtpSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RegisterOtpSession, error) {
	o := &RegisterOtpSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for register_otp_session")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RegisterOtpSession records from the query.
func (q registerOtpSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (RegisterOtpSessionSlice, error) {
	var o []*RegisterOtpSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RegisterOtpSession slice")
	}

	if len(registerOtpSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RegisterOtpSession records in the query.
func (q registerOtpSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count register_otp_session rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q registerOtpSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if register_otp_session exists")
	}

	return count > 0, nil
}

// RegisterOtpSessions retrieves all the records using an executor.
func RegisterOtpSessions(mods ...qm.QueryMod) registerOtpSessionQuery {
	mods = append(mods, qm.From("`register_otp_session`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`register_otp_session`.*"})
	}

	return registerOtpSessionQuery{q}
}

// FindRegisterOtpSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRegisterOtpSession(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*RegisterOtpSession, error) {
	registerOtpSessionObj := &RegisterOtpSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `register_otp_session` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, registerOtpSessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from register_otp_session")
	}

	if err = registerOtpSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return registerOtpSessionObj, err
	}

	return registerOtpSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RegisterOtpSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_otp_session provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerOtpSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	registerOtpSessionInsertCacheMut.RLock()
	cache, cached := registerOtpSessionInsertCache[key]
	registerOtpSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			registerOtpSessionAllColumns,
			registerOtpSessionColumnsWithDefault,
			registerOtpSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `register_otp_session` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `register_otp_session` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `register_otp_session` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, registerOtpSessionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into register_otp_session")
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
		return errors.Wrap(err, "models: unable to populate default values for register_otp_session")
	}

CacheNoHooks:
	if !cached {
		registerOtpSessionInsertCacheMut.Lock()
		registerOtpSessionInsertCache[key] = cache
		registerOtpSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RegisterOtpSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RegisterOtpSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	registerOtpSessionUpdateCacheMut.RLock()
	cache, cached := registerOtpSessionUpdateCache[key]
	registerOtpSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			registerOtpSessionAllColumns,
			registerOtpSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update register_otp_session, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `register_otp_session` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, registerOtpSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, append(wl, registerOtpSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update register_otp_session row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for register_otp_session")
	}

	if !cached {
		registerOtpSessionUpdateCacheMut.Lock()
		registerOtpSessionUpdateCache[key] = cache
		registerOtpSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q registerOtpSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for register_otp_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for register_otp_session")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RegisterOtpSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerOtpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `register_otp_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerOtpSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in registerOtpSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all registerOtpSession")
	}
	return rowsAff, nil
}

var mySQLRegisterOtpSessionUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RegisterOtpSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_otp_session provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerOtpSessionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRegisterOtpSessionUniqueColumns, o)

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

	registerOtpSessionUpsertCacheMut.RLock()
	cache, cached := registerOtpSessionUpsertCache[key]
	registerOtpSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			registerOtpSessionAllColumns,
			registerOtpSessionColumnsWithDefault,
			registerOtpSessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			registerOtpSessionAllColumns,
			registerOtpSessionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert register_otp_session, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`register_otp_session`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `register_otp_session` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for register_otp_session")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(registerOtpSessionType, registerOtpSessionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for register_otp_session")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for register_otp_session")
	}

CacheNoHooks:
	if !cached {
		registerOtpSessionUpsertCacheMut.Lock()
		registerOtpSessionUpsertCache[key] = cache
		registerOtpSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RegisterOtpSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RegisterOtpSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RegisterOtpSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), registerOtpSessionPrimaryKeyMapping)
	sql := "DELETE FROM `register_otp_session` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from register_otp_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for register_otp_session")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q registerOtpSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no registerOtpSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from register_otp_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_otp_session")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RegisterOtpSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(registerOtpSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerOtpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `register_otp_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerOtpSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from registerOtpSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_otp_session")
	}

	if len(registerOtpSessionAfterDeleteHooks) != 0 {
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
func (o *RegisterOtpSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRegisterOtpSession(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RegisterOtpSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RegisterOtpSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerOtpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `register_otp_session`.* FROM `register_otp_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, registerOtpSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RegisterOtpSessionSlice")
	}

	*o = slice

	return nil
}

// RegisterOtpSessionExists checks if the RegisterOtpSession row exists.
func RegisterOtpSessionExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `register_otp_session` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if register_otp_session exists")
	}

	return exists, nil
}

// Exists checks if the RegisterOtpSession row exists.
func (o *RegisterOtpSession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RegisterOtpSessionExists(ctx, exec, o.ID)
}
