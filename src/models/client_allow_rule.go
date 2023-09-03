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

// ClientAllowRule is an object representing the database table.
type ClientAllowRule struct {
	ID          uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClientID    string      `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	UserID      null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	EmailDomain null.String `boil:"email_domain" json:"email_domain,omitempty" toml:"email_domain" yaml:"email_domain,omitempty"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *clientAllowRuleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientAllowRuleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientAllowRuleColumns = struct {
	ID          string
	ClientID    string
	UserID      string
	EmailDomain string
	CreatedAt   string
}{
	ID:          "id",
	ClientID:    "client_id",
	UserID:      "user_id",
	EmailDomain: "email_domain",
	CreatedAt:   "created_at",
}

var ClientAllowRuleTableColumns = struct {
	ID          string
	ClientID    string
	UserID      string
	EmailDomain string
	CreatedAt   string
}{
	ID:          "client_allow_rule.id",
	ClientID:    "client_allow_rule.client_id",
	UserID:      "client_allow_rule.user_id",
	EmailDomain: "client_allow_rule.email_domain",
	CreatedAt:   "client_allow_rule.created_at",
}

// Generated where

var ClientAllowRuleWhere = struct {
	ID          whereHelperuint
	ClientID    whereHelperstring
	UserID      whereHelpernull_String
	EmailDomain whereHelpernull_String
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperuint{field: "`client_allow_rule`.`id`"},
	ClientID:    whereHelperstring{field: "`client_allow_rule`.`client_id`"},
	UserID:      whereHelpernull_String{field: "`client_allow_rule`.`user_id`"},
	EmailDomain: whereHelpernull_String{field: "`client_allow_rule`.`email_domain`"},
	CreatedAt:   whereHelpertime_Time{field: "`client_allow_rule`.`created_at`"},
}

// ClientAllowRuleRels is where relationship names are stored.
var ClientAllowRuleRels = struct {
	Client string
}{
	Client: "Client",
}

// clientAllowRuleR is where relationships are stored.
type clientAllowRuleR struct {
	Client *Client `boil:"Client" json:"Client" toml:"Client" yaml:"Client"`
}

// NewStruct creates a new relationship struct
func (*clientAllowRuleR) NewStruct() *clientAllowRuleR {
	return &clientAllowRuleR{}
}

func (r *clientAllowRuleR) GetClient() *Client {
	if r == nil {
		return nil
	}
	return r.Client
}

// clientAllowRuleL is where Load methods for each relationship are stored.
type clientAllowRuleL struct{}

var (
	clientAllowRuleAllColumns            = []string{"id", "client_id", "user_id", "email_domain", "created_at"}
	clientAllowRuleColumnsWithoutDefault = []string{"client_id", "user_id", "email_domain"}
	clientAllowRuleColumnsWithDefault    = []string{"id", "created_at"}
	clientAllowRulePrimaryKeyColumns     = []string{"id"}
	clientAllowRuleGeneratedColumns      = []string{}
)

type (
	// ClientAllowRuleSlice is an alias for a slice of pointers to ClientAllowRule.
	// This should almost always be used instead of []ClientAllowRule.
	ClientAllowRuleSlice []*ClientAllowRule
	// ClientAllowRuleHook is the signature for custom ClientAllowRule hook methods
	ClientAllowRuleHook func(context.Context, boil.ContextExecutor, *ClientAllowRule) error

	clientAllowRuleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientAllowRuleType                 = reflect.TypeOf(&ClientAllowRule{})
	clientAllowRuleMapping              = queries.MakeStructMapping(clientAllowRuleType)
	clientAllowRulePrimaryKeyMapping, _ = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, clientAllowRulePrimaryKeyColumns)
	clientAllowRuleInsertCacheMut       sync.RWMutex
	clientAllowRuleInsertCache          = make(map[string]insertCache)
	clientAllowRuleUpdateCacheMut       sync.RWMutex
	clientAllowRuleUpdateCache          = make(map[string]updateCache)
	clientAllowRuleUpsertCacheMut       sync.RWMutex
	clientAllowRuleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientAllowRuleAfterSelectHooks []ClientAllowRuleHook

var clientAllowRuleBeforeInsertHooks []ClientAllowRuleHook
var clientAllowRuleAfterInsertHooks []ClientAllowRuleHook

var clientAllowRuleBeforeUpdateHooks []ClientAllowRuleHook
var clientAllowRuleAfterUpdateHooks []ClientAllowRuleHook

var clientAllowRuleBeforeDeleteHooks []ClientAllowRuleHook
var clientAllowRuleAfterDeleteHooks []ClientAllowRuleHook

var clientAllowRuleBeforeUpsertHooks []ClientAllowRuleHook
var clientAllowRuleAfterUpsertHooks []ClientAllowRuleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClientAllowRule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClientAllowRule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClientAllowRule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClientAllowRule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClientAllowRule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClientAllowRule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClientAllowRule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClientAllowRule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClientAllowRule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAllowRuleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientAllowRuleHook registers your hook function for all future operations.
func AddClientAllowRuleHook(hookPoint boil.HookPoint, clientAllowRuleHook ClientAllowRuleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientAllowRuleAfterSelectHooks = append(clientAllowRuleAfterSelectHooks, clientAllowRuleHook)
	case boil.BeforeInsertHook:
		clientAllowRuleBeforeInsertHooks = append(clientAllowRuleBeforeInsertHooks, clientAllowRuleHook)
	case boil.AfterInsertHook:
		clientAllowRuleAfterInsertHooks = append(clientAllowRuleAfterInsertHooks, clientAllowRuleHook)
	case boil.BeforeUpdateHook:
		clientAllowRuleBeforeUpdateHooks = append(clientAllowRuleBeforeUpdateHooks, clientAllowRuleHook)
	case boil.AfterUpdateHook:
		clientAllowRuleAfterUpdateHooks = append(clientAllowRuleAfterUpdateHooks, clientAllowRuleHook)
	case boil.BeforeDeleteHook:
		clientAllowRuleBeforeDeleteHooks = append(clientAllowRuleBeforeDeleteHooks, clientAllowRuleHook)
	case boil.AfterDeleteHook:
		clientAllowRuleAfterDeleteHooks = append(clientAllowRuleAfterDeleteHooks, clientAllowRuleHook)
	case boil.BeforeUpsertHook:
		clientAllowRuleBeforeUpsertHooks = append(clientAllowRuleBeforeUpsertHooks, clientAllowRuleHook)
	case boil.AfterUpsertHook:
		clientAllowRuleAfterUpsertHooks = append(clientAllowRuleAfterUpsertHooks, clientAllowRuleHook)
	}
}

