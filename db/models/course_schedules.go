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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CourseSchedule is an object representing the database table.
type CourseSchedule struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Module   string      `boil:"module" json:"module" toml:"module" yaml:"module"`
	Day      string      `boil:"day" json:"day" toml:"day" yaml:"day"`
	Period   int16       `boil:"period" json:"period" toml:"period" yaml:"period"`
	Room     string      `boil:"room" json:"room" toml:"room" yaml:"room"`
	CourseID null.String `boil:"course_id" json:"course_id,omitempty" toml:"course_id" yaml:"course_id,omitempty"`

	R *courseScheduleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L courseScheduleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CourseScheduleColumns = struct {
	ID       string
	Module   string
	Day      string
	Period   string
	Room     string
	CourseID string
}{
	ID:       "id",
	Module:   "module",
	Day:      "day",
	Period:   "period",
	Room:     "room",
	CourseID: "course_id",
}

var CourseScheduleTableColumns = struct {
	ID       string
	Module   string
	Day      string
	Period   string
	Room     string
	CourseID string
}{
	ID:       "course_schedules.id",
	Module:   "course_schedules.module",
	Day:      "course_schedules.day",
	Period:   "course_schedules.period",
	Room:     "course_schedules.room",
	CourseID: "course_schedules.course_id",
}

// Generated where

var CourseScheduleWhere = struct {
	ID       whereHelperint
	Module   whereHelperstring
	Day      whereHelperstring
	Period   whereHelperint16
	Room     whereHelperstring
	CourseID whereHelpernull_String
}{
	ID:       whereHelperint{field: "\"course_schedules\".\"id\""},
	Module:   whereHelperstring{field: "\"course_schedules\".\"module\""},
	Day:      whereHelperstring{field: "\"course_schedules\".\"day\""},
	Period:   whereHelperint16{field: "\"course_schedules\".\"period\""},
	Room:     whereHelperstring{field: "\"course_schedules\".\"room\""},
	CourseID: whereHelpernull_String{field: "\"course_schedules\".\"course_id\""},
}

// CourseScheduleRels is where relationship names are stored.
var CourseScheduleRels = struct {
	Course string
}{
	Course: "Course",
}

// courseScheduleR is where relationships are stored.
type courseScheduleR struct {
	Course *Course `boil:"Course" json:"Course" toml:"Course" yaml:"Course"`
}

// NewStruct creates a new relationship struct
func (*courseScheduleR) NewStruct() *courseScheduleR {
	return &courseScheduleR{}
}

func (r *courseScheduleR) GetCourse() *Course {
	if r == nil {
		return nil
	}
	return r.Course
}

// courseScheduleL is where Load methods for each relationship are stored.
type courseScheduleL struct{}

var (
	courseScheduleAllColumns            = []string{"id", "module", "day", "period", "room", "course_id"}
	courseScheduleColumnsWithoutDefault = []string{"module", "day", "period", "room"}
	courseScheduleColumnsWithDefault    = []string{"id", "course_id"}
	courseSchedulePrimaryKeyColumns     = []string{"id"}
	courseScheduleGeneratedColumns      = []string{}
)

