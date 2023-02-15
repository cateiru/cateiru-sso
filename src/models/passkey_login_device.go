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

// PasskeyLoginDevice is an object representing the database table.
type PasskeyLoginDevice struct {
	ID               uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	PasskeyID        uint        `boil:"passkey_id" json:"passkey_id" toml:"passkey_id" yaml:"passkey_id"`
	UserID           []byte      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Device           null.String `boil:"device" json:"device,omitempty" toml:"device" yaml:"device,omitempty"`
	Os               null.String `boil:"os" json:"os,omitempty" toml:"os" yaml:"os,omitempty"`
	Browser          null.String `boil:"browser" json:"browser,omitempty" toml:"browser" yaml:"browser,omitempty"`
	IsRegisterDevice bool        `boil:"is_register_device" json:"is_register_device" toml:"is_register_device" yaml:"is_register_device"`
	Created          time.Time   `boil:"created" json:"created" toml:"created" yaml:"created"`

	R *passkeyLoginDeviceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L passkeyLoginDeviceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PasskeyLoginDeviceColumns = struct {
	ID               string
	PasskeyID        string
	UserID           string
	Device           string
	Os               string
	Browser          string
	IsRegisterDevice string
	Created          string
}{
	ID:               "id",
	PasskeyID:        "passkey_id",
	UserID:           "user_id",
	Device:           "device",
	Os:               "os",
	Browser:          "browser",
	IsRegisterDevice: "is_register_device",
	Created:          "created",
}

var PasskeyLoginDeviceTableColumns = struct {
	ID               string
	PasskeyID        string
	UserID           string
	Device           string
	Os               string
	Browser          string
	IsRegisterDevice string
	Created          string
}{
	ID:               "passkey_login_device.id",
	PasskeyID:        "passkey_login_device.passkey_id",
	UserID:           "passkey_login_device.user_id",
	Device:           "passkey_login_device.device",
	Os:               "passkey_login_device.os",
	Browser:          "passkey_login_device.browser",
	IsRegisterDevice: "passkey_login_device.is_register_device",
	Created:          "passkey_login_device.created",
}

// Generated where

var PasskeyLoginDeviceWhere = struct {
	ID               whereHelperuint
	PasskeyID        whereHelperuint
	UserID           whereHelper__byte
	Device           whereHelpernull_String
	Os               whereHelpernull_String
	Browser          whereHelpernull_String
	IsRegisterDevice whereHelperbool
	Created          whereHelpertime_Time
}{
	ID:               whereHelperuint{field: "`passkey_login_device`.`id`"},
	PasskeyID:        whereHelperuint{field: "`passkey_login_device`.`passkey_id`"},
	UserID:           whereHelper__byte{field: "`passkey_login_device`.`user_id`"},
	Device:           whereHelpernull_String{field: "`passkey_login_device`.`device`"},
	Os:               whereHelpernull_String{field: "`passkey_login_device`.`os`"},
	Browser:          whereHelpernull_String{field: "`passkey_login_device`.`browser`"},
	IsRegisterDevice: whereHelperbool{field: "`passkey_login_device`.`is_register_device`"},
	Created:          whereHelpertime_Time{field: "`passkey_login_device`.`created`"},
}

// PasskeyLoginDeviceRels is where relationship names are stored.
var PasskeyLoginDeviceRels = struct {
}{}

// passkeyLoginDeviceR is where relationships are stored.
type passkeyLoginDeviceR struct {
}

// NewStruct creates a new relationship struct
func (*passkeyLoginDeviceR) NewStruct() *passkeyLoginDeviceR {
	return &passkeyLoginDeviceR{}
}

// passkeyLoginDeviceL is where Load methods for each relationship are stored.
type passkeyLoginDeviceL struct{}

var (
	passkeyLoginDeviceAllColumns            = []string{"id", "passkey_id", "user_id", "device", "os", "browser", "is_register_device", "created"}
	passkeyLoginDeviceColumnsWithoutDefault = []string{"passkey_id", "user_id", "device", "os", "browser"}
	passkeyLoginDeviceColumnsWithDefault    = []string{"id", "is_register_device", "created"}
	passkeyLoginDevicePrimaryKeyColumns     = []string{"id"}
	passkeyLoginDeviceGeneratedColumns      = []string{}
)

type (
	// PasskeyLoginDeviceSlice is an alias for a slice of pointers to PasskeyLoginDevice.
	// This should almost always be used instead of []PasskeyLoginDevice.
	PasskeyLoginDeviceSlice []*PasskeyLoginDevice
	// PasskeyLoginDeviceHook is the signature for custom PasskeyLoginDevice hook methods
	PasskeyLoginDeviceHook func(context.Context, boil.ContextExecutor, *PasskeyLoginDevice) error

	passkeyLoginDeviceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	passkeyLoginDeviceType                 = reflect.TypeOf(&PasskeyLoginDevice{})
	passkeyLoginDeviceMapping              = queries.MakeStructMapping(passkeyLoginDeviceType)
	passkeyLoginDevicePrimaryKeyMapping, _ = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, passkeyLoginDevicePrimaryKeyColumns)
	passkeyLoginDeviceInsertCacheMut       sync.RWMutex
	passkeyLoginDeviceInsertCache          = make(map[string]insertCache)
	passkeyLoginDeviceUpdateCacheMut       sync.RWMutex
	passkeyLoginDeviceUpdateCache          = make(map[string]updateCache)
	passkeyLoginDeviceUpsertCacheMut       sync.RWMutex
	passkeyLoginDeviceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var passkeyLoginDeviceAfterSelectHooks []PasskeyLoginDeviceHook

var passkeyLoginDeviceBeforeInsertHooks []PasskeyLoginDeviceHook
var passkeyLoginDeviceAfterInsertHooks []PasskeyLoginDeviceHook

var passkeyLoginDeviceBeforeUpdateHooks []PasskeyLoginDeviceHook
var passkeyLoginDeviceAfterUpdateHooks []PasskeyLoginDeviceHook

var passkeyLoginDeviceBeforeDeleteHooks []PasskeyLoginDeviceHook
var passkeyLoginDeviceAfterDeleteHooks []PasskeyLoginDeviceHook

var passkeyLoginDeviceBeforeUpsertHooks []PasskeyLoginDeviceHook
var passkeyLoginDeviceAfterUpsertHooks []PasskeyLoginDeviceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PasskeyLoginDevice) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PasskeyLoginDevice) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PasskeyLoginDevice) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PasskeyLoginDevice) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PasskeyLoginDevice) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PasskeyLoginDevice) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PasskeyLoginDevice) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PasskeyLoginDevice) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PasskeyLoginDevice) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passkeyLoginDeviceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPasskeyLoginDeviceHook registers your hook function for all future operations.
func AddPasskeyLoginDeviceHook(hookPoint boil.HookPoint, passkeyLoginDeviceHook PasskeyLoginDeviceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		passkeyLoginDeviceAfterSelectHooks = append(passkeyLoginDeviceAfterSelectHooks, passkeyLoginDeviceHook)
	case boil.BeforeInsertHook:
		passkeyLoginDeviceBeforeInsertHooks = append(passkeyLoginDeviceBeforeInsertHooks, passkeyLoginDeviceHook)
	case boil.AfterInsertHook:
		passkeyLoginDeviceAfterInsertHooks = append(passkeyLoginDeviceAfterInsertHooks, passkeyLoginDeviceHook)
	case boil.BeforeUpdateHook:
		passkeyLoginDeviceBeforeUpdateHooks = append(passkeyLoginDeviceBeforeUpdateHooks, passkeyLoginDeviceHook)
	case boil.AfterUpdateHook:
		passkeyLoginDeviceAfterUpdateHooks = append(passkeyLoginDeviceAfterUpdateHooks, passkeyLoginDeviceHook)
	case boil.BeforeDeleteHook:
		passkeyLoginDeviceBeforeDeleteHooks = append(passkeyLoginDeviceBeforeDeleteHooks, passkeyLoginDeviceHook)
	case boil.AfterDeleteHook:
		passkeyLoginDeviceAfterDeleteHooks = append(passkeyLoginDeviceAfterDeleteHooks, passkeyLoginDeviceHook)
	case boil.BeforeUpsertHook:
		passkeyLoginDeviceBeforeUpsertHooks = append(passkeyLoginDeviceBeforeUpsertHooks, passkeyLoginDeviceHook)
	case boil.AfterUpsertHook:
		passkeyLoginDeviceAfterUpsertHooks = append(passkeyLoginDeviceAfterUpsertHooks, passkeyLoginDeviceHook)
	}
}

