package brandService

import (
	"CarApp/entities"
	"CarApp/service"
	"CarApp/store"
)

type brandServe struct {
	bs store.BrandInterface // handler of store type
}

func New(b store.BrandInterface) service.BrandService {
	return brandServe{bs: b}
}

func (b brandServe) GetById(id int) (entities.Brand, error) {
	return b.bs.GetById(id)
}

func (b brandServe) Create(brand entities.Brand) (int, error) {
	return b.bs.Create(&brand)
}

func (b brandServe) Delete(id int) (int, error) {
	return b.bs.Delete(id)
}

func (b brandServe) Update(id int, name string) (int, error) {
	return b.bs.Update(id, name)
}
