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

// ClientRedirect is an object representing the database table.
type ClientRedirect struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClientID  string    `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	URL       string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *clientRedirectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientRedirectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientRedirectColumns = struct {
	ID        string
	ClientID  string
	URL       string
	CreatedAt string
}{
	ID:        "id",
	ClientID:  "client_id",
	URL:       "url",
	CreatedAt: "created_at",
}

var ClientRedirectTableColumns = struct {
	ID        string
	ClientID  string
	URL       string
	CreatedAt string
}{
	ID:        "client_redirect.id",
	ClientID:  "client_redirect.client_id",
	URL:       "client_redirect.url",
	CreatedAt: "client_redirect.created_at",
}

// Generated where

var ClientRedirectWhere = struct {
	ID        whereHelperuint
	ClientID  whereHelperstring
	URL       whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperuint{field: "`client_redirect`.`id`"},
	ClientID:  whereHelperstring{field: "`client_redirect`.`client_id`"},
	URL:       whereHelperstring{field: "`client_redirect`.`url`"},
	CreatedAt: whereHelpertime_Time{field: "`client_redirect`.`created_at`"},
}

// ClientRedirectRels is where relationship names are stored.
var ClientRedirectRels = struct {
	Client string
}{
	Client: "Client",
}

// clientRedirectR is where relationships are stored.
type clientRedirectR struct {
	Client *Client `boil:"Client" json:"Client" toml:"Client" yaml:"Client"`
}

// NewStruct creates a new relationship struct
func (*clientRedirectR) NewStruct() *clientRedirectR {
	return &clientRedirectR{}
}

func (r *clientRedirectR) GetClient() *Client {
	if r == nil {
		return nil
	}
	return r.Client
}

// clientRedirectL is where Load methods for each relationship are stored.
type clientRedirectL struct{}

var (
	clientRedirectAllColumns            = []string{"id", "client_id", "url", "created_at"}
	clientRedirectColumnsWithoutDefault = []string{"client_id", "url"}
	clientRedirectColumnsWithDefault    = []string{"id", "created_at"}
	clientRedirectPrimaryKeyColumns     = []string{"id"}
	clientRedirectGeneratedColumns      = []string{}
)

type (
	// ClientRedirectSlice is an alias for a slice of pointers to ClientRedirect.
	// This should almost always be used instead of []ClientRedirect.
	ClientRedirectSlice []*ClientRedirect
	// ClientRedirectHook is the signature for custom ClientRedirect hook methods
	ClientRedirectHook func(context.Context, boil.ContextExecutor, *ClientRedirect) error

	clientRedirectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientRedirectType                 = reflect.TypeOf(&ClientRedirect{})
	clientRedirectMapping              = queries.MakeStructMapping(clientRedirectType)
	clientRedirectPrimaryKeyMapping, _ = queries.BindMapping(clientRedirectType, clientRedirectMapping, clientRedirectPrimaryKeyColumns)
	clientRedirectInsertCacheMut       sync.RWMutex
	clientRedirectInsertCache          = make(map[string]insertCache)
	clientRedirectUpdateCacheMut       sync.RWMutex
	clientRedirectUpdateCache          = make(map[string]updateCache)
	clientRedirectUpsertCacheMut       sync.RWMutex
	clientRedirectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientRedirectAfterSelectHooks []ClientRedirectHook

var clientRedirectBeforeInsertHooks []ClientRedirectHook
var clientRedirectAfterInsertHooks []ClientRedirectHook

var clientRedirectBeforeUpdateHooks []ClientRedirectHook
var clientRedirectAfterUpdateHooks []ClientRedirectHook

var clientRedirectBeforeDeleteHooks []ClientRedirectHook
var clientRedirectAfterDeleteHooks []ClientRedirectHook

var clientRedirectBeforeUpsertHooks []ClientRedirectHook
var clientRedirectAfterUpsertHooks []ClientRedirectHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClientRedirect) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClientRedirect) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClientRedirect) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClientRedirect) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClientRedirect) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClientRedirect) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClientRedirect) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClientRedirect) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClientRedirect) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientRedirectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientRedirectHook registers your hook function for all future operations.
func AddClientRedirectHook(hookPoint boil.HookPoint, clientRedirectHook ClientRedirectHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientRedirectAfterSelectHooks = append(clientRedirectAfterSelectHooks, clientRedirectHook)
	case boil.BeforeInsertHook:
		clientRedirectBeforeInsertHooks = append(clientRedirectBeforeInsertHooks, clientRedirectHook)
	case boil.AfterInsertHook:
		clientRedirectAfterInsertHooks = append(clientRedirectAfterInsertHooks, clientRedirectHook)
	case boil.BeforeUpdateHook:
		clientRedirectBeforeUpdateHooks = append(clientRedirectBeforeUpdateHooks, clientRedirectHook)
	case boil.AfterUpdateHook:
		clientRedirectAfterUpdateHooks = append(clientRedirectAfterUpdateHooks, clientRedirectHook)
	case boil.BeforeDeleteHook:
		clientRedirectBeforeDeleteHooks = append(clientRedirectBeforeDeleteHooks, clientRedirectHook)
	case boil.AfterDeleteHook:
		clientRedirectAfterDeleteHooks = append(clientRedirectAfterDeleteHooks, clientRedirectHook)
	case boil.BeforeUpsertHook:
		clientRedirectBeforeUpsertHooks = append(clientRedirectBeforeUpsertHooks, clientRedirectHook)
	case boil.AfterUpsertHook:
		clientRedirectAfterUpsertHooks = append(clientRedirectAfterUpsertHooks, clientRedirectHook)
	}
}

// One returns a single clientRedirect record from the query.
func (q clientRedirectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClientRedirect, error) {
	o := &ClientRedirect{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for client_redirect")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClientRedirect records from the query.
func (q clientRedirectQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientRedirectSlice, error) {
	var o []*ClientRedirect

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClientRedirect slice")
	}

	if len(clientRedirectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClientRedirect records in the query.
func (q clientRedirectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count client_redirect rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clientRedirectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if client_redirect exists")
	}

	return count > 0, nil
}

// Client pointed to by the foreign key.
func (o *ClientRedirect) Client(mods ...qm.QueryMod) clientQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`client_id` = ?", o.ClientID),
	}

	queryMods = append(queryMods, mods...)

	return Clients(queryMods...)
}

// LoadClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (clientRedirectL) LoadClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClientRedirect interface{}, mods queries.Applicator) error {
	var slice []*ClientRedirect
	var object *ClientRedirect

	if singular {
		var ok bool
		object, ok = maybeClientRedirect.(*ClientRedirect)
		if !ok {
			object = new(ClientRedirect)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeClientRedirect)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeClientRedirect))
			}
		}
	} else {
		s, ok := maybeClientRedirect.(*[]*ClientRedirect)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeClientRedirect)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeClientRedirect))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clientRedirectR{}
		}
		args = append(args, object.ClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clientRedirectR{}
			}

			for _, a := range args {
				if a == obj.ClientID {
					continue Outer
				}
			}

			args = append(args, obj.ClientID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`client`),
		qm.WhereIn(`client.client_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Client")
	}

	var resultSlice []*Client
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Client")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for client")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for client")
	}

	if len(clientAfterSelectHooks) != 0 {
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
		object.R.Client = foreign
		if foreign.R == nil {
			foreign.R = &clientR{}
		}
		foreign.R.ClientRedirects = append(foreign.R.ClientRedirects, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClientID == foreign.ClientID {
				local.R.Client = foreign
				if foreign.R == nil {
					foreign.R = &clientR{}
				}
				foreign.R.ClientRedirects = append(foreign.R.ClientRedirects, local)
				break
			}
		}
	}

	return nil
}

