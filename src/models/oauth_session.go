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

// OauthSession is an object representing the database table.
type OauthSession struct {
	Code      string      `boil:"code" json:"code" toml:"code" yaml:"code"`
	UserID    string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ClientID  string      `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	State     null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Nonce     null.String `boil:"nonce" json:"nonce,omitempty" toml:"nonce" yaml:"nonce,omitempty"`
	Period    time.Time   `boil:"period" json:"period" toml:"period" yaml:"period"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *oauthSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oauthSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OauthSessionColumns = struct {
	Code      string
	UserID    string
	ClientID  string
	State     string
	Nonce     string
	Period    string
	CreatedAt string
	UpdatedAt string
}{
	Code:      "code",
	UserID:    "user_id",
	ClientID:  "client_id",
	State:     "state",
	Nonce:     "nonce",
	Period:    "period",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var OauthSessionTableColumns = struct {
	Code      string
	UserID    string
	ClientID  string
	State     string
	Nonce     string
	Period    string
	CreatedAt string
	UpdatedAt string
}{
	Code:      "oauth_session.code",
	UserID:    "oauth_session.user_id",
	ClientID:  "oauth_session.client_id",
	State:     "oauth_session.state",
	Nonce:     "oauth_session.nonce",
	Period:    "oauth_session.period",
	CreatedAt: "oauth_session.created_at",
	UpdatedAt: "oauth_session.updated_at",
}

// Generated where

var OauthSessionWhere = struct {
	Code      whereHelperstring
	UserID    whereHelperstring
	ClientID  whereHelperstring
	State     whereHelpernull_String
	Nonce     whereHelpernull_String
	Period    whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	Code:      whereHelperstring{field: "`oauth_session`.`code`"},
	UserID:    whereHelperstring{field: "`oauth_session`.`user_id`"},
	ClientID:  whereHelperstring{field: "`oauth_session`.`client_id`"},
	State:     whereHelpernull_String{field: "`oauth_session`.`state`"},
	Nonce:     whereHelpernull_String{field: "`oauth_session`.`nonce`"},
	Period:    whereHelpertime_Time{field: "`oauth_session`.`period`"},
	CreatedAt: whereHelpertime_Time{field: "`oauth_session`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`oauth_session`.`updated_at`"},
}

// OauthSessionRels is where relationship names are stored.
var OauthSessionRels = struct {
	User   string
	Client string
}{
	User:   "User",
	Client: "Client",
}

// oauthSessionR is where relationships are stored.
type oauthSessionR struct {
	User   *User   `boil:"User" json:"User" toml:"User" yaml:"User"`
	Client *Client `boil:"Client" json:"Client" toml:"Client" yaml:"Client"`
}

// NewStruct creates a new relationship struct
func (*oauthSessionR) NewStruct() *oauthSessionR {
	return &oauthSessionR{}
}

func (r *oauthSessionR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *oauthSessionR) GetClient() *Client {
	if r == nil {
		return nil
	}
	return r.Client
}

// oauthSessionL is where Load methods for each relationship are stored.
type oauthSessionL struct{}

var (
	oauthSessionAllColumns            = []string{"code", "user_id", "client_id", "state", "nonce", "period", "created_at", "updated_at"}
	oauthSessionColumnsWithoutDefault = []string{"code", "user_id", "client_id", "state", "nonce"}
	oauthSessionColumnsWithDefault    = []string{"period", "created_at", "updated_at"}
	oauthSessionPrimaryKeyColumns     = []string{"code"}
	oauthSessionGeneratedColumns      = []string{}
)

type (
	// OauthSessionSlice is an alias for a slice of pointers to OauthSession.
	// This should almost always be used instead of []OauthSession.
	OauthSessionSlice []*OauthSession
	// OauthSessionHook is the signature for custom OauthSession hook methods
	OauthSessionHook func(context.Context, boil.ContextExecutor, *OauthSession) error

	oauthSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oauthSessionType                 = reflect.TypeOf(&OauthSession{})
	oauthSessionMapping              = queries.MakeStructMapping(oauthSessionType)
	oauthSessionPrimaryKeyMapping, _ = queries.BindMapping(oauthSessionType, oauthSessionMapping, oauthSessionPrimaryKeyColumns)
	oauthSessionInsertCacheMut       sync.RWMutex
	oauthSessionInsertCache          = make(map[string]insertCache)
	oauthSessionUpdateCacheMut       sync.RWMutex
	oauthSessionUpdateCache          = make(map[string]updateCache)
	oauthSessionUpsertCacheMut       sync.RWMutex
	oauthSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var oauthSessionAfterSelectHooks []OauthSessionHook

var oauthSessionBeforeInsertHooks []OauthSessionHook
var oauthSessionAfterInsertHooks []OauthSessionHook

var oauthSessionBeforeUpdateHooks []OauthSessionHook
var oauthSessionAfterUpdateHooks []OauthSessionHook

var oauthSessionBeforeDeleteHooks []OauthSessionHook
var oauthSessionAfterDeleteHooks []OauthSessionHook

var oauthSessionBeforeUpsertHooks []OauthSessionHook
var oauthSessionAfterUpsertHooks []OauthSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OauthSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OauthSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OauthSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OauthSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OauthSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OauthSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OauthSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OauthSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OauthSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOauthSessionHook registers your hook function for all future operations.
func AddOauthSessionHook(hookPoint boil.HookPoint, oauthSessionHook OauthSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		oauthSessionAfterSelectHooks = append(oauthSessionAfterSelectHooks, oauthSessionHook)
	case boil.BeforeInsertHook:
		oauthSessionBeforeInsertHooks = append(oauthSessionBeforeInsertHooks, oauthSessionHook)
	case boil.AfterInsertHook:
		oauthSessionAfterInsertHooks = append(oauthSessionAfterInsertHooks, oauthSessionHook)
	case boil.BeforeUpdateHook:
		oauthSessionBeforeUpdateHooks = append(oauthSessionBeforeUpdateHooks, oauthSessionHook)
	case boil.AfterUpdateHook:
		oauthSessionAfterUpdateHooks = append(oauthSessionAfterUpdateHooks, oauthSessionHook)
	case boil.BeforeDeleteHook:
		oauthSessionBeforeDeleteHooks = append(oauthSessionBeforeDeleteHooks, oauthSessionHook)
	case boil.AfterDeleteHook:
		oauthSessionAfterDeleteHooks = append(oauthSessionAfterDeleteHooks, oauthSessionHook)
	case boil.BeforeUpsertHook:
		oauthSessionBeforeUpsertHooks = append(oauthSessionBeforeUpsertHooks, oauthSessionHook)
	case boil.AfterUpsertHook:
		oauthSessionAfterUpsertHooks = append(oauthSessionAfterUpsertHooks, oauthSessionHook)
	}
}

// One returns a single oauthSession record from the query.
func (q oauthSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OauthSession, error) {
	o := &OauthSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for oauth_session")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OauthSession records from the query.
func (q oauthSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (OauthSessionSlice, error) {
	var o []*OauthSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OauthSession slice")
	}

	if len(oauthSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OauthSession records in the query.
func (q oauthSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count oauth_session rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oauthSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if oauth_session exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *OauthSession) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Client pointed to by the foreign key.
func (o *OauthSession) Client(mods ...qm.QueryMod) clientQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`client_id` = ?", o.ClientID),
	}

	queryMods = append(queryMods, mods...)

	return Clients(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (oauthSessionL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOauthSession interface{}, mods queries.Applicator) error {
	var slice []*OauthSession
	var object *OauthSession

	if singular {
		var ok bool
		object, ok = maybeOauthSession.(*OauthSession)
		if !ok {
			object = new(OauthSession)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOauthSession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOauthSession))
			}
		}
	} else {
		s, ok := maybeOauthSession.(*[]*OauthSession)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOauthSession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOauthSession))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &oauthSessionR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &oauthSessionR{}
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
		foreign.R.OauthSessions = append(foreign.R.OauthSessions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OauthSessions = append(foreign.R.OauthSessions, local)
				break
			}
		}
	}

	return nil
}

// LoadClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (oauthSessionL) LoadClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOauthSession interface{}, mods queries.Applicator) error {
	var slice []*OauthSession
	var object *OauthSession

	if singular {
		var ok bool
		object, ok = maybeOauthSession.(*OauthSession)
		if !ok {
			object = new(OauthSession)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOauthSession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOauthSession))
			}
		}
	} else {
		s, ok := maybeOauthSession.(*[]*OauthSession)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOauthSession)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOauthSession))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &oauthSessionR{}
		}
		args = append(args, object.ClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &oauthSessionR{}
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
		foreign.R.OauthSessions = append(foreign.R.OauthSessions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClientID == foreign.ClientID {
				local.R.Client = foreign
				if foreign.R == nil {
					foreign.R = &clientR{}
				}
				foreign.R.OauthSessions = append(foreign.R.OauthSessions, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the oauthSession to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OauthSessions.
func (o *OauthSession) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `oauth_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, oauthSessionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Code}

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
		o.R = &oauthSessionR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OauthSessions: OauthSessionSlice{o},
		}
	} else {
		related.R.OauthSessions = append(related.R.OauthSessions, o)
	}

	return nil
}

// SetClient of the oauthSession to the related item.
// Sets o.R.Client to related.
// Adds o to related.R.OauthSessions.
func (o *OauthSession) SetClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Client) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `oauth_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"client_id"}),
		strmangle.WhereClause("`", "`", 0, oauthSessionPrimaryKeyColumns),
	)
	values := []interface{}{related.ClientID, o.Code}

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
		o.R = &oauthSessionR{
			Client: related,
		}
	} else {
		o.R.Client = related
	}

	if related.R == nil {
		related.R = &clientR{
			OauthSessions: OauthSessionSlice{o},
		}
	} else {
		related.R.OauthSessions = append(related.R.OauthSessions, o)
	}

	return nil
}

