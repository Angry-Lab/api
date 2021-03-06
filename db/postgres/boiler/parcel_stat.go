// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

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

// ParcelStat is an object representing the database table.
type ParcelStat struct {
	Hid       string  `boil:"hid" json:"hid" toml:"hid" yaml:"hid"`
	CNT       int     `boil:"cnt" json:"cnt" toml:"cnt" yaml:"cnt"`
	TotalCost float32 `boil:"total_cost" json:"total_cost" toml:"total_cost" yaml:"total_cost"`
	TotalNP   float32 `boil:"total_np" json:"total_np" toml:"total_np" yaml:"total_np"`

	R *parcelStatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L parcelStatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ParcelStatColumns = struct {
	Hid       string
	CNT       string
	TotalCost string
	TotalNP   string
}{
	Hid:       "hid",
	CNT:       "cnt",
	TotalCost: "total_cost",
	TotalNP:   "total_np",
}

// Generated where

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ParcelStatWhere = struct {
	Hid       whereHelperstring
	CNT       whereHelperint
	TotalCost whereHelperfloat32
	TotalNP   whereHelperfloat32
}{
	Hid:       whereHelperstring{field: "\"parcel_stat\".\"hid\""},
	CNT:       whereHelperint{field: "\"parcel_stat\".\"cnt\""},
	TotalCost: whereHelperfloat32{field: "\"parcel_stat\".\"total_cost\""},
	TotalNP:   whereHelperfloat32{field: "\"parcel_stat\".\"total_np\""},
}

// ParcelStatRels is where relationship names are stored.
var ParcelStatRels = struct {
}{}

// parcelStatR is where relationships are stored.
type parcelStatR struct {
}

// NewStruct creates a new relationship struct
func (*parcelStatR) NewStruct() *parcelStatR {
	return &parcelStatR{}
}

// parcelStatL is where Load methods for each relationship are stored.
type parcelStatL struct{}

var (
	parcelStatAllColumns            = []string{"hid", "cnt", "total_cost", "total_np"}
	parcelStatColumnsWithoutDefault = []string{"hid"}
	parcelStatColumnsWithDefault    = []string{"cnt", "total_cost", "total_np"}
	parcelStatPrimaryKeyColumns     = []string{"hid"}
)

type (
	// ParcelStatSlice is an alias for a slice of pointers to ParcelStat.
	// This should generally be used opposed to []ParcelStat.
	ParcelStatSlice []*ParcelStat
	// ParcelStatHook is the signature for custom ParcelStat hook methods
	ParcelStatHook func(context.Context, boil.ContextExecutor, *ParcelStat) error

	parcelStatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	parcelStatType                 = reflect.TypeOf(&ParcelStat{})
	parcelStatMapping              = queries.MakeStructMapping(parcelStatType)
	parcelStatPrimaryKeyMapping, _ = queries.BindMapping(parcelStatType, parcelStatMapping, parcelStatPrimaryKeyColumns)
	parcelStatInsertCacheMut       sync.RWMutex
	parcelStatInsertCache          = make(map[string]insertCache)
	parcelStatUpdateCacheMut       sync.RWMutex
	parcelStatUpdateCache          = make(map[string]updateCache)
	parcelStatUpsertCacheMut       sync.RWMutex
	parcelStatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var parcelStatBeforeInsertHooks []ParcelStatHook
var parcelStatBeforeUpdateHooks []ParcelStatHook
var parcelStatBeforeDeleteHooks []ParcelStatHook
var parcelStatBeforeUpsertHooks []ParcelStatHook

var parcelStatAfterInsertHooks []ParcelStatHook
var parcelStatAfterSelectHooks []ParcelStatHook
var parcelStatAfterUpdateHooks []ParcelStatHook
var parcelStatAfterDeleteHooks []ParcelStatHook
var parcelStatAfterUpsertHooks []ParcelStatHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ParcelStat) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ParcelStat) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ParcelStat) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ParcelStat) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ParcelStat) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ParcelStat) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ParcelStat) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ParcelStat) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ParcelStat) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range parcelStatAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddParcelStatHook registers your hook function for all future operations.
func AddParcelStatHook(hookPoint boil.HookPoint, parcelStatHook ParcelStatHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		parcelStatBeforeInsertHooks = append(parcelStatBeforeInsertHooks, parcelStatHook)
	case boil.BeforeUpdateHook:
		parcelStatBeforeUpdateHooks = append(parcelStatBeforeUpdateHooks, parcelStatHook)
	case boil.BeforeDeleteHook:
		parcelStatBeforeDeleteHooks = append(parcelStatBeforeDeleteHooks, parcelStatHook)
	case boil.BeforeUpsertHook:
		parcelStatBeforeUpsertHooks = append(parcelStatBeforeUpsertHooks, parcelStatHook)
	case boil.AfterInsertHook:
		parcelStatAfterInsertHooks = append(parcelStatAfterInsertHooks, parcelStatHook)
	case boil.AfterSelectHook:
		parcelStatAfterSelectHooks = append(parcelStatAfterSelectHooks, parcelStatHook)
	case boil.AfterUpdateHook:
		parcelStatAfterUpdateHooks = append(parcelStatAfterUpdateHooks, parcelStatHook)
	case boil.AfterDeleteHook:
		parcelStatAfterDeleteHooks = append(parcelStatAfterDeleteHooks, parcelStatHook)
	case boil.AfterUpsertHook:
		parcelStatAfterUpsertHooks = append(parcelStatAfterUpsertHooks, parcelStatHook)
	}
}