// One returns a single clientAllowRule record from the query.
func (q clientAllowRuleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClientAllowRule, error) {
	o := &ClientAllowRule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for client_allow_rule")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClientAllowRule records from the query.
func (q clientAllowRuleQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientAllowRuleSlice, error) {
	var o []*ClientAllowRule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClientAllowRule slice")
	}

	if len(clientAllowRuleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClientAllowRule records in the query.
func (q clientAllowRuleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count client_allow_rule rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clientAllowRuleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if client_allow_rule exists")
	}

	return count > 0, nil
}

// Client pointed to by the foreign key.
func (o *ClientAllowRule) Client(mods ...qm.QueryMod) clientQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`client_id` = ?", o.ClientID),
	}

	queryMods = append(queryMods, mods...)

	return Clients(queryMods...)
}

// LoadClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (clientAllowRuleL) LoadClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClientAllowRule interface{}, mods queries.Applicator) error {
	var slice []*ClientAllowRule
	var object *ClientAllowRule

	if singular {
		var ok bool
		object, ok = maybeClientAllowRule.(*ClientAllowRule)
		if !ok {
			object = new(ClientAllowRule)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeClientAllowRule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeClientAllowRule))
			}
		}
	} else {
		s, ok := maybeClientAllowRule.(*[]*ClientAllowRule)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeClientAllowRule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeClientAllowRule))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clientAllowRuleR{}
		}
		args = append(args, object.ClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clientAllowRuleR{}
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
		foreign.R.ClientAllowRules = append(foreign.R.ClientAllowRules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClientID == foreign.ClientID {
				local.R.Client = foreign
				if foreign.R == nil {
					foreign.R = &clientR{}
				}
				foreign.R.ClientAllowRules = append(foreign.R.ClientAllowRules, local)
				break
			}
		}
	}

	return nil
}

// SetClient of the clientAllowRule to the related item.
// Sets o.R.Client to related.
// Adds o to related.R.ClientAllowRules.
func (o *ClientAllowRule) SetClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Client) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `client_allow_rule` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"client_id"}),
		strmangle.WhereClause("`", "`", 0, clientAllowRulePrimaryKeyColumns),
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
		o.R = &clientAllowRuleR{
			Client: related,
		}
	} else {
		o.R.Client = related
	}

	if related.R == nil {
		related.R = &clientR{
			ClientAllowRules: ClientAllowRuleSlice{o},
		}
	} else {
		related.R.ClientAllowRules = append(related.R.ClientAllowRules, o)
	}

	return nil
}