// One returns a single passkeyLoginDevice record from the query.
func (q passkeyLoginDeviceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PasskeyLoginDevice, error) {
	o := &PasskeyLoginDevice{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for passkey_login_device")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PasskeyLoginDevice records from the query.
func (q passkeyLoginDeviceQuery) All(ctx context.Context, exec boil.ContextExecutor) (PasskeyLoginDeviceSlice, error) {
	var o []*PasskeyLoginDevice

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PasskeyLoginDevice slice")
	}

	if len(passkeyLoginDeviceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PasskeyLoginDevice records in the query.
func (q passkeyLoginDeviceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count passkey_login_device rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q passkeyLoginDeviceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if passkey_login_device exists")
	}

	return count > 0, nil
}

// PasskeyLoginDevices retrieves all the records using an executor.
func PasskeyLoginDevices(mods ...qm.QueryMod) passkeyLoginDeviceQuery {
	mods = append(mods, qm.From("`passkey_login_device`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`passkey_login_device`.*"})
	}

	return passkeyLoginDeviceQuery{q}
}

// FindPasskeyLoginDevice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPasskeyLoginDevice(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*PasskeyLoginDevice, error) {
	passkeyLoginDeviceObj := &PasskeyLoginDevice{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `passkey_login_device` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, passkeyLoginDeviceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from passkey_login_device")
	}

	if err = passkeyLoginDeviceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return passkeyLoginDeviceObj, err
	}

	return passkeyLoginDeviceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PasskeyLoginDevice) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no passkey_login_device provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(passkeyLoginDeviceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	passkeyLoginDeviceInsertCacheMut.RLock()
	cache, cached := passkeyLoginDeviceInsertCache[key]
	passkeyLoginDeviceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			passkeyLoginDeviceAllColumns,
			passkeyLoginDeviceColumnsWithDefault,
			passkeyLoginDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `passkey_login_device` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `passkey_login_device` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `passkey_login_device` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, passkeyLoginDevicePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into passkey_login_device")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == passkeyLoginDeviceMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for passkey_login_device")
	}

