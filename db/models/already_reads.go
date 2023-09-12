// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodel

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

// AlreadyRead is an object representing the database table.
type AlreadyRead struct {
	ID            string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	InformationID string    `boil:"information_id" json:"information_id" toml:"information_id" yaml:"information_id"`
	ReadUser      string    `boil:"read_user" json:"read_user" toml:"read_user" yaml:"read_user"`
	ReadAt        time.Time `boil:"read_at" json:"read_at" toml:"read_at" yaml:"read_at"`

	R *alreadyReadR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L alreadyReadL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AlreadyReadColumns = struct {
	ID            string
	InformationID string
	ReadUser      string
	ReadAt        string
}{
	ID:            "id",
	InformationID: "information_id",
	ReadUser:      "read_user",
	ReadAt:        "read_at",
}

var AlreadyReadTableColumns = struct {
	ID            string
	InformationID string
	ReadUser      string
	ReadAt        string
}{
	ID:            "already_reads.id",
	InformationID: "already_reads.information_id",
	ReadUser:      "already_reads.read_user",
	ReadAt:        "already_reads.read_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AlreadyReadWhere = struct {
	ID            whereHelperstring
	InformationID whereHelperstring
	ReadUser      whereHelperstring
	ReadAt        whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "\"already_reads\".\"id\""},
	InformationID: whereHelperstring{field: "\"already_reads\".\"information_id\""},
	ReadUser:      whereHelperstring{field: "\"already_reads\".\"read_user\""},
	ReadAt:        whereHelpertime_Time{field: "\"already_reads\".\"read_at\""},
}

// AlreadyReadRels is where relationship names are stored.
var AlreadyReadRels = struct {
}{}

// alreadyReadR is where relationships are stored.
type alreadyReadR struct {
}

// NewStruct creates a new relationship struct
func (*alreadyReadR) NewStruct() *alreadyReadR {
	return &alreadyReadR{}
}

// alreadyReadL is where Load methods for each relationship are stored.
type alreadyReadL struct{}

var (
	alreadyReadAllColumns            = []string{"id", "information_id", "read_user", "read_at"}
	alreadyReadColumnsWithoutDefault = []string{"information_id", "read_user", "read_at"}
	alreadyReadColumnsWithDefault    = []string{"id"}
	alreadyReadPrimaryKeyColumns     = []string{"id"}
	alreadyReadGeneratedColumns      = []string{}
)

type (
	// AlreadyReadSlice is an alias for a slice of pointers to AlreadyRead.
	// This should almost always be used instead of []AlreadyRead.
	AlreadyReadSlice []*AlreadyRead
	// AlreadyReadHook is the signature for custom AlreadyRead hook methods
	AlreadyReadHook func(context.Context, boil.ContextExecutor, *AlreadyRead) error

	alreadyReadQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	alreadyReadType                 = reflect.TypeOf(&AlreadyRead{})
	alreadyReadMapping              = queries.MakeStructMapping(alreadyReadType)
	alreadyReadPrimaryKeyMapping, _ = queries.BindMapping(alreadyReadType, alreadyReadMapping, alreadyReadPrimaryKeyColumns)
	alreadyReadInsertCacheMut       sync.RWMutex
	alreadyReadInsertCache          = make(map[string]insertCache)
	alreadyReadUpdateCacheMut       sync.RWMutex
	alreadyReadUpdateCache          = make(map[string]updateCache)
	alreadyReadUpsertCacheMut       sync.RWMutex
	alreadyReadUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var alreadyReadAfterSelectHooks []AlreadyReadHook

var alreadyReadBeforeInsertHooks []AlreadyReadHook
var alreadyReadAfterInsertHooks []AlreadyReadHook

var alreadyReadBeforeUpdateHooks []AlreadyReadHook
var alreadyReadAfterUpdateHooks []AlreadyReadHook

var alreadyReadBeforeDeleteHooks []AlreadyReadHook
var alreadyReadAfterDeleteHooks []AlreadyReadHook

var alreadyReadBeforeUpsertHooks []AlreadyReadHook
var alreadyReadAfterUpsertHooks []AlreadyReadHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AlreadyRead) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AlreadyRead) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AlreadyRead) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AlreadyRead) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AlreadyRead) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AlreadyRead) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AlreadyRead) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AlreadyRead) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AlreadyRead) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range alreadyReadAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAlreadyReadHook registers your hook function for all future operations.
func AddAlreadyReadHook(hookPoint boil.HookPoint, alreadyReadHook AlreadyReadHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		alreadyReadAfterSelectHooks = append(alreadyReadAfterSelectHooks, alreadyReadHook)
	case boil.BeforeInsertHook:
		alreadyReadBeforeInsertHooks = append(alreadyReadBeforeInsertHooks, alreadyReadHook)
	case boil.AfterInsertHook:
		alreadyReadAfterInsertHooks = append(alreadyReadAfterInsertHooks, alreadyReadHook)
	case boil.BeforeUpdateHook:
		alreadyReadBeforeUpdateHooks = append(alreadyReadBeforeUpdateHooks, alreadyReadHook)
	case boil.AfterUpdateHook:
		alreadyReadAfterUpdateHooks = append(alreadyReadAfterUpdateHooks, alreadyReadHook)
	case boil.BeforeDeleteHook:
		alreadyReadBeforeDeleteHooks = append(alreadyReadBeforeDeleteHooks, alreadyReadHook)
	case boil.AfterDeleteHook:
		alreadyReadAfterDeleteHooks = append(alreadyReadAfterDeleteHooks, alreadyReadHook)
	case boil.BeforeUpsertHook:
		alreadyReadBeforeUpsertHooks = append(alreadyReadBeforeUpsertHooks, alreadyReadHook)
	case boil.AfterUpsertHook:
		alreadyReadAfterUpsertHooks = append(alreadyReadAfterUpsertHooks, alreadyReadHook)
	}
}

