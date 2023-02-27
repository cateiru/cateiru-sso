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
	"github.com/volatiletech/strmangle"
)

// LoginTryHistory is an object representing the database table.
type LoginTryHistory struct {
	ID       uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID   string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Device   null.String `boil:"device" json:"device,omitempty" toml:"device" yaml:"device,omitempty"`
	Os       null.String `boil:"os" json:"os,omitempty" toml:"os" yaml:"os,omitempty"`
	Browser  null.String `boil:"browser" json:"browser,omitempty" toml:"browser" yaml:"browser,omitempty"`
	IsMobile null.Bool   `boil:"is_mobile" json:"is_mobile,omitempty" toml:"is_mobile" yaml:"is_mobile,omitempty"`
	IP       []byte      `boil:"ip" json:"ip" toml:"ip" yaml:"ip"`
	Created  time.Time   `boil:"created" json:"created" toml:"created" yaml:"created"`

	R *loginTryHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L loginTryHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LoginTryHistoryColumns = struct {
	ID       string
	UserID   string
	Device   string
	Os       string
	Browser  string
	IsMobile string
	IP       string
	Created  string
}{
	ID:       "id",
	UserID:   "user_id",
	Device:   "device",
	Os:       "os",
	Browser:  "browser",
	IsMobile: "is_mobile",
	IP:       "ip",
	Created:  "created",
}

var LoginTryHistoryTableColumns = struct {
	ID       string
	UserID   string
	Device   string
	Os       string
	Browser  string
	IsMobile string
	IP       string
	Created  string
}{
	ID:       "login_try_history.id",
	UserID:   "login_try_history.user_id",
	Device:   "login_try_history.device",
	Os:       "login_try_history.os",
	Browser:  "login_try_history.browser",
	IsMobile: "login_try_history.is_mobile",
	IP:       "login_try_history.ip",
	Created:  "login_try_history.created",
}

// Generated where

var LoginTryHistoryWhere = struct {
	ID       whereHelperuint
	UserID   whereHelperstring
	Device   whereHelpernull_String
	Os       whereHelpernull_String
	Browser  whereHelpernull_String
	IsMobile whereHelpernull_Bool
	IP       whereHelper__byte
	Created  whereHelpertime_Time
}{
	ID:       whereHelperuint{field: "`login_try_history`.`id`"},
	UserID:   whereHelperstring{field: "`login_try_history`.`user_id`"},
	Device:   whereHelpernull_String{field: "`login_try_history`.`device`"},
	Os:       whereHelpernull_String{field: "`login_try_history`.`os`"},
	Browser:  whereHelpernull_String{field: "`login_try_history`.`browser`"},
	IsMobile: whereHelpernull_Bool{field: "`login_try_history`.`is_mobile`"},
	IP:       whereHelper__byte{field: "`login_try_history`.`ip`"},
	Created:  whereHelpertime_Time{field: "`login_try_history`.`created`"},
}

// LoginTryHistoryRels is where relationship names are stored.
var LoginTryHistoryRels = struct {
}{}

// loginTryHistoryR is where relationships are stored.
type loginTryHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*loginTryHistoryR) NewStruct() *loginTryHistoryR {
	return &loginTryHistoryR{}
}

// loginTryHistoryL is where Load methods for each relationship are stored.
type loginTryHistoryL struct{}

var (
	loginTryHistoryAllColumns            = []string{"id", "user_id", "device", "os", "browser", "is_mobile", "ip", "created"}
	loginTryHistoryColumnsWithoutDefault = []string{"user_id", "device", "os", "browser", "is_mobile", "ip"}
	loginTryHistoryColumnsWithDefault    = []string{"id", "created"}
	loginTryHistoryPrimaryKeyColumns     = []string{"id"}
	loginTryHistoryGeneratedColumns      = []string{}
)

