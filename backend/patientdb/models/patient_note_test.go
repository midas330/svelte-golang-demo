// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPatientNotes(t *testing.T) {
	t.Parallel()

	query := PatientNotes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPatientNotesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
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

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientNotesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PatientNotes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientNotesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientNoteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientNotesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PatientNoteExists(ctx, tx, o.Noteid)
	if err != nil {
		t.Errorf("Unable to check if PatientNote exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PatientNoteExists to return true, but got false.")
	}
}

func testPatientNotesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	patientNoteFound, err := FindPatientNote(ctx, tx, o.Noteid)
	if err != nil {
		t.Error(err)
	}

	if patientNoteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPatientNotesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PatientNotes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPatientNotesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PatientNotes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPatientNotesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	patientNoteOne := &PatientNote{}
	patientNoteTwo := &PatientNote{}
	if err = randomize.Struct(seed, patientNoteOne, patientNoteDBTypes, false, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}
	if err = randomize.Struct(seed, patientNoteTwo, patientNoteDBTypes, false, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientNoteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientNoteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PatientNotes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPatientNotesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	patientNoteOne := &PatientNote{}
	patientNoteTwo := &PatientNote{}
	if err = randomize.Struct(seed, patientNoteOne, patientNoteDBTypes, false, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}
	if err = randomize.Struct(seed, patientNoteTwo, patientNoteDBTypes, false, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientNoteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientNoteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func patientNoteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func patientNoteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PatientNote) error {
	*o = PatientNote{}
	return nil
}

func testPatientNotesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PatientNote{}
	o := &PatientNote{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, patientNoteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PatientNote object: %s", err)
	}

	AddPatientNoteHook(boil.BeforeInsertHook, patientNoteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	patientNoteBeforeInsertHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.AfterInsertHook, patientNoteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	patientNoteAfterInsertHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.AfterSelectHook, patientNoteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	patientNoteAfterSelectHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.BeforeUpdateHook, patientNoteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	patientNoteBeforeUpdateHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.AfterUpdateHook, patientNoteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	patientNoteAfterUpdateHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.BeforeDeleteHook, patientNoteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	patientNoteBeforeDeleteHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.AfterDeleteHook, patientNoteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	patientNoteAfterDeleteHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.BeforeUpsertHook, patientNoteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	patientNoteBeforeUpsertHooks = []PatientNoteHook{}

	AddPatientNoteHook(boil.AfterUpsertHook, patientNoteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	patientNoteAfterUpsertHooks = []PatientNoteHook{}
}

func testPatientNotesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientNotesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(patientNoteColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientNoteToOnePatientUsingPatientid(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PatientNote
	var foreign Patient

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Patient_Id, foreign.Patientid)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Patientid().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.Patientid, foreign.Patientid) {
		t.Errorf("want: %v, got %v", foreign.Patientid, check.Patientid)
	}

	slice := PatientNoteSlice{&local}
	if err = local.L.LoadPatientid(ctx, tx, false, (*[]*PatientNote)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Patientid == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Patientid = nil
	if err = local.L.LoadPatientid(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Patientid == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPatientNoteToOneUserUsingUserid(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PatientNote
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.User_Id, foreign.Userid)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Userid().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.Userid, foreign.Userid) {
		t.Errorf("want: %v, got %v", foreign.Userid, check.Userid)
	}

	slice := PatientNoteSlice{&local}
	if err = local.L.LoadUserid(ctx, tx, false, (*[]*PatientNote)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Userid == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Userid = nil
	if err = local.L.LoadUserid(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Userid == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPatientNoteToOneSetOpPatientUsingPatientid(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PatientNote
	var b, c Patient

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientNoteDBTypes, false, strmangle.SetComplement(patientNotePrimaryKeyColumns, patientNoteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Patient{&b, &c} {
		err = a.SetPatientid(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Patientid != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PatientidPatientNotes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Patient_Id, x.Patientid) {
			t.Error("foreign key was wrong value", a.Patient_Id)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Patient_Id))
		reflect.Indirect(reflect.ValueOf(&a.Patient_Id)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Patient_Id, x.Patientid) {
			t.Error("foreign key was wrong value", a.Patient_Id, x.Patientid)
		}
	}
}

func testPatientNoteToOneRemoveOpPatientUsingPatientid(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PatientNote
	var b Patient

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientNoteDBTypes, false, strmangle.SetComplement(patientNotePrimaryKeyColumns, patientNoteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPatientid(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePatientid(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Patientid().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Patientid != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Patient_Id) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PatientidPatientNotes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPatientNoteToOneSetOpUserUsingUserid(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PatientNote
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientNoteDBTypes, false, strmangle.SetComplement(patientNotePrimaryKeyColumns, patientNoteColumnsWithoutDefault)...); err != nil {
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
		err = a.SetUserid(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Userid != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UseridPatientNotes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.User_Id, x.Userid) {
			t.Error("foreign key was wrong value", a.User_Id)
		}

		zero := reflect.Zero(reflect.TypeOf(a.User_Id))
		reflect.Indirect(reflect.ValueOf(&a.User_Id)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.User_Id, x.Userid) {
			t.Error("foreign key was wrong value", a.User_Id, x.Userid)
		}
	}
}

func testPatientNoteToOneRemoveOpUserUsingUserid(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PatientNote
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientNoteDBTypes, false, strmangle.SetComplement(patientNotePrimaryKeyColumns, patientNoteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUserid(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUserid(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Userid().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Userid != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.User_Id) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.UseridPatientNotes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPatientNotesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
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

func testPatientNotesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientNoteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatientNotesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PatientNotes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	patientNoteDBTypes = map[string]string{`Noteid`: `character`, `Patient_Id`: `character varying`, `User_Id`: `character varying`, `Note`: `character varying`, `Created`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testPatientNotesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(patientNotePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(patientNoteAllColumns) == len(patientNotePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNotePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPatientNotesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(patientNoteAllColumns) == len(patientNotePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PatientNote{}
	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNoteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientNoteDBTypes, true, patientNotePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(patientNoteAllColumns, patientNotePrimaryKeyColumns) {
		fields = patientNoteAllColumns
	} else {
		fields = strmangle.SetComplement(
			patientNoteAllColumns,
			patientNotePrimaryKeyColumns,
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

	slice := PatientNoteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPatientNotesUpsert(t *testing.T) {
	t.Parallel()

	if len(patientNoteAllColumns) == len(patientNotePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PatientNote{}
	if err = randomize.Struct(seed, &o, patientNoteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PatientNote: %s", err)
	}

	count, err := PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, patientNoteDBTypes, false, patientNotePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PatientNote struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PatientNote: %s", err)
	}

	count, err = PatientNotes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}