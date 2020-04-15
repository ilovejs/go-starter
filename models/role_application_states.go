// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RoleApplicationState is an object representing the database table.
type RoleApplicationState struct {
	RoleID             string    `boil:"role_id" json:"role_id" toml:"role_id" yaml:"role_id"`
	ApplicationStateID string    `boil:"application_state_id" json:"application_state_id" toml:"application_state_id" yaml:"application_state_id"`
	CreatedAt          time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt          time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *roleApplicationStateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roleApplicationStateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoleApplicationStateColumns = struct {
	RoleID             string
	ApplicationStateID string
	CreatedAt          string
	UpdatedAt          string
}{
	RoleID:             "role_id",
	ApplicationStateID: "application_state_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// Generated where

var RoleApplicationStateWhere = struct {
	RoleID             whereHelperstring
	ApplicationStateID whereHelperstring
	CreatedAt          whereHelpertime_Time
	UpdatedAt          whereHelpertime_Time
}{
	RoleID:             whereHelperstring{field: "\"role_application_states\".\"role_id\""},
	ApplicationStateID: whereHelperstring{field: "\"role_application_states\".\"application_state_id\""},
	CreatedAt:          whereHelpertime_Time{field: "\"role_application_states\".\"created_at\""},
	UpdatedAt:          whereHelpertime_Time{field: "\"role_application_states\".\"updated_at\""},
}

// RoleApplicationStateRels is where relationship names are stored.
var RoleApplicationStateRels = struct {
	ApplicationState string
	Role             string
}{
	ApplicationState: "ApplicationState",
	Role:             "Role",
}

// roleApplicationStateR is where relationships are stored.
type roleApplicationStateR struct {
	ApplicationState *ApplicationState
	Role             *Role
}

// NewStruct creates a new relationship struct
func (*roleApplicationStateR) NewStruct() *roleApplicationStateR {
	return &roleApplicationStateR{}
}

// roleApplicationStateL is where Load methods for each relationship are stored.
type roleApplicationStateL struct{}

var (
	roleApplicationStateAllColumns            = []string{"role_id", "application_state_id", "created_at", "updated_at"}
	roleApplicationStateColumnsWithoutDefault = []string{"role_id", "application_state_id", "created_at", "updated_at"}
	roleApplicationStateColumnsWithDefault    = []string{}
	roleApplicationStatePrimaryKeyColumns     = []string{"role_id", "application_state_id"}
)

type (
	// RoleApplicationStateSlice is an alias for a slice of pointers to RoleApplicationState.
	// This should generally be used opposed to []RoleApplicationState.
	RoleApplicationStateSlice []*RoleApplicationState

	roleApplicationStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roleApplicationStateType                 = reflect.TypeOf(&RoleApplicationState{})
	roleApplicationStateMapping              = queries.MakeStructMapping(roleApplicationStateType)
	roleApplicationStatePrimaryKeyMapping, _ = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, roleApplicationStatePrimaryKeyColumns)
	roleApplicationStateInsertCacheMut       sync.RWMutex
	roleApplicationStateInsertCache          = make(map[string]insertCache)
	roleApplicationStateUpdateCacheMut       sync.RWMutex
	roleApplicationStateUpdateCache          = make(map[string]updateCache)
	roleApplicationStateUpsertCacheMut       sync.RWMutex
	roleApplicationStateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single roleApplicationState record from the query.
func (q roleApplicationStateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RoleApplicationState, error) {
	o := &RoleApplicationState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for role_application_states")
	}

	return o, nil
}

// All returns all RoleApplicationState records from the query.
func (q roleApplicationStateQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoleApplicationStateSlice, error) {
	var o []*RoleApplicationState

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RoleApplicationState slice")
	}

	return o, nil
}

// Count returns the count of all RoleApplicationState records in the query.
func (q roleApplicationStateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count role_application_states rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roleApplicationStateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if role_application_states exists")
	}

	return count > 0, nil
}

// ApplicationState pointed to by the foreign key.
func (o *RoleApplicationState) ApplicationState(mods ...qm.QueryMod) applicationStateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ApplicationStateID),
	}

	queryMods = append(queryMods, mods...)

	query := ApplicationStates(queryMods...)
	queries.SetFrom(query.Query, "\"application_states\"")

	return query
}