// One returns a single parcelStat record from the query.
func (q parcelStatQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ParcelStat, error) {
	o := &ParcelStat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for parcel_stat")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ParcelStat records from the query.
func (q parcelStatQuery) All(ctx context.Context, exec boil.ContextExecutor) (ParcelStatSlice, error) {
	var o []*ParcelStat

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to ParcelStat slice")
	}

	if len(parcelStatAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ParcelStat records in the query.
func (q parcelStatQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count parcel_stat rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q parcelStatQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if parcel_stat exists")
	}

	return count > 0, nil
}

// ParcelStats retrieves all the records using an executor.
func ParcelStats(mods ...qm.QueryMod) parcelStatQuery {
	mods = append(mods, qm.From("\"parcel_stat\""))
	return parcelStatQuery{NewQuery(mods...)}
}

// FindParcelStat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindParcelStat(ctx context.Context, exec boil.ContextExecutor, hid string, selectCols ...string) (*ParcelStat, error) {
	parcelStatObj := &ParcelStat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"parcel_stat\" where \"hid\"=$1", sel,
	)

	q := queries.Raw(query, hid)

	err := q.Bind(ctx, exec, parcelStatObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from parcel_stat")
	}

	return parcelStatObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ParcelStat) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no parcel_stat provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(parcelStatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	parcelStatInsertCacheMut.RLock()
	cache, cached := parcelStatInsertCache[key]
	parcelStatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			parcelStatAllColumns,
			parcelStatColumnsWithDefault,
			parcelStatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(parcelStatType, parcelStatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(parcelStatType, parcelStatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"parcel_stat\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"parcel_stat\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into parcel_stat")
	}

	if !cached {
		parcelStatInsertCacheMut.Lock()
		parcelStatInsertCache[key] = cache
		parcelStatInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ParcelStat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ParcelStat) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	parcelStatUpdateCacheMut.RLock()
	cache, cached := parcelStatUpdateCache[key]
	parcelStatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			parcelStatAllColumns,
			parcelStatPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update parcel_stat, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"parcel_stat\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, parcelStatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(parcelStatType, parcelStatMapping, append(wl, parcelStatPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update parcel_stat row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for parcel_stat")
	}

	if !cached {
		parcelStatUpdateCacheMut.Lock()
		parcelStatUpdateCache[key] = cache
		parcelStatUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q parcelStatQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for parcel_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for parcel_stat")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ParcelStatSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parcelStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"parcel_stat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, parcelStatPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in parcelStat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all parcelStat")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ParcelStat) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no parcel_stat provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(parcelStatColumnsWithDefault, o)

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

	parcelStatUpsertCacheMut.RLock()
	cache, cached := parcelStatUpsertCache[key]
	parcelStatUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			parcelStatAllColumns,
			parcelStatColumnsWithDefault,
			parcelStatColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			parcelStatAllColumns,
			parcelStatPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert parcel_stat, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(parcelStatPrimaryKeyColumns))
			copy(conflict, parcelStatPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"parcel_stat\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(parcelStatType, parcelStatMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(parcelStatType, parcelStatMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert parcel_stat")
	}

	if !cached {
		parcelStatUpsertCacheMut.Lock()
		parcelStatUpsertCache[key] = cache
		parcelStatUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ParcelStat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ParcelStat) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no ParcelStat provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), parcelStatPrimaryKeyMapping)
	sql := "DELETE FROM \"parcel_stat\" WHERE \"hid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from parcel_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for parcel_stat")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q parcelStatQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no parcelStatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from parcel_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for parcel_stat")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ParcelStatSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(parcelStatBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parcelStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"parcel_stat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, parcelStatPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from parcelStat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for parcel_stat")
	}

	if len(parcelStatAfterDeleteHooks) != 0 {
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
func (o *ParcelStat) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindParcelStat(ctx, exec, o.Hid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ParcelStatSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ParcelStatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parcelStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"parcel_stat\".* FROM \"parcel_stat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, parcelStatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in ParcelStatSlice")
	}

	*o = slice

	return nil
}

// ParcelStatExists checks if the ParcelStat row exists.
func ParcelStatExists(ctx context.Context, exec boil.ContextExecutor, hid string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"parcel_stat\" where \"hid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, hid)
	}
	row := exec.QueryRowContext(ctx, sql, hid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if parcel_stat exists")
	}

	return exists, nil
}
