package variant

import (
	"CarApp/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {

	data := []struct {
		input    entities.Variant
		hasError string
	}{
		{
			entities.Variant{1, "Renault", 30, "50", "70", 2},
			"Success",
		}, {
			entities.Variant{2, "Safari", 30, "40", "60", 1},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var v VariantStorer
	v.variantDB = db
	defer v.variantDB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "displace", "peakPower", "peakTorque", "modelId"}).
		AddRow(1, "Renault", 30, "50", "70", 2).AddRow(2, "Safari", 30, "40", "60", 1)

	for _, d := range data {
		mock.ExpectQuery("SELECT * FROM Variant where id = ?").
			WithArgs(d.input.Id).WillReturnRows(rows)
		_, err1 := v.GetById(d.input.Id)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}

func TestCreate(t *testing.T) {
	data := []struct {
		input    entities.Variant
		hasError string
	}{
		{
			entities.Variant{1, "Renault", 30, "50", "70", 2},
			"Success",
		}, {
			entities.Variant{2, "Safari", 30, "40", "60", 1},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var v VariantStorer
	v.variantDB = db
	defer v.variantDB.Close()

	for _, d := range data {

		mock.ExpectPrepare("INSERT INTO Variant (id,name,displace,peakPower,peakTorque,modelId) values (?,?,?,?,?,?)").
			ExpectExec().
			WithArgs(d.input.Id, d.input.Name, d.input.Displace, d.input.PeakPower, d.input.PeakTorque, d.input.ModelsId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := v.Create(&d.input)

		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}

}

func TestDelete(t *testing.T) {
	data := []struct {
		input    entities.Variant
		hasError string
	}{
		{
			entities.Variant{1, "Renault", 30, "50", "70", 2},
			"Success",
		}, {
			entities.Variant{2, "Safari", 30, "40", "60", 1},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	var v VariantStorer
	v.variantDB = db
	defer v.variantDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("DELETE FROM Variant WHERE id=?").
			ExpectExec().WithArgs(d.input.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := v.Delete(d.input.Id)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}

func TestUpdate(t *testing.T) {
	data := []struct {
		input    entities.Variant
		hasError string
	}{
		{
			entities.Variant{1, "Renault", 30, "50", "70", 2},
			"Success",
		}, {
			entities.Variant{2, "Safari", 30, "40", "60", 1},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var v VariantStorer
	v.variantDB = db
	defer v.variantDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("UPDATE Variant SET name=? WHERE id=?").
			ExpectExec().WithArgs(d.input.Id, d.input.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := v.Update(d.input.Id, d.input.Name)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}
