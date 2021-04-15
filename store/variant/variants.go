package variant

import (
	"CarApp/entities"
	"CarApp/store"
	"database/sql"
	"errors"
	"strings"
)

type VariantStorer struct {
	variantDB *sql.DB
}

func New(db *sql.DB) store.VariantInterface {
	return VariantStorer{variantDB: db}
}

// Get method imlementation
func (v VariantStorer) GetById(id int) (entities.Variant, error) {
	var variant entities.Variant
	rows, err := v.variantDB.Query("SELECT * FROM Variant where id = ?", id)
	if err != nil {
		return variant, errors.New("Failed")
	}
	if rows.Next() {
		_ = rows.Scan(&variant.Id, &variant.Name, &variant.Displace, &variant.PeakPower, &variant.PeakTorque, &variant.ModelsId)
		return variant, errors.New("Success")
	}
	return variant, errors.New("Failed")
}

func (m VariantStorer) Create(variant *entities.Variant) (int, error) {
	rows, err := m.variantDB.Prepare("INSERT INTO Variant (id,name,displace,peakPower,peakTorque,modelId) values (?,?,?,?,?,?)")
	affect, err := rows.Exec(variant.Id, variant.Name, variant.Displace, variant.PeakPower, variant.PeakTorque, variant.ModelsId)
	if err != nil {
		return 0, errors.New("Failed")
	}
	_, err1 := affect.RowsAffected()
	if err1 != nil {
		return 0, errors.New("Failed")
	}
	return 1, errors.New("Success")
}

func (m VariantStorer) Delete(id int) (int, error) {
	deleteQuery, err := m.variantDB.Prepare("DELETE FROM Variant WHERE id=?")
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

func (m VariantStorer) Update(id int, name string) (int, error) {
	updateQuery, err := m.variantDB.Prepare("UPDATE Variant SET name=? WHERE id=?")

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

func (m VariantStorer) ReadAll(all map[string]string) ([]entities.Variant, error) {
	var s string = "select * from Variant where "
	for i, v := range all {
		s += i + "=" + v + " and "
	}
	s = strings.Trim(s, "and ")
	rows, err := m.variantDB.Query(s)
	if err != nil {
		return []entities.Variant{}, nil
	}
	var res []entities.Variant
	for rows.Next() {
		var md entities.Variant
		err := rows.Scan(&md.Id, &md.Name, &md.Displace, &md.PeakPower, &md.PeakTorque, &md.ModelsId)
		if err != nil {
			return []entities.Variant{}, err
		}
		res = append(res, md)
	}
	return res, nil
}