// Role pointed to by the foreign key.
func (o *RoleApplicationState) Role(mods ...qm.QueryMod) roleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RoleID),
	}

	queryMods = append(queryMods, mods...)

	query := Roles(queryMods...)
	queries.SetFrom(query.Query, "\"roles\"")

	return query
}

// LoadApplicationState allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (roleApplicationStateL) LoadApplicationState(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoleApplicationState interface{}, mods queries.Applicator) error {
	var slice []*RoleApplicationState
	var object *RoleApplicationState

	if singular {
		object = maybeRoleApplicationState.(*RoleApplicationState)
	} else {
		slice = *maybeRoleApplicationState.(*[]*RoleApplicationState)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleApplicationStateR{}
		}
		args = append(args, object.ApplicationStateID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleApplicationStateR{}
			}

			for _, a := range args {
				if a == obj.ApplicationStateID {
					continue Outer
				}
			}

			args = append(args, obj.ApplicationStateID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`application_states`), qm.WhereIn(`application_states.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ApplicationState")
	}

	var resultSlice []*ApplicationState
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ApplicationState")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for application_states")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for application_states")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ApplicationState = foreign
		if foreign.R == nil {
			foreign.R = &applicationStateR{}
		}
		foreign.R.RoleApplicationStates = append(foreign.R.RoleApplicationStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ApplicationStateID == foreign.ID {
				local.R.ApplicationState = foreign
				if foreign.R == nil {
					foreign.R = &applicationStateR{}
				}
				foreign.R.RoleApplicationStates = append(foreign.R.RoleApplicationStates, local)
				break
			}
		}
	}

	return nil
}

// LoadRole allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (roleApplicationStateL) LoadRole(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoleApplicationState interface{}, mods queries.Applicator) error {
	var slice []*RoleApplicationState
	var object *RoleApplicationState

	if singular {
		object = maybeRoleApplicationState.(*RoleApplicationState)
	} else {
		slice = *maybeRoleApplicationState.(*[]*RoleApplicationState)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleApplicationStateR{}
		}
		args = append(args, object.RoleID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleApplicationStateR{}
			}

			for _, a := range args {
				if a == obj.RoleID {
					continue Outer
				}
			}

			args = append(args, obj.RoleID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`roles`), qm.WhereIn(`roles.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Role")
	}

	var resultSlice []*Role
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Role")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for roles")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Role = foreign
		if foreign.R == nil {
			foreign.R = &roleR{}
		}
		foreign.R.RoleApplicationStates = append(foreign.R.RoleApplicationStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RoleID == foreign.ID {
				local.R.Role = foreign
				if foreign.R == nil {
					foreign.R = &roleR{}
				}
				foreign.R.RoleApplicationStates = append(foreign.R.RoleApplicationStates, local)
				break
			}
		}
	}

	return nil
}

// SetApplicationState of the roleApplicationState to the related item.
// Sets o.R.ApplicationState to related.
// Adds o to related.R.RoleApplicationStates.
func (o *RoleApplicationState) SetApplicationState(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ApplicationState) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"role_application_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"application_state_id"}),
		strmangle.WhereClause("\"", "\"", 2, roleApplicationStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.RoleID, o.ApplicationStateID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ApplicationStateID = related.ID
	if o.R == nil {
		o.R = &roleApplicationStateR{
			ApplicationState: related,
		}
	} else {
		o.R.ApplicationState = related
	}

	if related.R == nil {
		related.R = &applicationStateR{
			RoleApplicationStates: RoleApplicationStateSlice{o},
		}
	} else {
		related.R.RoleApplicationStates = append(related.R.RoleApplicationStates, o)
	}

	return nil
}

// SetRole of the roleApplicationState to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleApplicationStates.
func (o *RoleApplicationState) SetRole(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Role) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"role_application_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
		strmangle.WhereClause("\"", "\"", 2, roleApplicationStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.RoleID, o.ApplicationStateID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RoleID = related.ID
	if o.R == nil {
		o.R = &roleApplicationStateR{
			Role: related,
		}
	} else {
		o.R.Role = related
	}

	if related.R == nil {
		related.R = &roleR{
			RoleApplicationStates: RoleApplicationStateSlice{o},
		}
	} else {
		related.R.RoleApplicationStates = append(related.R.RoleApplicationStates, o)
	}

	return nil
}

