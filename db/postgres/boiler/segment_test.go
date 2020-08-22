// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSegments(t *testing.T) {
	t.Parallel()

	query := Segments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSegmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSegmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Segments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSegmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SegmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSegmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SegmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Segment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SegmentExists to return true, but got false.")
	}
}

func testSegmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	segmentFound, err := FindSegment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if segmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSegmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Segments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSegmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Segments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSegmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	segmentOne := &Segment{}
	segmentTwo := &Segment{}
	if err = randomize.Struct(seed, segmentOne, segmentDBTypes, false, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}
	if err = randomize.Struct(seed, segmentTwo, segmentDBTypes, false, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = segmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = segmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Segments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSegmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	segmentOne := &Segment{}
	segmentTwo := &Segment{}
	if err = randomize.Struct(seed, segmentOne, segmentDBTypes, false, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}
	if err = randomize.Struct(seed, segmentTwo, segmentDBTypes, false, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = segmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = segmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func segmentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func segmentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Segment) error {
	*o = Segment{}
	return nil
}

func testSegmentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Segment{}
	o := &Segment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, segmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Segment object: %s", err)
	}

	AddSegmentHook(boil.BeforeInsertHook, segmentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	segmentBeforeInsertHooks = []SegmentHook{}

	AddSegmentHook(boil.AfterInsertHook, segmentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	segmentAfterInsertHooks = []SegmentHook{}

	AddSegmentHook(boil.AfterSelectHook, segmentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	segmentAfterSelectHooks = []SegmentHook{}

	AddSegmentHook(boil.BeforeUpdateHook, segmentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	segmentBeforeUpdateHooks = []SegmentHook{}

	AddSegmentHook(boil.AfterUpdateHook, segmentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	segmentAfterUpdateHooks = []SegmentHook{}

	AddSegmentHook(boil.BeforeDeleteHook, segmentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	segmentBeforeDeleteHooks = []SegmentHook{}

	AddSegmentHook(boil.AfterDeleteHook, segmentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	segmentAfterDeleteHooks = []SegmentHook{}

	AddSegmentHook(boil.BeforeUpsertHook, segmentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	segmentBeforeUpsertHooks = []SegmentHook{}

	AddSegmentHook(boil.AfterUpsertHook, segmentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	segmentAfterUpsertHooks = []SegmentHook{}
}

func testSegmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSegmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(segmentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSegmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSegmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SegmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSegmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Segments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	segmentDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Description`: `character varying`, `Condition`: `character varying`}
	_              = bytes.MinRead
)

func testSegmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(segmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(segmentAllColumns) == len(segmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSegmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(segmentAllColumns) == len(segmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Segment{}
	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, segmentDBTypes, true, segmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(segmentAllColumns, segmentPrimaryKeyColumns) {
		fields = segmentAllColumns
	} else {
		fields = strmangle.SetComplement(
			segmentAllColumns,
			segmentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SegmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSegmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(segmentAllColumns) == len(segmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Segment{}
	if err = randomize.Struct(seed, &o, segmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Segment: %s", err)
	}

	count, err := Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, segmentDBTypes, false, segmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Segment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Segment: %s", err)
	}

	count, err = Segments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}