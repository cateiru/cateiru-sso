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

// ClientQuiz is an object representing the database table.
type ClientQuiz struct {
	ID           uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClientID     []byte    `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	Title        string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	AnswerRegexp string    `boil:"answer_regexp" json:"answer_regexp" toml:"answer_regexp" yaml:"answer_regexp"`
	Choices      null.JSON `boil:"choices" json:"choices,omitempty" toml:"choices" yaml:"choices,omitempty"`
	Created      time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`
	Modified     time.Time `boil:"modified" json:"modified" toml:"modified" yaml:"modified"`

	R *clientQuizR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientQuizL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientQuizColumns = struct {
	ID           string
	ClientID     string
	Title        string
	AnswerRegexp string
	Choices      string
	Created      string
	Modified     string
}{
	ID:           "id",
	ClientID:     "client_id",
	Title:        "title",
	AnswerRegexp: "answer_regexp",
	Choices:      "choices",
	Created:      "created",
	Modified:     "modified",
}

var ClientQuizTableColumns = struct {
	ID           string
	ClientID     string
	Title        string
	AnswerRegexp string
	Choices      string
	Created      string
	Modified     string
}{
	ID:           "client_quiz.id",
	ClientID:     "client_quiz.client_id",
	Title:        "client_quiz.title",
	AnswerRegexp: "client_quiz.answer_regexp",
	Choices:      "client_quiz.choices",
	Created:      "client_quiz.created",
	Modified:     "client_quiz.modified",
}

// Generated where

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ClientQuizWhere = struct {
	ID           whereHelperuint
	ClientID     whereHelper__byte
	Title        whereHelperstring
	AnswerRegexp whereHelperstring
	Choices      whereHelpernull_JSON
	Created      whereHelpertime_Time
	Modified     whereHelpertime_Time
}{
	ID:           whereHelperuint{field: "`client_quiz`.`id`"},
	ClientID:     whereHelper__byte{field: "`client_quiz`.`client_id`"},
	Title:        whereHelperstring{field: "`client_quiz`.`title`"},
	AnswerRegexp: whereHelperstring{field: "`client_quiz`.`answer_regexp`"},
	Choices:      whereHelpernull_JSON{field: "`client_quiz`.`choices`"},
	Created:      whereHelpertime_Time{field: "`client_quiz`.`created`"},
	Modified:     whereHelpertime_Time{field: "`client_quiz`.`modified`"},
}

// ClientQuizRels is where relationship names are stored.
var ClientQuizRels = struct {
}{}

// clientQuizR is where relationships are stored.
type clientQuizR struct {
}

// NewStruct creates a new relationship struct
func (*clientQuizR) NewStruct() *clientQuizR {
	return &clientQuizR{}
}

// clientQuizL is where Load methods for each relationship are stored.
type clientQuizL struct{}

var (
	clientQuizAllColumns            = []string{"id", "client_id", "title", "answer_regexp", "choices", "created", "modified"}
	clientQuizColumnsWithoutDefault = []string{"client_id", "title", "answer_regexp", "choices"}
	clientQuizColumnsWithDefault    = []string{"id", "created", "modified"}
	clientQuizPrimaryKeyColumns     = []string{"id"}
	clientQuizGeneratedColumns      = []string{}
)

type (
	// ClientQuizSlice is an alias for a slice of pointers to ClientQuiz.
	// This should almost always be used instead of []ClientQuiz.
	ClientQuizSlice []*ClientQuiz
	// ClientQuizHook is the signature for custom ClientQuiz hook methods
	ClientQuizHook func(context.Context, boil.ContextExecutor, *ClientQuiz) error

	clientQuizQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientQuizType                 = reflect.TypeOf(&ClientQuiz{})
	clientQuizMapping              = queries.MakeStructMapping(clientQuizType)
	clientQuizPrimaryKeyMapping, _ = queries.BindMapping(clientQuizType, clientQuizMapping, clientQuizPrimaryKeyColumns)
	clientQuizInsertCacheMut       sync.RWMutex
	clientQuizInsertCache          = make(map[string]insertCache)
	clientQuizUpdateCacheMut       sync.RWMutex
	clientQuizUpdateCache          = make(map[string]updateCache)
	clientQuizUpsertCacheMut       sync.RWMutex
	clientQuizUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientQuizAfterSelectHooks []ClientQuizHook

var clientQuizBeforeInsertHooks []ClientQuizHook
var clientQuizAfterInsertHooks []ClientQuizHook

var clientQuizBeforeUpdateHooks []ClientQuizHook
var clientQuizAfterUpdateHooks []ClientQuizHook

var clientQuizBeforeDeleteHooks []ClientQuizHook
var clientQuizAfterDeleteHooks []ClientQuizHook

var clientQuizBeforeUpsertHooks []ClientQuizHook
var clientQuizAfterUpsertHooks []ClientQuizHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClientQuiz) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClientQuiz) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClientQuiz) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClientQuiz) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClientQuiz) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClientQuiz) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClientQuiz) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClientQuiz) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClientQuiz) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientQuizAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientQuizHook registers your hook function for all future operations.
func AddClientQuizHook(hookPoint boil.HookPoint, clientQuizHook ClientQuizHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientQuizAfterSelectHooks = append(clientQuizAfterSelectHooks, clientQuizHook)
	case boil.BeforeInsertHook:
		clientQuizBeforeInsertHooks = append(clientQuizBeforeInsertHooks, clientQuizHook)
	case boil.AfterInsertHook:
		clientQuizAfterInsertHooks = append(clientQuizAfterInsertHooks, clientQuizHook)
	case boil.BeforeUpdateHook:
		clientQuizBeforeUpdateHooks = append(clientQuizBeforeUpdateHooks, clientQuizHook)
	case boil.AfterUpdateHook:
		clientQuizAfterUpdateHooks = append(clientQuizAfterUpdateHooks, clientQuizHook)
	case boil.BeforeDeleteHook:
		clientQuizBeforeDeleteHooks = append(clientQuizBeforeDeleteHooks, clientQuizHook)
	case boil.AfterDeleteHook:
		clientQuizAfterDeleteHooks = append(clientQuizAfterDeleteHooks, clientQuizHook)
	case boil.BeforeUpsertHook:
		clientQuizBeforeUpsertHooks = append(clientQuizBeforeUpsertHooks, clientQuizHook)
	case boil.AfterUpsertHook:
		clientQuizAfterUpsertHooks = append(clientQuizAfterUpsertHooks, clientQuizHook)
	}
}

