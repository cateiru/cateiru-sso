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

// ClientReferrer is an object representing the database table.
type ClientReferrer struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClientID  string    `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	Host      string    `boil:"host" json:"host" toml:"host" yaml:"host"`
	URL       string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *clientReferrerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientReferrerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientReferrerColumns = struct {
	ID        string
	ClientID  string
	Host      string
	URL       string
	CreatedAt string
}{
	ID:        "id",
	ClientID:  "client_id",
	Host:      "host",
	URL:       "url",
	CreatedAt: "created_at",
}

var ClientReferrerTableColumns = struct {
	ID        string
	ClientID  string
	Host      string
	URL       string
	CreatedAt string
}{
	ID:        "client_referrer.id",
	ClientID:  "client_referrer.client_id",
	Host:      "client_referrer.host",
	URL:       "client_referrer.url",
	CreatedAt: "client_referrer.created_at",
}

// Generated where

var ClientReferrerWhere = struct {
	ID        whereHelperuint
	ClientID  whereHelperstring
	Host      whereHelperstring
	URL       whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperuint{field: "`client_referrer`.`id`"},
	ClientID:  whereHelperstring{field: "`client_referrer`.`client_id`"},
	Host:      whereHelperstring{field: "`client_referrer`.`host`"},
	URL:       whereHelperstring{field: "`client_referrer`.`url`"},
	CreatedAt: whereHelpertime_Time{field: "`client_referrer`.`created_at`"},
}

// ClientReferrerRels is where relationship names are stored.
var ClientReferrerRels = struct {
	Client string
}{
	Client: "Client",
}

// clientReferrerR is where relationships are stored.
type clientReferrerR struct {
	Client *Client `boil:"Client" json:"Client" toml:"Client" yaml:"Client"`
}

// NewStruct creates a new relationship struct
func (*clientReferrerR) NewStruct() *clientReferrerR {
	return &clientReferrerR{}
}

func (r *clientReferrerR) GetClient() *Client {
	if r == nil {
		return nil
	}
	return r.Client
}

// clientReferrerL is where Load methods for each relationship are stored.
type clientReferrerL struct{}

var (
	clientReferrerAllColumns            = []string{"id", "client_id", "host", "url", "created_at"}
	clientReferrerColumnsWithoutDefault = []string{"client_id", "host", "url"}
	clientReferrerColumnsWithDefault    = []string{"id", "created_at"}
	clientReferrerPrimaryKeyColumns     = []string{"id"}
	clientReferrerGeneratedColumns      = []string{}
)

type (
	// ClientReferrerSlice is an alias for a slice of pointers to ClientReferrer.
	// This should almost always be used instead of []ClientReferrer.
	ClientReferrerSlice []*ClientReferrer
	// ClientReferrerHook is the signature for custom ClientReferrer hook methods
	ClientReferrerHook func(context.Context, boil.ContextExecutor, *ClientReferrer) error

	clientReferrerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientReferrerType                 = reflect.TypeOf(&ClientReferrer{})
	clientReferrerMapping              = queries.MakeStructMapping(clientReferrerType)
	clientReferrerPrimaryKeyMapping, _ = queries.BindMapping(clientReferrerType, clientReferrerMapping, clientReferrerPrimaryKeyColumns)
	clientReferrerInsertCacheMut       sync.RWMutex
	clientReferrerInsertCache          = make(map[string]insertCache)
	clientReferrerUpdateCacheMut       sync.RWMutex
	clientReferrerUpdateCache          = make(map[string]updateCache)
	clientReferrerUpsertCacheMut       sync.RWMutex
	clientReferrerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientReferrerAfterSelectHooks []ClientReferrerHook

var clientReferrerBeforeInsertHooks []ClientReferrerHook
var clientReferrerAfterInsertHooks []ClientReferrerHook

var clientReferrerBeforeUpdateHooks []ClientReferrerHook
var clientReferrerAfterUpdateHooks []ClientReferrerHook

var clientReferrerBeforeDeleteHooks []ClientReferrerHook
var clientReferrerAfterDeleteHooks []ClientReferrerHook

var clientReferrerBeforeUpsertHooks []ClientReferrerHook
var clientReferrerAfterUpsertHooks []ClientReferrerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClientReferrer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClientReferrer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClientReferrer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClientReferrer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClientReferrer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClientReferrer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClientReferrer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClientReferrer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClientReferrer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientReferrerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientReferrerHook registers your hook function for all future operations.
func AddClientReferrerHook(hookPoint boil.HookPoint, clientReferrerHook ClientReferrerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientReferrerAfterSelectHooks = append(clientReferrerAfterSelectHooks, clientReferrerHook)
	case boil.BeforeInsertHook:
		clientReferrerBeforeInsertHooks = append(clientReferrerBeforeInsertHooks, clientReferrerHook)
	case boil.AfterInsertHook:
		clientReferrerAfterInsertHooks = append(clientReferrerAfterInsertHooks, clientReferrerHook)
	case boil.BeforeUpdateHook:
		clientReferrerBeforeUpdateHooks = append(clientReferrerBeforeUpdateHooks, clientReferrerHook)
	case boil.AfterUpdateHook:
		clientReferrerAfterUpdateHooks = append(clientReferrerAfterUpdateHooks, clientReferrerHook)
	case boil.BeforeDeleteHook:
		clientReferrerBeforeDeleteHooks = append(clientReferrerBeforeDeleteHooks, clientReferrerHook)
	case boil.AfterDeleteHook:
		clientReferrerAfterDeleteHooks = append(clientReferrerAfterDeleteHooks, clientReferrerHook)
	case boil.BeforeUpsertHook:
		clientReferrerBeforeUpsertHooks = append(clientReferrerBeforeUpsertHooks, clientReferrerHook)
	case boil.AfterUpsertHook:
		clientReferrerAfterUpsertHooks = append(clientReferrerAfterUpsertHooks, clientReferrerHook)
	}
}

// One returns a single clientReferrer record from the query.
func (q clientReferrerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClientReferrer, error) {
	o := &ClientReferrer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for client_referrer")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClientReferrer records from the query.
func (q clientReferrerQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientReferrerSlice, error) {
	var o []*ClientReferrer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClientReferrer slice")
	}

	if len(clientReferrerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClientReferrer records in the query.
func (q clientReferrerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count client_referrer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clientReferrerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if client_referrer exists")
	}

	return count > 0, nil
}

// Client pointed to by the foreign key.
func (o *ClientReferrer) Client(mods ...qm.QueryMod) clientQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`client_id` = ?", o.ClientID),
	}

	queryMods = append(queryMods, mods...)

	return Clients(queryMods...)
}

// LoadClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (clientReferrerL) LoadClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClientReferrer interface{}, mods queries.Applicator) error {
	var slice []*ClientReferrer
	var object *ClientReferrer

	if singular {
		var ok bool
		object, ok = maybeClientReferrer.(*ClientReferrer)
		if !ok {
			object = new(ClientReferrer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeClientReferrer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeClientReferrer))
			}
		}
	} else {
		s, ok := maybeClientReferrer.(*[]*ClientReferrer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeClientReferrer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeClientReferrer))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clientReferrerR{}
		}
		args = append(args, object.ClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clientReferrerR{}
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
		foreign.R.ClientReferrers = append(foreign.R.ClientReferrers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClientID == foreign.ClientID {
				local.R.Client = foreign
				if foreign.R == nil {
					foreign.R = &clientR{}
				}
				foreign.R.ClientReferrers = append(foreign.R.ClientReferrers, local)
				break
			}
		}
	}

	return nil
}

// SetClient of the clientReferrer to the related item.
// Sets o.R.Client to related.
// Adds o to related.R.ClientReferrers.
func (o *ClientReferrer) SetClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Client) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `client_referrer` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"client_id"}),
		strmangle.WhereClause("`", "`", 0, clientReferrerPrimaryKeyColumns),
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
		o.R = &clientReferrerR{
			Client: related,
		}
	} else {
		o.R.Client = related
	}

	if related.R == nil {
		related.R = &clientR{
			ClientReferrers: ClientReferrerSlice{o},
		}
	} else {
		related.R.ClientReferrers = append(related.R.ClientReferrers, o)
	}

	return nil
}

// ClientReferrers retrieves all the records using an executor.
func ClientReferrers(mods ...qm.QueryMod) clientReferrerQuery {
	mods = append(mods, qm.From("`client_referrer`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`client_referrer`.*"})
	}

	return clientReferrerQuery{q}
}

// FindClientReferrer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClientReferrer(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ClientReferrer, error) {
	clientReferrerObj := &ClientReferrer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `client_referrer` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clientReferrerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from client_referrer")
	}

	if err = clientReferrerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientReferrerObj, err
	}

	return clientReferrerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClientReferrer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_referrer provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(clientReferrerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientReferrerInsertCacheMut.RLock()
	cache, cached := clientReferrerInsertCache[key]
	clientReferrerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientReferrerAllColumns,
			clientReferrerColumnsWithDefault,
			clientReferrerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `client_referrer` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `client_referrer` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `client_referrer` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, clientReferrerPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into client_referrer")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientReferrerMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for client_referrer")
	}

CacheNoHooks:
	if !cached {
		clientReferrerInsertCacheMut.Lock()
		clientReferrerInsertCache[key] = cache
		clientReferrerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClientReferrer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClientReferrer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientReferrerUpdateCacheMut.RLock()
	cache, cached := clientReferrerUpdateCache[key]
	clientReferrerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientReferrerAllColumns,
			clientReferrerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update client_referrer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `client_referrer` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, clientReferrerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, append(wl, clientReferrerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update client_referrer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for client_referrer")
	}

	if !cached {
		clientReferrerUpdateCacheMut.Lock()
		clientReferrerUpdateCache[key] = cache
		clientReferrerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clientReferrerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for client_referrer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for client_referrer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientReferrerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientReferrerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `client_referrer` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientReferrerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clientReferrer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clientReferrer")
	}
	return rowsAff, nil
}

var mySQLClientReferrerUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClientReferrer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_referrer provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(clientReferrerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLClientReferrerUniqueColumns, o)

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

	clientReferrerUpsertCacheMut.RLock()
	cache, cached := clientReferrerUpsertCache[key]
	clientReferrerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientReferrerAllColumns,
			clientReferrerColumnsWithDefault,
			clientReferrerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			clientReferrerAllColumns,
			clientReferrerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert client_referrer, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`client_referrer`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `client_referrer` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for client_referrer")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientReferrerMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(clientReferrerType, clientReferrerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for client_referrer")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for client_referrer")
	}

CacheNoHooks:
	if !cached {
		clientReferrerUpsertCacheMut.Lock()
		clientReferrerUpsertCache[key] = cache
		clientReferrerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClientReferrer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClientReferrer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClientReferrer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientReferrerPrimaryKeyMapping)
	sql := "DELETE FROM `client_referrer` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from client_referrer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for client_referrer")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clientReferrerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clientReferrerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from client_referrer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_referrer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientReferrerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientReferrerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientReferrerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `client_referrer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientReferrerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clientReferrer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_referrer")
	}

	if len(clientReferrerAfterDeleteHooks) != 0 {
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
func (o *ClientReferrer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClientReferrer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientReferrerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientReferrerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientReferrerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `client_referrer`.* FROM `client_referrer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientReferrerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClientReferrerSlice")
	}

	*o = slice

	return nil
}

// ClientReferrerExists checks if the ClientReferrer row exists.
func ClientReferrerExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `client_referrer` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if client_referrer exists")
	}

	return exists, nil
}

// Exists checks if the ClientReferrer row exists.
func (o *ClientReferrer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ClientReferrerExists(ctx, exec, o.ID)
}
