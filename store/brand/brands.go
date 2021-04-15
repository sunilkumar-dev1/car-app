package brand

import (
	"CarApp/entities"
	"CarApp/store"
	"database/sql"
	"errors"
)

type BrandStorer struct {
	brandDB *sql.DB
}

func New(db *sql.DB) store.BrandInterface {
	return BrandStorer{brandDB: db} // this become interface when we return interface and can call interface.methods
}

// Get method imlementation
func (b BrandStorer) GetById(id int) (entities.Brand, error) {
	var brand entities.Brand
	rows, err := b.brandDB.Query("SELECT * FROM Brands where id = ?", id)
	if err != nil {
		return brand, errors.New("Failed")
	}
	if rows.Next() {
		_ = rows.Scan(&brand.Id, &brand.Name)
		return brand, errors.New("Success")
	}
	return brand, errors.New("Failed")
}

func (b BrandStorer) Create(brand *entities.Brand) (int, error) {
	rows, _ := b.brandDB.Prepare("INSERT INTO Brands (id,bname) values (?,?)")

	affect, err := rows.Exec(brand.Id, brand.Name)
	if err != nil {
		return 0, errors.New("Failed")
	}
	_, err1 := affect.RowsAffected()
	if err1 == nil {
		return 1, errors.New("Success")
	}
	return 0, errors.New("Failed")
}

func (b BrandStorer) Delete(id int) (int, error) {
	deleteQuery, err := b.brandDB.Prepare("DELETE FROM Brands WHERE id=?")
	if err != nil {
		return 0, errors.New("Failed")
	}
	result, err := deleteQuery.Exec(id)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 0 {
		return 1, errors.New("Success")
	}
	return 0, errors.New("Failed")
}

func (b BrandStorer) Update(id int, name string) (int, error) {
	updateQuery, err := b.brandDB.Prepare("UPDATE Brands SET bname=? WHERE id=?")

	if err != nil {
		return 0, errors.New("Failed")
	} else {
		result, _ := updateQuery.Exec(name, id)
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return 0, errors.New("Failed")
		} else {
			return 1, errors.New("Success")
		}
	}
	return 0, errors.New("Failed")
}
