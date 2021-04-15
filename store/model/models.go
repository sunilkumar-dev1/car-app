package model

import (
	"CarApp/entities"
	"CarApp/store"
	"database/sql"
	"errors"
	"strings"
)

type ModelStorer struct {
	modelDB *sql.DB
}

func New(db *sql.DB) store.ModelInterface {
	return ModelStorer{modelDB: db}
}

// Get method imlementation
func (m ModelStorer) GetById(id int) (entities.Models, error) {
	var model entities.Models
	rows, err := m.modelDB.Query("SELECT * FROM Model where id = ?", id)
	if err != nil {
		return model, errors.New("Failed")
	}
	if rows.Next() {
		_ = rows.Scan(&model.Id, &model.Name, &model.BrandId)
		return model, errors.New("Success")
	}
	return model, errors.New("Failed")
}

func (m ModelStorer) Create(model *entities.Models) (int, error) {
	rows, err := m.modelDB.Prepare("INSERT INTO Model (id,mname,brandId) values (?,?,?)")

	affect, err := rows.Exec(model.Id, model.Name, model.BrandId)
	if err != nil {
		return 0, errors.New("Failed")
	}
	_, err1 := affect.RowsAffected()
	if err1 != nil {
		return 0, errors.New("Failed")
	}
	return 1, errors.New("Success")
}

func (m ModelStorer) Delete(id int) (int, error) {
	deleteQuery, err := m.modelDB.Prepare("DELETE FROM Model WHERE id=?")
	if err != nil {
		return 0, errors.New("failed")
	}
	result, err := deleteQuery.Exec(id)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 0 {
		return 1, errors.New("Success")
	}
	return 0, errors.New("failed")
}

func (m ModelStorer) Update(id int, name string) (int, error) {
	updateQuery, err := m.modelDB.Prepare("UPDATE Model SET mname=? WHERE id=?")

	if err != nil {
		return 0, errors.New("failed")
	} else {
		result, _ := updateQuery.Exec(name, id)
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return 0, errors.New("failed")
		} else {
			return 1, errors.New("Success")
		}
	}
	return 0, errors.New("failed")
}

func (m ModelStorer) ReadAll(all map[string]string) ([]entities.Models, error) {
	var s string = "select * from Model where "
	for i, v := range all {
		s += i + "=" + v + " and "
	}
	s = strings.Trim(s, "and ")
	rows, err := m.modelDB.Query(s)
	if err != nil {
		return []entities.Models{}, nil
	}
	var res []entities.Models
	for rows.Next() {
		var md entities.Models
		err := rows.Scan(&md.Id, &md.Name, &md.BrandId)
		if err != nil {
			return []entities.Models{}, err
		}
		res = append(res, md)
	}
	return res, nil
}