CacheNoHooks:
	if !cached {
		passkeyLoginDeviceInsertCacheMut.Lock()
		passkeyLoginDeviceInsertCache[key] = cache
		passkeyLoginDeviceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PasskeyLoginDevice.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PasskeyLoginDevice) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	passkeyLoginDeviceUpdateCacheMut.RLock()
	cache, cached := passkeyLoginDeviceUpdateCache[key]
	passkeyLoginDeviceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			passkeyLoginDeviceAllColumns,
			passkeyLoginDevicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update passkey_login_device, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `passkey_login_device` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, passkeyLoginDevicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, append(wl, passkeyLoginDevicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update passkey_login_device row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for passkey_login_device")
	}

	if !cached {
		passkeyLoginDeviceUpdateCacheMut.Lock()
		passkeyLoginDeviceUpdateCache[key] = cache
		passkeyLoginDeviceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q passkeyLoginDeviceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for passkey_login_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for passkey_login_device")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PasskeyLoginDeviceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passkeyLoginDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `passkey_login_device` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, passkeyLoginDevicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in passkeyLoginDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all passkeyLoginDevice")
	}
	return rowsAff, nil
}

var mySQLPasskeyLoginDeviceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PasskeyLoginDevice) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no passkey_login_device provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(passkeyLoginDeviceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPasskeyLoginDeviceUniqueColumns, o)

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

	passkeyLoginDeviceUpsertCacheMut.RLock()
	cache, cached := passkeyLoginDeviceUpsertCache[key]
	passkeyLoginDeviceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			passkeyLoginDeviceAllColumns,
			passkeyLoginDeviceColumnsWithDefault,
			passkeyLoginDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			passkeyLoginDeviceAllColumns,
			passkeyLoginDevicePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert passkey_login_device, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`passkey_login_device`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `passkey_login_device` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for passkey_login_device")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == passkeyLoginDeviceMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(passkeyLoginDeviceType, passkeyLoginDeviceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for passkey_login_device")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for passkey_login_device")
	}

CacheNoHooks:
	if !cached {
		passkeyLoginDeviceUpsertCacheMut.Lock()
		passkeyLoginDeviceUpsertCache[key] = cache
		passkeyLoginDeviceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PasskeyLoginDevice record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PasskeyLoginDevice) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PasskeyLoginDevice provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), passkeyLoginDevicePrimaryKeyMapping)
	sql := "DELETE FROM `passkey_login_device` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from passkey_login_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for passkey_login_device")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q passkeyLoginDeviceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no passkeyLoginDeviceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from passkey_login_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for passkey_login_device")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PasskeyLoginDeviceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(passkeyLoginDeviceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passkeyLoginDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `passkey_login_device` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, passkeyLoginDevicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from passkeyLoginDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for passkey_login_device")
	}

	if len(passkeyLoginDeviceAfterDeleteHooks) != 0 {
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
func (o *PasskeyLoginDevice) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPasskeyLoginDevice(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PasskeyLoginDeviceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PasskeyLoginDeviceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passkeyLoginDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `passkey_login_device`.* FROM `passkey_login_device` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, passkeyLoginDevicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PasskeyLoginDeviceSlice")
	}

	*o = slice

	return nil
}

// PasskeyLoginDeviceExists checks if the PasskeyLoginDevice row exists.
func PasskeyLoginDeviceExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `passkey_login_device` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if passkey_login_device exists")
	}

	return exists, nil
}

// Exists checks if the PasskeyLoginDevice row exists.
func (o *PasskeyLoginDevice) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PasskeyLoginDeviceExists(ctx, exec, o.ID)
}