// SetClient of the clientRedirect to the related item.
// Sets o.R.Client to related.
// Adds o to related.R.ClientRedirects.
func (o *ClientRedirect) SetClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Client) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `client_redirect` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"client_id"}),
		strmangle.WhereClause("`", "`", 0, clientRedirectPrimaryKeyColumns),
	)
	values := []interface{}{related.ClientID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ClientID = related.ClientID
	if o.R == nil {
		o.R = &clientRedirectR{
			Client: related,
		}
	} else {
		o.R.Client = related
	}

	if related.R == nil {
		related.R = &clientR{
			ClientRedirects: ClientRedirectSlice{o},
		}
	} else {
		related.R.ClientRedirects = append(related.R.ClientRedirects, o)
	}

	return nil
}

// ClientRedirects retrieves all the records using an executor.
func ClientRedirects(mods ...qm.QueryMod) clientRedirectQuery {
	mods = append(mods, qm.From("`client_redirect`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`client_redirect`.*"})
	}

	return clientRedirectQuery{q}
}

// FindClientRedirect retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClientRedirect(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ClientRedirect, error) {
	clientRedirectObj := &ClientRedirect{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `client_redirect` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clientRedirectObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from client_redirect")
	}

	if err = clientRedirectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientRedirectObj, err
	}

	return clientRedirectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClientRedirect) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_redirect provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(clientRedirectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientRedirectInsertCacheMut.RLock()
	cache, cached := clientRedirectInsertCache[key]
	clientRedirectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientRedirectAllColumns,
			clientRedirectColumnsWithDefault,
			clientRedirectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `client_redirect` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `client_redirect` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `client_redirect` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, clientRedirectPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into client_redirect")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientRedirectMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for client_redirect")
	}

