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

func testAuthTokens(t *testing.T) {
	t.Parallel()

	query := AuthTokens()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
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

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AuthTokens().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthTokenSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthTokenExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AuthToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthTokenExists to return true, but got false.")
	}
}

func testAuthTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authTokenFound, err := FindAuthToken(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authTokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AuthTokens().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AuthTokens().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authTokenOne := &AuthToken{}
	authTokenTwo := &AuthToken{}
	if err = randomize.Struct(seed, authTokenOne, authTokenDBTypes, false, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}
	if err = randomize.Struct(seed, authTokenTwo, authTokenDBTypes, false, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authTokenOne := &AuthToken{}
	authTokenTwo := &AuthToken{}
	if err = randomize.Struct(seed, authTokenOne, authTokenDBTypes, false, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}
	if err = randomize.Struct(seed, authTokenTwo, authTokenDBTypes, false, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authTokenBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func authTokenAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthToken) error {
	*o = AuthToken{}
	return nil
}

func testAuthTokensHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AuthToken{}
	o := &AuthToken{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authTokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthToken object: %s", err)
	}

	AddAuthTokenHook(boil.BeforeInsertHook, authTokenBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authTokenBeforeInsertHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.AfterInsertHook, authTokenAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authTokenAfterInsertHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.AfterSelectHook, authTokenAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authTokenAfterSelectHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.BeforeUpdateHook, authTokenBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authTokenBeforeUpdateHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.AfterUpdateHook, authTokenAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authTokenAfterUpdateHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.BeforeDeleteHook, authTokenBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authTokenBeforeDeleteHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.AfterDeleteHook, authTokenAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authTokenAfterDeleteHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.BeforeUpsertHook, authTokenBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authTokenBeforeUpsertHooks = []AuthTokenHook{}

	AddAuthTokenHook(boil.AfterUpsertHook, authTokenAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authTokenAfterUpsertHooks = []AuthTokenHook{}
}

func testAuthTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authTokenColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthTokenToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AuthToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authTokenDBTypes, false, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthTokenSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*AuthToken)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthTokenToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authTokenDBTypes, false, strmangle.SetComplement(authTokenPrimaryKeyColumns, authTokenColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testAuthTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
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

func testAuthTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthTokenSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authTokenDBTypes = map[string]string{`ID`: `bigint`, `UserID`: `bigint`, `Token`: `character`, `Type`: `enum.token_type('auth')`, `DTExpired`: `timestamp without time zone`, `DTCreated`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testAuthTokensUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authTokenAllColumns) == len(authTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authTokenAllColumns) == len(authTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthToken{}
	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authTokenDBTypes, true, authTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authTokenAllColumns, authTokenPrimaryKeyColumns) {
		fields = authTokenAllColumns
	} else {
		fields = strmangle.SetComplement(
			authTokenAllColumns,
			authTokenPrimaryKeyColumns,
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

	slice := AuthTokenSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(authTokenAllColumns) == len(authTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AuthToken{}
	if err = randomize.Struct(seed, &o, authTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthToken: %s", err)
	}

	count, err := AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authTokenDBTypes, false, authTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthToken struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthToken: %s", err)
	}

	count, err = AuthTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