type (
	// CourseScheduleSlice is an alias for a slice of pointers to CourseSchedule.
	// This should almost always be used instead of []CourseSchedule.
	CourseScheduleSlice []*CourseSchedule
	// CourseScheduleHook is the signature for custom CourseSchedule hook methods
	CourseScheduleHook func(context.Context, boil.ContextExecutor, *CourseSchedule) error

	courseScheduleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	courseScheduleType                 = reflect.TypeOf(&CourseSchedule{})
	courseScheduleMapping              = queries.MakeStructMapping(courseScheduleType)
	courseSchedulePrimaryKeyMapping, _ = queries.BindMapping(courseScheduleType, courseScheduleMapping, courseSchedulePrimaryKeyColumns)
	courseScheduleInsertCacheMut       sync.RWMutex
	courseScheduleInsertCache          = make(map[string]insertCache)
	courseScheduleUpdateCacheMut       sync.RWMutex
	courseScheduleUpdateCache          = make(map[string]updateCache)
	courseScheduleUpsertCacheMut       sync.RWMutex
	courseScheduleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var courseScheduleAfterSelectHooks []CourseScheduleHook

var courseScheduleBeforeInsertHooks []CourseScheduleHook
var courseScheduleAfterInsertHooks []CourseScheduleHook

var courseScheduleBeforeUpdateHooks []CourseScheduleHook
var courseScheduleAfterUpdateHooks []CourseScheduleHook

var courseScheduleBeforeDeleteHooks []CourseScheduleHook
var courseScheduleAfterDeleteHooks []CourseScheduleHook

var courseScheduleBeforeUpsertHooks []CourseScheduleHook
var courseScheduleAfterUpsertHooks []CourseScheduleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CourseSchedule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CourseSchedule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CourseSchedule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CourseSchedule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CourseSchedule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CourseSchedule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CourseSchedule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CourseSchedule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CourseSchedule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range courseScheduleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCourseScheduleHook registers your hook function for all future operations.
func AddCourseScheduleHook(hookPoint boil.HookPoint, courseScheduleHook CourseScheduleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		courseScheduleAfterSelectHooks = append(courseScheduleAfterSelectHooks, courseScheduleHook)
	case boil.BeforeInsertHook:
		courseScheduleBeforeInsertHooks = append(courseScheduleBeforeInsertHooks, courseScheduleHook)
	case boil.AfterInsertHook:
		courseScheduleAfterInsertHooks = append(courseScheduleAfterInsertHooks, courseScheduleHook)
	case boil.BeforeUpdateHook:
		courseScheduleBeforeUpdateHooks = append(courseScheduleBeforeUpdateHooks, courseScheduleHook)
	case boil.AfterUpdateHook:
		courseScheduleAfterUpdateHooks = append(courseScheduleAfterUpdateHooks, courseScheduleHook)
	case boil.BeforeDeleteHook:
		courseScheduleBeforeDeleteHooks = append(courseScheduleBeforeDeleteHooks, courseScheduleHook)
	case boil.AfterDeleteHook:
		courseScheduleAfterDeleteHooks = append(courseScheduleAfterDeleteHooks, courseScheduleHook)
	case boil.BeforeUpsertHook:
		courseScheduleBeforeUpsertHooks = append(courseScheduleBeforeUpsertHooks, courseScheduleHook)
	case boil.AfterUpsertHook:
		courseScheduleAfterUpsertHooks = append(courseScheduleAfterUpsertHooks, courseScheduleHook)
	}
}

// One returns a single courseSchedule record from the query.
func (q courseScheduleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CourseSchedule, error) {
	o := &CourseSchedule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for course_schedules")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CourseSchedule records from the query.
func (q courseScheduleQuery) All(ctx context.Context, exec boil.ContextExecutor) (CourseScheduleSlice, error) {
	var o []*CourseSchedule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to CourseSchedule slice")
	}

	if len(courseScheduleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CourseSchedule records in the query.
func (q courseScheduleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count course_schedules rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q courseScheduleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if course_schedules exists")
	}

	return count > 0, nil
}

// Course pointed to by the foreign key.
func (o *CourseSchedule) Course(mods ...qm.QueryMod) courseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CourseID),
	}

	queryMods = append(queryMods, mods...)

	return Courses(queryMods...)
}

