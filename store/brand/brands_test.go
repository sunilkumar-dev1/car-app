package brand

import (
	"CarApp/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {

	data := []struct {
		input    entities.Brand
		hasError string
	}{
		{
			entities.Brand{1, "Maruti"},
			"Success",
		}, {
			entities.Brand{-1, "Audi"},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	var b BrandStorer
	b.brandDB = db
	defer b.brandDB.Close()

	rows := sqlmock.NewRows([]string{"id", "bname"}).AddRow(1, "Maruti").AddRow(1, "Hyundai").AddRow(-1, "Audi").
		AddRow(2, "Hyundai")

	for _, d := range data {
		mock.ExpectQuery("SELECT * FROM Brands where id = ?").
			WithArgs(d.input.Id).
			WillReturnRows(rows)
		_, err1 := b.GetById(d.input.Id)
		if err1.Error() != d.hasError {

			t.Errorf(err1.Error())
		}
	}
}

func TestCreate(t *testing.T) {
	data := []struct {
		input    entities.Brand
		hasError string
	}{
		{
			entities.Brand{1, "Maruti"},
			"Success",
		}, {
			entities.Brand{1, "Maruti"},
			"Failed",
		},
		{
			entities.Brand{-1, "Audi"},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	var b BrandStorer
	b.brandDB = db
	defer b.brandDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("INSERT INTO Brands (id,bname) values (?,?)").
			ExpectExec().
			WithArgs(d.input.Id, d.input.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := b.Create(&d.input)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}
func TestDelete(t *testing.T) {
	data := []struct {
		input    entities.Brand
		hasError string
	}{
		{
			entities.Brand{1, "Maruti"},
			"Success",
		}, {
			entities.Brand{2, "Hyundai"},
			"Success",
		},
	}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	var b BrandStorer
	b.brandDB = db
	defer b.brandDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("DELETE FROM Brands WHERE id=?").
			ExpectExec().
			WithArgs(d.input.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := b.Delete(d.input.Id)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}

func TestUpdate(t *testing.T) {
	data := []struct {
		input    entities.Brand
		hasError string
	}{
		{
			entities.Brand{1, "Maruti"},
			"Success",
		}, {
			entities.Brand{2, "Hyundai"},
			"Success",
		},
	}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	var b BrandStorer
	b.brandDB = db
	defer b.brandDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("UPDATE Brands SET bname=? WHERE id=?").
			ExpectExec().
			WithArgs(d.input.Id, d.input.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))

		_, err1 := b.Update(d.input.Id, d.input.Name)
		if err1.Error() != "Success" {
			t.Errorf(err1.Error())
		}
	}
}