// One returns a single clientQuiz record from the query.
func (q clientQuizQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClientQuiz, error) {
	o := &ClientQuiz{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for client_quiz")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClientQuiz records from the query.
func (q clientQuizQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientQuizSlice, error) {
	var o []*ClientQuiz

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClientQuiz slice")
	}

	if len(clientQuizAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClientQuiz records in the query.
func (q clientQuizQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count client_quiz rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clientQuizQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if client_quiz exists")
	}

	return count > 0, nil
}

// ClientQuizzes retrieves all the records using an executor.
func ClientQuizzes(mods ...qm.QueryMod) clientQuizQuery {
	mods = append(mods, qm.From("`client_quiz`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`client_quiz`.*"})
	}

	return clientQuizQuery{q}
}

// FindClientQuiz retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClientQuiz(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ClientQuiz, error) {
	clientQuizObj := &ClientQuiz{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `client_quiz` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clientQuizObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from client_quiz")
	}

	if err = clientQuizObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientQuizObj, err
	}

	return clientQuizObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClientQuiz) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_quiz provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clientQuizColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientQuizInsertCacheMut.RLock()
	cache, cached := clientQuizInsertCache[key]
	clientQuizInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientQuizAllColumns,
			clientQuizColumnsWithDefault,
			clientQuizColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientQuizType, clientQuizMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientQuizType, clientQuizMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `client_quiz` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `client_quiz` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `client_quiz` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, clientQuizPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into client_quiz")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientQuizMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for client_quiz")
	}

CacheNoHooks:
	if !cached {
		clientQuizInsertCacheMut.Lock()
		clientQuizInsertCache[key] = cache
		clientQuizInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClientQuiz.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClientQuiz) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientQuizUpdateCacheMut.RLock()
	cache, cached := clientQuizUpdateCache[key]
	clientQuizUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientQuizAllColumns,
			clientQuizPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update client_quiz, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `client_quiz` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, clientQuizPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientQuizType, clientQuizMapping, append(wl, clientQuizPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update client_quiz row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for client_quiz")
	}

	if !cached {
		clientQuizUpdateCacheMut.Lock()
		clientQuizUpdateCache[key] = cache
		clientQuizUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clientQuizQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for client_quiz")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for client_quiz")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientQuizSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientQuizPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `client_quiz` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientQuizPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clientQuiz slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clientQuiz")
	}
	return rowsAff, nil
}

var mySQLClientQuizUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClientQuiz) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no client_quiz provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clientQuizColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLClientQuizUniqueColumns, o)

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

	clientQuizUpsertCacheMut.RLock()
	cache, cached := clientQuizUpsertCache[key]
	clientQuizUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientQuizAllColumns,
			clientQuizColumnsWithDefault,
			clientQuizColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			clientQuizAllColumns,
			clientQuizPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert client_quiz, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`client_quiz`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `client_quiz` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(clientQuizType, clientQuizMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientQuizType, clientQuizMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for client_quiz")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clientQuizMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(clientQuizType, clientQuizMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for client_quiz")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for client_quiz")
	}

CacheNoHooks:
	if !cached {
		clientQuizUpsertCacheMut.Lock()
		clientQuizUpsertCache[key] = cache
		clientQuizUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClientQuiz record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClientQuiz) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClientQuiz provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientQuizPrimaryKeyMapping)
	sql := "DELETE FROM `client_quiz` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from client_quiz")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for client_quiz")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clientQuizQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clientQuizQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from client_quiz")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_quiz")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientQuizSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientQuizBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientQuizPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `client_quiz` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientQuizPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clientQuiz slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for client_quiz")
	}

	if len(clientQuizAfterDeleteHooks) != 0 {
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
func (o *ClientQuiz) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClientQuiz(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientQuizSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientQuizSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientQuizPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `client_quiz`.* FROM `client_quiz` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientQuizPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClientQuizSlice")
	}

	*o = slice

	return nil
}

// ClientQuizExists checks if the ClientQuiz row exists.
func ClientQuizExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `client_quiz` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if client_quiz exists")
	}

	return exists, nil
}

// Exists checks if the ClientQuiz row exists.
func (o *ClientQuiz) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ClientQuizExists(ctx, exec, o.ID)
}