// LoadCourse allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (courseScheduleL) LoadCourse(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCourseSchedule interface{}, mods queries.Applicator) error {
	var slice []*CourseSchedule
	var object *CourseSchedule

	if singular {
		var ok bool
		object, ok = maybeCourseSchedule.(*CourseSchedule)
		if !ok {
			object = new(CourseSchedule)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCourseSchedule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCourseSchedule))
			}
		}
	} else {
		s, ok := maybeCourseSchedule.(*[]*CourseSchedule)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCourseSchedule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCourseSchedule))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &courseScheduleR{}
		}
		if !queries.IsNil(object.CourseID) {
			args = append(args, object.CourseID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &courseScheduleR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CourseID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.CourseID) {
				args = append(args, obj.CourseID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`courses`),
		qm.WhereIn(`courses.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Course")
	}

	var resultSlice []*Course
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Course")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for courses")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for courses")
	}

	if len(courseAfterSelectHooks) != 0 {
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
		object.R.Course = foreign
		if foreign.R == nil {
			foreign.R = &courseR{}
		}
		foreign.R.CourseSchedules = append(foreign.R.CourseSchedules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.CourseID, foreign.ID) {
				local.R.Course = foreign
				if foreign.R == nil {
					foreign.R = &courseR{}
				}
				foreign.R.CourseSchedules = append(foreign.R.CourseSchedules, local)
				break
			}
		}
	}

	return nil
}

// SetCourse of the courseSchedule to the related item.
// Sets o.R.Course to related.
// Adds o to related.R.CourseSchedules.
func (o *CourseSchedule) SetCourse(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Course) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"course_schedules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
		strmangle.WhereClause("\"", "\"", 2, courseSchedulePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.CourseID, related.ID)
	if o.R == nil {
		o.R = &courseScheduleR{
			Course: related,
		}
	} else {
		o.R.Course = related
	}

	if related.R == nil {
		related.R = &courseR{
			CourseSchedules: CourseScheduleSlice{o},
		}
	} else {
		related.R.CourseSchedules = append(related.R.CourseSchedules, o)
	}

	return nil
}

// RemoveCourse relationship.
// Sets o.R.Course to nil.
// Removes o from all passed in related items' relationships struct.
func (o *CourseSchedule) RemoveCourse(ctx context.Context, exec boil.ContextExecutor, related *Course) error {
	var err error

	queries.SetScanner(&o.CourseID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("course_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Course = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.CourseSchedules {
		if queries.Equal(o.CourseID, ri.CourseID) {
			continue
		}

		ln := len(related.R.CourseSchedules)
		if ln > 1 && i < ln-1 {
			related.R.CourseSchedules[i] = related.R.CourseSchedules[ln-1]
		}
		related.R.CourseSchedules = related.R.CourseSchedules[:ln-1]
		break
	}
	return nil
}

// CourseSchedules retrieves all the records using an executor.
func CourseSchedules(mods ...qm.QueryMod) courseScheduleQuery {
	mods = append(mods, qm.From("\"course_schedules\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"course_schedules\".*"})
	}

	return courseScheduleQuery{q}
}

// FindCourseSchedule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCourseSchedule(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CourseSchedule, error) {
	courseScheduleObj := &CourseSchedule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"course_schedules\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, courseScheduleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from course_schedules")
	}

	if err = courseScheduleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return courseScheduleObj, err
	}

	return courseScheduleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CourseSchedule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no course_schedules provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(courseScheduleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	courseScheduleInsertCacheMut.RLock()
	cache, cached := courseScheduleInsertCache[key]
	courseScheduleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			courseScheduleAllColumns,
			courseScheduleColumnsWithDefault,
			courseScheduleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(courseScheduleType, courseScheduleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(courseScheduleType, courseScheduleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"course_schedules\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"course_schedules\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodel: unable to insert into course_schedules")
	}

	if !cached {
		courseScheduleInsertCacheMut.Lock()
		courseScheduleInsertCache[key] = cache
		courseScheduleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CourseSchedule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CourseSchedule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	courseScheduleUpdateCacheMut.RLock()
	cache, cached := courseScheduleUpdateCache[key]
	courseScheduleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			courseScheduleAllColumns,
			courseSchedulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update course_schedules, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"course_schedules\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, courseSchedulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(courseScheduleType, courseScheduleMapping, append(wl, courseSchedulePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodel: unable to update course_schedules row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for course_schedules")
	}

	if !cached {
		courseScheduleUpdateCacheMut.Lock()
		courseScheduleUpdateCache[key] = cache
		courseScheduleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q courseScheduleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for course_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for course_schedules")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CourseScheduleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), courseSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"course_schedules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, courseSchedulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in courseSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all courseSchedule")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CourseSchedule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no course_schedules provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(courseScheduleColumnsWithDefault, o)

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

	courseScheduleUpsertCacheMut.RLock()
	cache, cached := courseScheduleUpsertCache[key]
	courseScheduleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			courseScheduleAllColumns,
			courseScheduleColumnsWithDefault,
			courseScheduleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			courseScheduleAllColumns,
			courseSchedulePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodel: unable to upsert course_schedules, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(courseSchedulePrimaryKeyColumns))
			copy(conflict, courseSchedulePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"course_schedules\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(courseScheduleType, courseScheduleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(courseScheduleType, courseScheduleMapping, ret)
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
		return errors.Wrap(err, "dbmodel: unable to upsert course_schedules")
	}

	if !cached {
		courseScheduleUpsertCacheMut.Lock()
		courseScheduleUpsertCache[key] = cache
		courseScheduleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CourseSchedule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CourseSchedule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no CourseSchedule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), courseSchedulePrimaryKeyMapping)
	sql := "DELETE FROM \"course_schedules\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from course_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for course_schedules")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q courseScheduleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no courseScheduleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from course_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for course_schedules")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CourseScheduleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(courseScheduleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), courseSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"course_schedules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, courseSchedulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from courseSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for course_schedules")
	}

	if len(courseScheduleAfterDeleteHooks) != 0 {
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
func (o *CourseSchedule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCourseSchedule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CourseScheduleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CourseScheduleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), courseSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"course_schedules\".* FROM \"course_schedules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, courseSchedulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in CourseScheduleSlice")
	}

	*o = slice

	return nil
}

// CourseScheduleExists checks if the CourseSchedule row exists.
func CourseScheduleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"course_schedules\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if course_schedules exists")
	}

	return exists, nil
}

// Exists checks if the CourseSchedule row exists.
func (o *CourseSchedule) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CourseScheduleExists(ctx, exec, o.ID)
}