// OauthSessions retrieves all the records using an executor.
func OauthSessions(mods ...qm.QueryMod) oauthSessionQuery {
	mods = append(mods, qm.From("`oauth_session`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`oauth_session`.*"})
	}

	return oauthSessionQuery{q}
}

// FindOauthSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOauthSession(ctx context.Context, exec boil.ContextExecutor, code string, selectCols ...string) (*OauthSession, error) {
	oauthSessionObj := &OauthSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `oauth_session` where `code`=?", sel,
	)

	q := queries.Raw(query, code)

	err := q.Bind(ctx, exec, oauthSessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from oauth_session")
	}

	if err = oauthSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return oauthSessionObj, err
	}

	return oauthSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OauthSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauth_session provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(oauthSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oauthSessionInsertCacheMut.RLock()
	cache, cached := oauthSessionInsertCache[key]
	oauthSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oauthSessionAllColumns,
			oauthSessionColumnsWithDefault,
			oauthSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `oauth_session` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `oauth_session` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `oauth_session` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, oauthSessionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into oauth_session")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Code,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for oauth_session")
	}

CacheNoHooks:
	if !cached {
		oauthSessionInsertCacheMut.Lock()
		oauthSessionInsertCache[key] = cache
		oauthSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OauthSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OauthSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	oauthSessionUpdateCacheMut.RLock()
	cache, cached := oauthSessionUpdateCache[key]
	oauthSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oauthSessionAllColumns,
			oauthSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update oauth_session, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `oauth_session` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, oauthSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, append(wl, oauthSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update oauth_session row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for oauth_session")
	}

	if !cached {
		oauthSessionUpdateCacheMut.Lock()
		oauthSessionUpdateCache[key] = cache
		oauthSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q oauthSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for oauth_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for oauth_session")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OauthSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `oauth_session` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in oauthSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all oauthSession")
	}
	return rowsAff, nil
}

var mySQLOauthSessionUniqueColumns = []string{
	"code",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OauthSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauth_session provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(oauthSessionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOauthSessionUniqueColumns, o)

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

	oauthSessionUpsertCacheMut.RLock()
	cache, cached := oauthSessionUpsertCache[key]
	oauthSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oauthSessionAllColumns,
			oauthSessionColumnsWithDefault,
			oauthSessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			oauthSessionAllColumns,
			oauthSessionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert oauth_session, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`oauth_session`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `oauth_session` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for oauth_session")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(oauthSessionType, oauthSessionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for oauth_session")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for oauth_session")
	}

CacheNoHooks:
	if !cached {
		oauthSessionUpsertCacheMut.Lock()
		oauthSessionUpsertCache[key] = cache
		oauthSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OauthSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OauthSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OauthSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oauthSessionPrimaryKeyMapping)
	sql := "DELETE FROM `oauth_session` WHERE `code`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from oauth_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for oauth_session")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oauthSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no oauthSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauth_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauth_session")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OauthSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(oauthSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `oauth_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauthSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauth_session")
	}

	if len(oauthSessionAfterDeleteHooks) != 0 {
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
func (o *OauthSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOauthSession(ctx, exec, o.Code)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OauthSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OauthSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `oauth_session`.* FROM `oauth_session` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OauthSessionSlice")
	}

	*o = slice

	return nil
}

// OauthSessionExists checks if the OauthSession row exists.
func OauthSessionExists(ctx context.Context, exec boil.ContextExecutor, code string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `oauth_session` where `code`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, code)
	}
	row := exec.QueryRowContext(ctx, sql, code)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if oauth_session exists")
	}

	return exists, nil
}

// Exists checks if the OauthSession row exists.
func (o *OauthSession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OauthSessionExists(ctx, exec, o.Code)
}
