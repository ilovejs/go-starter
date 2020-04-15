// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRoleNotificationTemplates(t *testing.T) {
	t.Parallel()

	query := RoleNotificationTemplates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoleNotificationTemplatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
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

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleNotificationTemplatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RoleNotificationTemplates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleNotificationTemplatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleNotificationTemplateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleNotificationTemplatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoleNotificationTemplateExists(ctx, tx, o.RoleID, o.NotificationTemplateID)
	if err != nil {
		t.Errorf("Unable to check if RoleNotificationTemplate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoleNotificationTemplateExists to return true, but got false.")
	}
}

func testRoleNotificationTemplatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roleNotificationTemplateFound, err := FindRoleNotificationTemplate(ctx, tx, o.RoleID, o.NotificationTemplateID)
	if err != nil {
		t.Error(err)
	}

	if roleNotificationTemplateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoleNotificationTemplatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RoleNotificationTemplates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoleNotificationTemplatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RoleNotificationTemplates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoleNotificationTemplatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roleNotificationTemplateOne := &RoleNotificationTemplate{}
	roleNotificationTemplateTwo := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, roleNotificationTemplateOne, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}
	if err = randomize.Struct(seed, roleNotificationTemplateTwo, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleNotificationTemplateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleNotificationTemplateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleNotificationTemplates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoleNotificationTemplatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roleNotificationTemplateOne := &RoleNotificationTemplate{}
	roleNotificationTemplateTwo := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, roleNotificationTemplateOne, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}
	if err = randomize.Struct(seed, roleNotificationTemplateTwo, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleNotificationTemplateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleNotificationTemplateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testRoleNotificationTemplatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleNotificationTemplatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(roleNotificationTemplateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleNotificationTemplateToOneNotificationTemplateUsingNotificationTemplate(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleNotificationTemplate
	var foreign NotificationTemplate

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, notificationTemplateDBTypes, false, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.NotificationTemplateID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.NotificationTemplate().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoleNotificationTemplateSlice{&local}
	if err = local.L.LoadNotificationTemplate(ctx, tx, false, (*[]*RoleNotificationTemplate)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.NotificationTemplate == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.NotificationTemplate = nil
	if err = local.L.LoadNotificationTemplate(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.NotificationTemplate == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleNotificationTemplateToOneRoleUsingRole(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleNotificationTemplate
	var foreign Role

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoleID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Role().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoleNotificationTemplateSlice{&local}
	if err = local.L.LoadRole(ctx, tx, false, (*[]*RoleNotificationTemplate)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Role = nil
	if err = local.L.LoadRole(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleNotificationTemplateToOneSetOpNotificationTemplateUsingNotificationTemplate(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleNotificationTemplate
	var b, c NotificationTemplate

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleNotificationTemplateDBTypes, false, strmangle.SetComplement(roleNotificationTemplatePrimaryKeyColumns, roleNotificationTemplateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, notificationTemplateDBTypes, false, strmangle.SetComplement(notificationTemplatePrimaryKeyColumns, notificationTemplateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, notificationTemplateDBTypes, false, strmangle.SetComplement(notificationTemplatePrimaryKeyColumns, notificationTemplateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*NotificationTemplate{&b, &c} {
		err = a.SetNotificationTemplate(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.NotificationTemplate != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleNotificationTemplates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.NotificationTemplateID != x.ID {
			t.Error("foreign key was wrong value", a.NotificationTemplateID)
		}

		if exists, err := RoleNotificationTemplateExists(ctx, tx, a.RoleID, a.NotificationTemplateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testRoleNotificationTemplateToOneSetOpRoleUsingRole(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleNotificationTemplate
	var b, c Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleNotificationTemplateDBTypes, false, strmangle.SetComplement(roleNotificationTemplatePrimaryKeyColumns, roleNotificationTemplateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Role{&b, &c} {
		err = a.SetRole(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Role != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleNotificationTemplates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID)
		}

		if exists, err := RoleNotificationTemplateExists(ctx, tx, a.RoleID, a.NotificationTemplateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testRoleNotificationTemplatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
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

func testRoleNotificationTemplatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleNotificationTemplateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoleNotificationTemplatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleNotificationTemplates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roleNotificationTemplateDBTypes = map[string]string{`RoleID`: `uuid`, `NotificationTemplateID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                               = bytes.MinRead
)

func testRoleNotificationTemplatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(roleNotificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roleNotificationTemplateAllColumns) == len(roleNotificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoleNotificationTemplatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roleNotificationTemplateAllColumns) == len(roleNotificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleNotificationTemplate{}
	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleNotificationTemplateDBTypes, true, roleNotificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roleNotificationTemplateAllColumns, roleNotificationTemplatePrimaryKeyColumns) {
		fields = roleNotificationTemplateAllColumns
	} else {
		fields = strmangle.SetComplement(
			roleNotificationTemplateAllColumns,
			roleNotificationTemplatePrimaryKeyColumns,
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

	slice := RoleNotificationTemplateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoleNotificationTemplatesUpsert(t *testing.T) {
	t.Parallel()

	if len(roleNotificationTemplateAllColumns) == len(roleNotificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RoleNotificationTemplate{}
	if err = randomize.Struct(seed, &o, roleNotificationTemplateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleNotificationTemplate: %s", err)
	}

	count, err := RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roleNotificationTemplateDBTypes, false, roleNotificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleNotificationTemplate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleNotificationTemplate: %s", err)
	}

	count, err = RoleNotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
