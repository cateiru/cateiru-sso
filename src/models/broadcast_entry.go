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

// BroadcastEntry is an object representing the database table.
type BroadcastEntry struct {
	ID           uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreateUserID string      `boil:"create_user_id" json:"create_user_id" toml:"create_user_id" yaml:"create_user_id"`
	Title        string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Body         null.String `boil:"body" json:"body,omitempty" toml:"body" yaml:"body,omitempty"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ModifiedAt   time.Time   `boil:"modified_at" json:"modified_at" toml:"modified_at" yaml:"modified_at"`

	R *broadcastEntryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L broadcastEntryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BroadcastEntryColumns = struct {
	ID           string
	CreateUserID string
	Title        string
	Body         string
	CreatedAt    string
	ModifiedAt   string
}{
	ID:           "id",
	CreateUserID: "create_user_id",
	Title:        "title",
	Body:         "body",
	CreatedAt:    "created_at",
	ModifiedAt:   "modified_at",
}

var BroadcastEntryTableColumns = struct {
	ID           string
	CreateUserID string
	Title        string
	Body         string
	CreatedAt    string
	ModifiedAt   string
}{
	ID:           "broadcast_entry.id",
	CreateUserID: "broadcast_entry.create_user_id",
	Title:        "broadcast_entry.title",
	Body:         "broadcast_entry.body",
	CreatedAt:    "broadcast_entry.created_at",
	ModifiedAt:   "broadcast_entry.modified_at",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BroadcastEntryWhere = struct {
	ID           whereHelperuint
	CreateUserID whereHelperstring
	Title        whereHelperstring
	Body         whereHelpernull_String
	CreatedAt    whereHelpertime_Time
	ModifiedAt   whereHelpertime_Time
}{
	ID:           whereHelperuint{field: "`broadcast_entry`.`id`"},
	CreateUserID: whereHelperstring{field: "`broadcast_entry`.`create_user_id`"},
	Title:        whereHelperstring{field: "`broadcast_entry`.`title`"},
	Body:         whereHelpernull_String{field: "`broadcast_entry`.`body`"},
	CreatedAt:    whereHelpertime_Time{field: "`broadcast_entry`.`created_at`"},
	ModifiedAt:   whereHelpertime_Time{field: "`broadcast_entry`.`modified_at`"},
}

// BroadcastEntryRels is where relationship names are stored.
var BroadcastEntryRels = struct {
	EntryBroadcastNotices string
}{
	EntryBroadcastNotices: "EntryBroadcastNotices",
}

// broadcastEntryR is where relationships are stored.
type broadcastEntryR struct {
	EntryBroadcastNotices BroadcastNoticeSlice `boil:"EntryBroadcastNotices" json:"EntryBroadcastNotices" toml:"EntryBroadcastNotices" yaml:"EntryBroadcastNotices"`
}

// NewStruct creates a new relationship struct
func (*broadcastEntryR) NewStruct() *broadcastEntryR {
	return &broadcastEntryR{}
}

func (r *broadcastEntryR) GetEntryBroadcastNotices() BroadcastNoticeSlice {
	if r == nil {
		return nil
	}
	return r.EntryBroadcastNotices
}

// broadcastEntryL is where Load methods for each relationship are stored.
type broadcastEntryL struct{}

var (
	broadcastEntryAllColumns            = []string{"id", "create_user_id", "title", "body", "created_at", "modified_at"}
	broadcastEntryColumnsWithoutDefault = []string{"create_user_id", "title", "body"}
	broadcastEntryColumnsWithDefault    = []string{"id", "created_at", "modified_at"}
	broadcastEntryPrimaryKeyColumns     = []string{"id"}
	broadcastEntryGeneratedColumns      = []string{}
)

type (
	// BroadcastEntrySlice is an alias for a slice of pointers to BroadcastEntry.
	// This should almost always be used instead of []BroadcastEntry.
	BroadcastEntrySlice []*BroadcastEntry
	// BroadcastEntryHook is the signature for custom BroadcastEntry hook methods
	BroadcastEntryHook func(context.Context, boil.ContextExecutor, *BroadcastEntry) error

	broadcastEntryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	broadcastEntryType                 = reflect.TypeOf(&BroadcastEntry{})
	broadcastEntryMapping              = queries.MakeStructMapping(broadcastEntryType)
	broadcastEntryPrimaryKeyMapping, _ = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, broadcastEntryPrimaryKeyColumns)
	broadcastEntryInsertCacheMut       sync.RWMutex
	broadcastEntryInsertCache          = make(map[string]insertCache)
	broadcastEntryUpdateCacheMut       sync.RWMutex
	broadcastEntryUpdateCache          = make(map[string]updateCache)
	broadcastEntryUpsertCacheMut       sync.RWMutex
	broadcastEntryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var broadcastEntryAfterSelectHooks []BroadcastEntryHook

var broadcastEntryBeforeInsertHooks []BroadcastEntryHook
var broadcastEntryAfterInsertHooks []BroadcastEntryHook

var broadcastEntryBeforeUpdateHooks []BroadcastEntryHook
var broadcastEntryAfterUpdateHooks []BroadcastEntryHook

var broadcastEntryBeforeDeleteHooks []BroadcastEntryHook
var broadcastEntryAfterDeleteHooks []BroadcastEntryHook

var broadcastEntryBeforeUpsertHooks []BroadcastEntryHook
var broadcastEntryAfterUpsertHooks []BroadcastEntryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BroadcastEntry) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BroadcastEntry) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BroadcastEntry) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BroadcastEntry) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BroadcastEntry) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BroadcastEntry) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BroadcastEntry) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BroadcastEntry) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BroadcastEntry) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range broadcastEntryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBroadcastEntryHook registers your hook function for all future operations.
func AddBroadcastEntryHook(hookPoint boil.HookPoint, broadcastEntryHook BroadcastEntryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		broadcastEntryAfterSelectHooks = append(broadcastEntryAfterSelectHooks, broadcastEntryHook)
	case boil.BeforeInsertHook:
		broadcastEntryBeforeInsertHooks = append(broadcastEntryBeforeInsertHooks, broadcastEntryHook)
	case boil.AfterInsertHook:
		broadcastEntryAfterInsertHooks = append(broadcastEntryAfterInsertHooks, broadcastEntryHook)
	case boil.BeforeUpdateHook:
		broadcastEntryBeforeUpdateHooks = append(broadcastEntryBeforeUpdateHooks, broadcastEntryHook)
	case boil.AfterUpdateHook:
		broadcastEntryAfterUpdateHooks = append(broadcastEntryAfterUpdateHooks, broadcastEntryHook)
	case boil.BeforeDeleteHook:
		broadcastEntryBeforeDeleteHooks = append(broadcastEntryBeforeDeleteHooks, broadcastEntryHook)
	case boil.AfterDeleteHook:
		broadcastEntryAfterDeleteHooks = append(broadcastEntryAfterDeleteHooks, broadcastEntryHook)
	case boil.BeforeUpsertHook:
		broadcastEntryBeforeUpsertHooks = append(broadcastEntryBeforeUpsertHooks, broadcastEntryHook)
	case boil.AfterUpsertHook:
		broadcastEntryAfterUpsertHooks = append(broadcastEntryAfterUpsertHooks, broadcastEntryHook)
	}
}

// One returns a single broadcastEntry record from the query.
func (q broadcastEntryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BroadcastEntry, error) {
	o := &BroadcastEntry{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for broadcast_entry")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BroadcastEntry records from the query.
func (q broadcastEntryQuery) All(ctx context.Context, exec boil.ContextExecutor) (BroadcastEntrySlice, error) {
	var o []*BroadcastEntry

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BroadcastEntry slice")
	}

	if len(broadcastEntryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BroadcastEntry records in the query.
func (q broadcastEntryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count broadcast_entry rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q broadcastEntryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if broadcast_entry exists")
	}

	return count > 0, nil
}

// EntryBroadcastNotices retrieves all the broadcast_notice's BroadcastNotices with an executor via entry_id column.
func (o *BroadcastEntry) EntryBroadcastNotices(mods ...qm.QueryMod) broadcastNoticeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`broadcast_notice`.`entry_id`=?", o.ID),
	)

	return BroadcastNotices(queryMods...)
}

// LoadEntryBroadcastNotices allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (broadcastEntryL) LoadEntryBroadcastNotices(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBroadcastEntry interface{}, mods queries.Applicator) error {
	var slice []*BroadcastEntry
	var object *BroadcastEntry

	if singular {
		var ok bool
		object, ok = maybeBroadcastEntry.(*BroadcastEntry)
		if !ok {
			object = new(BroadcastEntry)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBroadcastEntry)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBroadcastEntry))
			}
		}
	} else {
		s, ok := maybeBroadcastEntry.(*[]*BroadcastEntry)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBroadcastEntry)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBroadcastEntry))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &broadcastEntryR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &broadcastEntryR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`broadcast_notice`),
		qm.WhereIn(`broadcast_notice.entry_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load broadcast_notice")
	}

	var resultSlice []*BroadcastNotice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice broadcast_notice")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on broadcast_notice")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for broadcast_notice")
	}

	if len(broadcastNoticeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.EntryBroadcastNotices = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &broadcastNoticeR{}
			}
			foreign.R.Entry = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EntryID {
				local.R.EntryBroadcastNotices = append(local.R.EntryBroadcastNotices, foreign)
				if foreign.R == nil {
					foreign.R = &broadcastNoticeR{}
				}
				foreign.R.Entry = local
				break
			}
		}
	}

	return nil
}

