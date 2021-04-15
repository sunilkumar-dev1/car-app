package service

import "CarApp/entities"

// this interface using the interface of store layers
//go:generate mockgen -destination=mockBrand_interface.go -package=store . BrandInterface
type BrandService interface {
	GetById(id int) (entities.Brand, error)
	Create(brand entities.Brand) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
}

type ModelService interface {
	GetById(id int) (entities.Models, error)
	Create(model entities.Models) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
	ReadAll(all map[string]string) ([]entities.Models, error)
}

type VariantService interface {
	GetById(id int) (entities.Variant, error)
	Create(variant entities.Variant) (int, error)
	Delete(id int) (int, error)
	Update(id int, name string) (int, error)
	ReadAll(all map[string]string) ([]entities.Variant, error)
}
