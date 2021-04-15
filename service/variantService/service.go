package variantService

import (
	"CarApp/entities"
	"CarApp/service"
	"CarApp/store"
)

type variantServe struct {
	vs store.VariantInterface
}

func New(v store.VariantInterface) service.VariantService {
	return variantServe{
		vs: v,
	}
}
func (v variantServe) GetById(id int) (entities.Variant, error) {
	return v.vs.GetById(id)
}

func (v variantServe) Create(variant entities.Variant) (int, error) {
	return v.vs.Create(&variant)
}

func (v variantServe) Delete(id int) (int, error) {
	return v.vs.Delete(id)
}

func (v variantServe) Update(id int, name string) (int, error) {
	return v.vs.Update(id, name)
}
func (v variantServe) ReadAll(all map[string]string) ([]entities.Variant, error) {
	return v.vs.ReadAll(all)
}