// AddEntryBroadcastNotices adds the given related objects to the existing relationships
// of the broadcast_entry, optionally inserting them as new records.
// Appends related to o.R.EntryBroadcastNotices.
// Sets related.R.Entry appropriately.
func (o *BroadcastEntry) AddEntryBroadcastNotices(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BroadcastNotice) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EntryID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `broadcast_notice` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"entry_id"}),
				strmangle.WhereClause("`", "`", 0, broadcastNoticePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EntryID = o.ID
		}
	}

	if o.R == nil {
		o.R = &broadcastEntryR{
			EntryBroadcastNotices: related,
		}
	} else {
		o.R.EntryBroadcastNotices = append(o.R.EntryBroadcastNotices, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &broadcastNoticeR{
				Entry: o,
			}
		} else {
			rel.R.Entry = o
		}
	}
	return nil
}

// BroadcastEntries retrieves all the records using an executor.
func BroadcastEntries(mods ...qm.QueryMod) broadcastEntryQuery {
	mods = append(mods, qm.From("`broadcast_entry`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`broadcast_entry`.*"})
	}

	return broadcastEntryQuery{q}
}

// FindBroadcastEntry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBroadcastEntry(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*BroadcastEntry, error) {
	broadcastEntryObj := &BroadcastEntry{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `broadcast_entry` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, broadcastEntryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from broadcast_entry")
	}

	if err = broadcastEntryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return broadcastEntryObj, err
	}

	return broadcastEntryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BroadcastEntry) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no broadcast_entry provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(broadcastEntryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	broadcastEntryInsertCacheMut.RLock()
	cache, cached := broadcastEntryInsertCache[key]
	broadcastEntryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			broadcastEntryAllColumns,
			broadcastEntryColumnsWithDefault,
			broadcastEntryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `broadcast_entry` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `broadcast_entry` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `broadcast_entry` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, broadcastEntryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into broadcast_entry")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == broadcastEntryMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for broadcast_entry")
	}