type (
	// LoginTryHistorySlice is an alias for a slice of pointers to LoginTryHistory.
	// This should almost always be used instead of []LoginTryHistory.
	LoginTryHistorySlice []*LoginTryHistory
	// LoginTryHistoryHook is the signature for custom LoginTryHistory hook methods
	LoginTryHistoryHook func(context.Context, boil.ContextExecutor, *LoginTryHistory) error

	loginTryHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	loginTryHistoryType                 = reflect.TypeOf(&LoginTryHistory{})
	loginTryHistoryMapping              = queries.MakeStructMapping(loginTryHistoryType)
	loginTryHistoryPrimaryKeyMapping, _ = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, loginTryHistoryPrimaryKeyColumns)
	loginTryHistoryInsertCacheMut       sync.RWMutex
	loginTryHistoryInsertCache          = make(map[string]insertCache)
	loginTryHistoryUpdateCacheMut       sync.RWMutex
	loginTryHistoryUpdateCache          = make(map[string]updateCache)
	loginTryHistoryUpsertCacheMut       sync.RWMutex
	loginTryHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var loginTryHistoryAfterSelectHooks []LoginTryHistoryHook

var loginTryHistoryBeforeInsertHooks []LoginTryHistoryHook
var loginTryHistoryAfterInsertHooks []LoginTryHistoryHook

var loginTryHistoryBeforeUpdateHooks []LoginTryHistoryHook
var loginTryHistoryAfterUpdateHooks []LoginTryHistoryHook

var loginTryHistoryBeforeDeleteHooks []LoginTryHistoryHook
var loginTryHistoryAfterDeleteHooks []LoginTryHistoryHook

var loginTryHistoryBeforeUpsertHooks []LoginTryHistoryHook
var loginTryHistoryAfterUpsertHooks []LoginTryHistoryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LoginTryHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LoginTryHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LoginTryHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LoginTryHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LoginTryHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LoginTryHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LoginTryHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LoginTryHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LoginTryHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginTryHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLoginTryHistoryHook registers your hook function for all future operations.
func AddLoginTryHistoryHook(hookPoint boil.HookPoint, loginTryHistoryHook LoginTryHistoryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		loginTryHistoryAfterSelectHooks = append(loginTryHistoryAfterSelectHooks, loginTryHistoryHook)
	case boil.BeforeInsertHook:
		loginTryHistoryBeforeInsertHooks = append(loginTryHistoryBeforeInsertHooks, loginTryHistoryHook)
	case boil.AfterInsertHook:
		loginTryHistoryAfterInsertHooks = append(loginTryHistoryAfterInsertHooks, loginTryHistoryHook)
	case boil.BeforeUpdateHook:
		loginTryHistoryBeforeUpdateHooks = append(loginTryHistoryBeforeUpdateHooks, loginTryHistoryHook)
	case boil.AfterUpdateHook:
		loginTryHistoryAfterUpdateHooks = append(loginTryHistoryAfterUpdateHooks, loginTryHistoryHook)
	case boil.BeforeDeleteHook:
		loginTryHistoryBeforeDeleteHooks = append(loginTryHistoryBeforeDeleteHooks, loginTryHistoryHook)
	case boil.AfterDeleteHook:
		loginTryHistoryAfterDeleteHooks = append(loginTryHistoryAfterDeleteHooks, loginTryHistoryHook)
	case boil.BeforeUpsertHook:
		loginTryHistoryBeforeUpsertHooks = append(loginTryHistoryBeforeUpsertHooks, loginTryHistoryHook)
	case boil.AfterUpsertHook:
		loginTryHistoryAfterUpsertHooks = append(loginTryHistoryAfterUpsertHooks, loginTryHistoryHook)
	}
}