// ClientAllowRules retrieves all the records using an executor.
func ClientAllowRules(mods ...qm.QueryMod) clientAllowRuleQuery {
	mods = append(mods, qm.From("`client_allow_rule`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`client_allow_rule`.*"})
	}

	return clientAllowRuleQuery{q}
}

// FindClientAllowRule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClientAllowRule(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ClientAllowRule, error) {
	clientAllowRuleObj := &ClientAllowRule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `client_allow_rule` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clientAllowRuleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from client_allow_rule")
	}

	if err = clientAllowRuleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientAllowRuleObj, err
	}

	return clientAllowRuleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClientAllowRule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_allow_rule provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(clientAllowRuleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientAllowRuleInsertCacheMut.RLock()
	cache, cached := clientAllowRuleInsertCache[key]
	clientAllowRuleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientAllowRuleAllColumns,
			clientAllowRuleColumnsWithDefault,
			clientAllowRuleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `client_allow_rule` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `client_allow_rule` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `client_allow_rule` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, clientAllowRulePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into client_allow_rule")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientAllowRuleMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for client_allow_rule")
	}

CacheNoHooks:
	if !cached {
		clientAllowRuleInsertCacheMut.Lock()
		clientAllowRuleInsertCache[key] = cache
		clientAllowRuleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClientAllowRule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClientAllowRule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientAllowRuleUpdateCacheMut.RLock()
	cache, cached := clientAllowRuleUpdateCache[key]
	clientAllowRuleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientAllowRuleAllColumns,
			clientAllowRulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update client_allow_rule, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `client_allow_rule` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, clientAllowRulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, append(wl, clientAllowRulePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update client_allow_rule row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for client_allow_rule")
	}

	if !cached {
		clientAllowRuleUpdateCacheMut.Lock()
		clientAllowRuleUpdateCache[key] = cache
		clientAllowRuleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clientAllowRuleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for client_allow_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for client_allow_rule")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientAllowRuleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientAllowRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `client_allow_rule` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientAllowRulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clientAllowRule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clientAllowRule")
	}
	return rowsAff, nil
}

var mySQLClientAllowRuleUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClientAllowRule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_allow_rule provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(clientAllowRuleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLClientAllowRuleUniqueColumns, o)

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

	clientAllowRuleUpsertCacheMut.RLock()
	cache, cached := clientAllowRuleUpsertCache[key]
	clientAllowRuleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientAllowRuleAllColumns,
			clientAllowRuleColumnsWithDefault,
			clientAllowRuleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			clientAllowRuleAllColumns,
			clientAllowRulePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert client_allow_rule, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`client_allow_rule`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `client_allow_rule` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for client_allow_rule")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientAllowRuleMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(clientAllowRuleType, clientAllowRuleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for client_allow_rule")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for client_allow_rule")
	}

CacheNoHooks:
	if !cached {
		clientAllowRuleUpsertCacheMut.Lock()
		clientAllowRuleUpsertCache[key] = cache
		clientAllowRuleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClientAllowRule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClientAllowRule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClientAllowRule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientAllowRulePrimaryKeyMapping)
	sql := "DELETE FROM `client_allow_rule` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from client_allow_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for client_allow_rule")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clientAllowRuleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clientAllowRuleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from client_allow_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_allow_rule")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientAllowRuleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientAllowRuleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientAllowRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `client_allow_rule` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientAllowRulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clientAllowRule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_allow_rule")
	}

	if len(clientAllowRuleAfterDeleteHooks) != 0 {
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
func (o *ClientAllowRule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClientAllowRule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientAllowRuleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientAllowRuleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientAllowRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `client_allow_rule`.* FROM `client_allow_rule` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientAllowRulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClientAllowRuleSlice")
	}

	*o = slice

	return nil
}

// ClientAllowRuleExists checks if the ClientAllowRule row exists.
func ClientAllowRuleExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `client_allow_rule` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if client_allow_rule exists")
	}

	return exists, nil
}

// Exists checks if the ClientAllowRule row exists.
func (o *ClientAllowRule) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ClientAllowRuleExists(ctx, exec, o.ID)
}