CacheNoHooks:
	if !cached {
		broadcastEntryInsertCacheMut.Lock()
		broadcastEntryInsertCache[key] = cache
		broadcastEntryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BroadcastEntry.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BroadcastEntry) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	broadcastEntryUpdateCacheMut.RLock()
	cache, cached := broadcastEntryUpdateCache[key]
	broadcastEntryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			broadcastEntryAllColumns,
			broadcastEntryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update broadcast_entry, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `broadcast_entry` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, broadcastEntryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, append(wl, broadcastEntryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update broadcast_entry row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for broadcast_entry")
	}

	if !cached {
		broadcastEntryUpdateCacheMut.Lock()
		broadcastEntryUpdateCache[key] = cache
		broadcastEntryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q broadcastEntryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for broadcast_entry")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for broadcast_entry")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BroadcastEntrySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), broadcastEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `broadcast_entry` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, broadcastEntryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in broadcastEntry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all broadcastEntry")
	}
	return rowsAff, nil
}

var mySQLBroadcastEntryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BroadcastEntry) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no broadcast_entry provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(broadcastEntryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBroadcastEntryUniqueColumns, o)

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

	broadcastEntryUpsertCacheMut.RLock()
	cache, cached := broadcastEntryUpsertCache[key]
	broadcastEntryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			broadcastEntryAllColumns,
			broadcastEntryColumnsWithDefault,
			broadcastEntryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			broadcastEntryAllColumns,
			broadcastEntryPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert broadcast_entry, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`broadcast_entry`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `broadcast_entry` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for broadcast_entry")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == broadcastEntryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(broadcastEntryType, broadcastEntryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for broadcast_entry")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for broadcast_entry")
	}

CacheNoHooks:
	if !cached {
		broadcastEntryUpsertCacheMut.Lock()
		broadcastEntryUpsertCache[key] = cache
		broadcastEntryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BroadcastEntry record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BroadcastEntry) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BroadcastEntry provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), broadcastEntryPrimaryKeyMapping)
	sql := "DELETE FROM `broadcast_entry` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from broadcast_entry")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for broadcast_entry")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q broadcastEntryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no broadcastEntryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from broadcast_entry")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for broadcast_entry")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BroadcastEntrySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(broadcastEntryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), broadcastEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `broadcast_entry` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, broadcastEntryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from broadcastEntry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for broadcast_entry")
	}

	if len(broadcastEntryAfterDeleteHooks) != 0 {
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
func (o *BroadcastEntry) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBroadcastEntry(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BroadcastEntrySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BroadcastEntrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), broadcastEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `broadcast_entry`.* FROM `broadcast_entry` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, broadcastEntryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BroadcastEntrySlice")
	}

	*o = slice

	return nil
}

// BroadcastEntryExists checks if the BroadcastEntry row exists.
func BroadcastEntryExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `broadcast_entry` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if broadcast_entry exists")
	}

	return exists, nil
}

// Exists checks if the BroadcastEntry row exists.
func (o *BroadcastEntry) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BroadcastEntryExists(ctx, exec, o.ID)
}