// One returns a single alreadyRead record from the query.
func (q alreadyReadQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AlreadyRead, error) {
	o := &AlreadyRead{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for already_reads")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AlreadyRead records from the query.
func (q alreadyReadQuery) All(ctx context.Context, exec boil.ContextExecutor) (AlreadyReadSlice, error) {
	var o []*AlreadyRead

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to AlreadyRead slice")
	}

	if len(alreadyReadAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AlreadyRead records in the query.
func (q alreadyReadQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count already_reads rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q alreadyReadQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if already_reads exists")
	}

	return count > 0, nil
}

// AlreadyReads retrieves all the records using an executor.
func AlreadyReads(mods ...qm.QueryMod) alreadyReadQuery {
	mods = append(mods, qm.From("\"already_reads\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"already_reads\".*"})
	}

	return alreadyReadQuery{q}
}

// FindAlreadyRead retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAlreadyRead(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AlreadyRead, error) {
	alreadyReadObj := &AlreadyRead{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"already_reads\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, alreadyReadObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from already_reads")
	}

	if err = alreadyReadObj.doAfterSelectHooks(ctx, exec); err != nil {
		return alreadyReadObj, err
	}

	return alreadyReadObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AlreadyRead) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no already_reads provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(alreadyReadColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	alreadyReadInsertCacheMut.RLock()
	cache, cached := alreadyReadInsertCache[key]
	alreadyReadInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			alreadyReadAllColumns,
			alreadyReadColumnsWithDefault,
			alreadyReadColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(alreadyReadType, alreadyReadMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(alreadyReadType, alreadyReadMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"already_reads\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"already_reads\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to insert into already_reads")
	}

	if !cached {
		alreadyReadInsertCacheMut.Lock()
		alreadyReadInsertCache[key] = cache
		alreadyReadInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AlreadyRead.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AlreadyRead) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	alreadyReadUpdateCacheMut.RLock()
	cache, cached := alreadyReadUpdateCache[key]
	alreadyReadUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			alreadyReadAllColumns,
			alreadyReadPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update already_reads, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"already_reads\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, alreadyReadPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(alreadyReadType, alreadyReadMapping, append(wl, alreadyReadPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodel: unable to update already_reads row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for already_reads")
	}

	if !cached {
		alreadyReadUpdateCacheMut.Lock()
		alreadyReadUpdateCache[key] = cache
		alreadyReadUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q alreadyReadQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for already_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for already_reads")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AlreadyReadSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alreadyReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"already_reads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, alreadyReadPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in alreadyRead slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all alreadyRead")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AlreadyRead) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no already_reads provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(alreadyReadColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	alreadyReadUpsertCacheMut.RLock()
	cache, cached := alreadyReadUpsertCache[key]
	alreadyReadUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			alreadyReadAllColumns,
			alreadyReadColumnsWithDefault,
			alreadyReadColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			alreadyReadAllColumns,
			alreadyReadPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodel: unable to upsert already_reads, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(alreadyReadPrimaryKeyColumns))
			copy(conflict, alreadyReadPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"already_reads\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(alreadyReadType, alreadyReadMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(alreadyReadType, alreadyReadMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to upsert already_reads")
	}

	if !cached {
		alreadyReadUpsertCacheMut.Lock()
		alreadyReadUpsertCache[key] = cache
		alreadyReadUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AlreadyRead record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AlreadyRead) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no AlreadyRead provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), alreadyReadPrimaryKeyMapping)
	sql := "DELETE FROM \"already_reads\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from already_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for already_reads")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q alreadyReadQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no alreadyReadQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from already_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for already_reads")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AlreadyReadSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(alreadyReadBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alreadyReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"already_reads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, alreadyReadPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from alreadyRead slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for already_reads")
	}

	if len(alreadyReadAfterDeleteHooks) != 0 {
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
func (o *AlreadyRead) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAlreadyRead(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AlreadyReadSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AlreadyReadSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alreadyReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"already_reads\".* FROM \"already_reads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, alreadyReadPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in AlreadyReadSlice")
	}

	*o = slice

	return nil
}

// AlreadyReadExists checks if the AlreadyRead row exists.
func AlreadyReadExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"already_reads\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if already_reads exists")
	}

	return exists, nil
}

// Exists checks if the AlreadyRead row exists.
func (o *AlreadyRead) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AlreadyReadExists(ctx, exec, o.ID)
}
