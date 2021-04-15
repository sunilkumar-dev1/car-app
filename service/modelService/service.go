package modelService

import (
	"CarApp/entities"
	"CarApp/service"
	"CarApp/store"
)

type modelServe struct {
	ms store.ModelInterface
}

func New(m store.ModelInterface) service.ModelService {
	return modelServe{
		ms: m,
	}
}
func (m modelServe) GetById(id int) (entities.Models, error) {
	return m.ms.GetById(id)
}

func (m modelServe) Create(model entities.Models) (int, error) {
	return m.ms.Create(&model)
}

func (m modelServe) Delete(id int) (int, error) {
	return m.ms.Delete(id)
}

func (m modelServe) Update(id int, name string) (int, error) {
	return m.ms.Update(id, name)
}
func (m modelServe) ReadAll(all map[string]string) ([]entities.Models, error) {
	return m.ms.ReadAll(all)
}