// RoleApplicationStates retrieves all the records using an executor.
func RoleApplicationStates(mods ...qm.QueryMod) roleApplicationStateQuery {
	mods = append(mods, qm.From("\"role_application_states\""))
	return roleApplicationStateQuery{NewQuery(mods...)}
}

// FindRoleApplicationState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRoleApplicationState(ctx context.Context, exec boil.ContextExecutor, roleID string, applicationStateID string, selectCols ...string) (*RoleApplicationState, error) {
	roleApplicationStateObj := &RoleApplicationState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"role_application_states\" where \"role_id\"=$1 AND \"application_state_id\"=$2", sel,
	)

	q := queries.Raw(query, roleID, applicationStateID)

	err := q.Bind(ctx, exec, roleApplicationStateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from role_application_states")
	}

	return roleApplicationStateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RoleApplicationState) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role_application_states provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(roleApplicationStateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roleApplicationStateInsertCacheMut.RLock()
	cache, cached := roleApplicationStateInsertCache[key]
	roleApplicationStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roleApplicationStateAllColumns,
			roleApplicationStateColumnsWithDefault,
			roleApplicationStateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"role_application_states\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"role_application_states\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into role_application_states")
	}

	if !cached {
		roleApplicationStateInsertCacheMut.Lock()
		roleApplicationStateInsertCache[key] = cache
		roleApplicationStateInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the RoleApplicationState.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RoleApplicationState) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	roleApplicationStateUpdateCacheMut.RLock()
	cache, cached := roleApplicationStateUpdateCache[key]
	roleApplicationStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roleApplicationStateAllColumns,
			roleApplicationStatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update role_application_states, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"role_application_states\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, roleApplicationStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, append(wl, roleApplicationStatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update role_application_states row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for role_application_states")
	}

	if !cached {
		roleApplicationStateUpdateCacheMut.Lock()
		roleApplicationStateUpdateCache[key] = cache
		roleApplicationStateUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q roleApplicationStateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for role_application_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for role_application_states")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoleApplicationStateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleApplicationStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"role_application_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, roleApplicationStatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in roleApplicationState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all roleApplicationState")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RoleApplicationState) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role_application_states provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(roleApplicationStateColumnsWithDefault, o)

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

	roleApplicationStateUpsertCacheMut.RLock()
	cache, cached := roleApplicationStateUpsertCache[key]
	roleApplicationStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roleApplicationStateAllColumns,
			roleApplicationStateColumnsWithDefault,
			roleApplicationStateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			roleApplicationStateAllColumns,
			roleApplicationStatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert role_application_states, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(roleApplicationStatePrimaryKeyColumns))
			copy(conflict, roleApplicationStatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"role_application_states\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roleApplicationStateType, roleApplicationStateMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert role_application_states")
	}

	if !cached {
		roleApplicationStateUpsertCacheMut.Lock()
		roleApplicationStateUpsertCache[key] = cache
		roleApplicationStateUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single RoleApplicationState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RoleApplicationState) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RoleApplicationState provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), roleApplicationStatePrimaryKeyMapping)
	sql := "DELETE FROM \"role_application_states\" WHERE \"role_id\"=$1 AND \"application_state_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from role_application_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for role_application_states")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roleApplicationStateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no roleApplicationStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from role_application_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role_application_states")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoleApplicationStateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleApplicationStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"role_application_states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, roleApplicationStatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from roleApplicationState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role_application_states")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RoleApplicationState) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRoleApplicationState(ctx, exec, o.RoleID, o.ApplicationStateID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoleApplicationStateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoleApplicationStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleApplicationStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"role_application_states\".* FROM \"role_application_states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, roleApplicationStatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RoleApplicationStateSlice")
	}

	*o = slice

	return nil
}

// RoleApplicationStateExists checks if the RoleApplicationState row exists.
func RoleApplicationStateExists(ctx context.Context, exec boil.ContextExecutor, roleID string, applicationStateID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"role_application_states\" where \"role_id\"=$1 AND \"application_state_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, roleID, applicationStateID)
	}
	row := exec.QueryRowContext(ctx, sql, roleID, applicationStateID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if role_application_states exists")
	}

	return exists, nil
}