CacheNoHooks:
	if !cached {
		clientRedirectInsertCacheMut.Lock()
		clientRedirectInsertCache[key] = cache
		clientRedirectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClientRedirect.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClientRedirect) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientRedirectUpdateCacheMut.RLock()
	cache, cached := clientRedirectUpdateCache[key]
	clientRedirectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientRedirectAllColumns,
			clientRedirectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update client_redirect, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `client_redirect` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, clientRedirectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, append(wl, clientRedirectPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update client_redirect row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for client_redirect")
	}

	if !cached {
		clientRedirectUpdateCacheMut.Lock()
		clientRedirectUpdateCache[key] = cache
		clientRedirectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clientRedirectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for client_redirect")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for client_redirect")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientRedirectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientRedirectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `client_redirect` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientRedirectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clientRedirect slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clientRedirect")
	}
	return rowsAff, nil
}

var mySQLClientRedirectUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClientRedirect) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_redirect provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(clientRedirectColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLClientRedirectUniqueColumns, o)

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

	clientRedirectUpsertCacheMut.RLock()
	cache, cached := clientRedirectUpsertCache[key]
	clientRedirectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientRedirectAllColumns,
			clientRedirectColumnsWithDefault,
			clientRedirectColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			clientRedirectAllColumns,
			clientRedirectPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert client_redirect, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`client_redirect`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `client_redirect` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for client_redirect")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientRedirectMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(clientRedirectType, clientRedirectMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for client_redirect")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for client_redirect")
	}

CacheNoHooks:
	if !cached {
		clientRedirectUpsertCacheMut.Lock()
		clientRedirectUpsertCache[key] = cache
		clientRedirectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClientRedirect record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClientRedirect) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClientRedirect provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientRedirectPrimaryKeyMapping)
	sql := "DELETE FROM `client_redirect` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from client_redirect")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for client_redirect")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clientRedirectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clientRedirectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from client_redirect")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_redirect")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientRedirectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientRedirectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientRedirectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `client_redirect` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientRedirectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clientRedirect slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_redirect")
	}

	if len(clientRedirectAfterDeleteHooks) != 0 {
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
func (o *ClientRedirect) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClientRedirect(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientRedirectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientRedirectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientRedirectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `client_redirect`.* FROM `client_redirect` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientRedirectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClientRedirectSlice")
	}

	*o = slice

	return nil
}

// ClientRedirectExists checks if the ClientRedirect row exists.
func ClientRedirectExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `client_redirect` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if client_redirect exists")
	}

	return exists, nil
}

// Exists checks if the ClientRedirect row exists.
func (o *ClientRedirect) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ClientRedirectExists(ctx, exec, o.ID)
}
