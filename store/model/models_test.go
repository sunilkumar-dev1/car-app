package model

import (
	"CarApp/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	data := []struct {
		input    entities.Models
		hasError string
	}{
		{
			entities.Models{1, "Swift", 1},
			"Success",
		}, {
			entities.Models{2, "Suzki", 2},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var m ModelStorer
	m.modelDB = db
	defer m.modelDB.Close()

	rows := sqlmock.NewRows([]string{"id", "mname", "brandId"}).
		AddRow(1, "Swift", 1).AddRow(2, "Suzki", 2)

	for _, d := range data {
		mock.ExpectQuery("SELECT * FROM Model where id = ?").
			WithArgs(d.input.Id).
			WillReturnRows(rows)
		_, err1 := m.GetById(d.input.Id)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}
func TestCreate(t *testing.T) {
	data := []struct {
		input    entities.Models
		hasError string
	}{
		{
			entities.Models{1, "Swift", 1},
			"Success",
		}, {
			entities.Models{2, "Suzki", 2},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var m ModelStorer
	m.modelDB = db
	defer m.modelDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("INSERT INTO Model (id,mname,brandId) values (?,?,?)").
			ExpectExec().
			WithArgs(d.input.Id, d.input.Name, d.input.BrandId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		_, err1 := m.Create(&d.input)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}

func TestDelete(t *testing.T) {

	data := []struct {
		input    entities.Models
		hasError string
	}{
		{
			entities.Models{1, "Swift", 1},
			"Success",
		}, {
			entities.Models{2, "Suzki", 2},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var m ModelStorer
	m.modelDB = db
	defer m.modelDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("DELETE FROM Model WHERE id=?").
			ExpectExec().
			WithArgs(d.input.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		_, err1 := m.Delete(d.input.Id)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}

func TestUpdate(t *testing.T) {

	data := []struct {
		input    entities.Models
		hasError string
	}{
		{
			entities.Models{1, "Swift", 1},
			"Success",
		}, {
			entities.Models{2, "Suzki", 2},
			"Success",
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}

	var m ModelStorer
	m.modelDB = db
	defer m.modelDB.Close()

	for _, d := range data {
		mock.ExpectPrepare("UPDATE Model SET mname=? WHERE id=?").
			ExpectExec().
			WithArgs(d.input.Id, d.input.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))

		_, err1 := m.Update(d.input.Id, d.input.Name)
		if err1.Error() != d.hasError {
			t.Errorf(err1.Error())
		}
	}
}
