package store

import "CarApp/entities"

type BrandInterface interface {
	GetById(id int) (entities.Brand, error)
	Create(brand *entities.Brand) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
}

type ModelInterface interface {
	GetById(id int) (entities.Models, error)
	Create(model *entities.Models) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
	ReadAll(all map[string]string) ([]entities.Models, error)
}

type VariantInterface interface {
	GetById(id int) (entities.Variant, error)
	Create(variant *entities.Variant) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
	ReadAll(all map[string]string) ([]entities.Variant, error)
}