// One returns a single loginTryHistory record from the query.
func (q loginTryHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LoginTryHistory, error) {
	o := &LoginTryHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for login_try_history")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LoginTryHistory records from the query.
func (q loginTryHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (LoginTryHistorySlice, error) {
	var o []*LoginTryHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LoginTryHistory slice")
	}

	if len(loginTryHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LoginTryHistory records in the query.
func (q loginTryHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count login_try_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q loginTryHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if login_try_history exists")
	}

	return count > 0, nil
}

// LoginTryHistories retrieves all the records using an executor.
func LoginTryHistories(mods ...qm.QueryMod) loginTryHistoryQuery {
	mods = append(mods, qm.From("`login_try_history`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`login_try_history`.*"})
	}

	return loginTryHistoryQuery{q}
}

// FindLoginTryHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLoginTryHistory(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*LoginTryHistory, error) {
	loginTryHistoryObj := &LoginTryHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `login_try_history` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, loginTryHistoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from login_try_history")
	}

	if err = loginTryHistoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return loginTryHistoryObj, err
	}

	return loginTryHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LoginTryHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no login_try_history provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loginTryHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	loginTryHistoryInsertCacheMut.RLock()
	cache, cached := loginTryHistoryInsertCache[key]
	loginTryHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			loginTryHistoryAllColumns,
			loginTryHistoryColumnsWithDefault,
			loginTryHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `login_try_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `login_try_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `login_try_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, loginTryHistoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into login_try_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == loginTryHistoryMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for login_try_history")
	}

CacheNoHooks:
	if !cached {
		loginTryHistoryInsertCacheMut.Lock()
		loginTryHistoryInsertCache[key] = cache
		loginTryHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LoginTryHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LoginTryHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	loginTryHistoryUpdateCacheMut.RLock()
	cache, cached := loginTryHistoryUpdateCache[key]
	loginTryHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			loginTryHistoryAllColumns,
			loginTryHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update login_try_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `login_try_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, loginTryHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, append(wl, loginTryHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update login_try_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for login_try_history")
	}

	if !cached {
		loginTryHistoryUpdateCacheMut.Lock()
		loginTryHistoryUpdateCache[key] = cache
		loginTryHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q loginTryHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for login_try_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for login_try_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LoginTryHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginTryHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `login_try_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginTryHistoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in loginTryHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all loginTryHistory")
	}
	return rowsAff, nil
}

var mySQLLoginTryHistoryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LoginTryHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no login_try_history provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loginTryHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLoginTryHistoryUniqueColumns, o)

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

	loginTryHistoryUpsertCacheMut.RLock()
	cache, cached := loginTryHistoryUpsertCache[key]
	loginTryHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			loginTryHistoryAllColumns,
			loginTryHistoryColumnsWithDefault,
			loginTryHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			loginTryHistoryAllColumns,
			loginTryHistoryPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert login_try_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`login_try_history`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `login_try_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for login_try_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == loginTryHistoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(loginTryHistoryType, loginTryHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for login_try_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for login_try_history")
	}

CacheNoHooks:
	if !cached {
		loginTryHistoryUpsertCacheMut.Lock()
		loginTryHistoryUpsertCache[key] = cache
		loginTryHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LoginTryHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LoginTryHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LoginTryHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), loginTryHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `login_try_history` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from login_try_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for login_try_history")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q loginTryHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no loginTryHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from login_try_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for login_try_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LoginTryHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(loginTryHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginTryHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `login_try_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginTryHistoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from loginTryHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for login_try_history")
	}

	if len(loginTryHistoryAfterDeleteHooks) != 0 {
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
func (o *LoginTryHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLoginTryHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LoginTryHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LoginTryHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginTryHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `login_try_history`.* FROM `login_try_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginTryHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LoginTryHistorySlice")
	}

	*o = slice

	return nil
}

// LoginTryHistoryExists checks if the LoginTryHistory row exists.
func LoginTryHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `login_try_history` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if login_try_history exists")
	}

	return exists, nil
}

// Exists checks if the LoginTryHistory row exists.
func (o *LoginTryHistory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LoginTryHistoryExists(ctx, exec, o.ID)
}
