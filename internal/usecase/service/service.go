package service

import (
	"github.com/gin-gonic/gin"
	"otus/internal/model/catalog"
	"otus/internal/repository"
	"otus/pkg/lib/mapstructure"
)

type Service[T catalog.HasId] struct {
	repo repository.IRepository[T]
}

func New[T catalog.HasId](repo repository.IRepository[T]) *Service[T] {
	return &Service[T]{repo: repo}
}

func (us *Service[T]) GetInner(id int) (T, error) {
	return us.repo.GetById(id)
}

func (us *Service[T]) GetListInner() ([]T, error) {
	items, err := us.repo.GetAll()
	return items, err
}

func (us *Service[T]) AddInner(binding *T) (*T, error) {
	var err error
	if err = us.repo.Save(*binding); err == nil {
		return binding, nil
	}
	return nil, err
}

func (us *Service[T]) UpdateInner(id int, c *gin.Context) (T, error) {
	var err error
	var entity T

	entity, err = us.repo.GetById(id)

	if err == nil {
		bindings := new(T)
		var inputFields map[string]any
		if err = c.ShouldBindJSON(&inputFields); err == nil {
			mapstructure.MapToStruct(inputFields, bindings)
			if err = entityAssign[T](entity, *bindings, inputFields); err == nil {
				err = us.repo.Save(entity)
			}
		}
	}
	return entity, err
}

func (us *Service[T]) DeleteInner(id int) error {
	var err error

	_, err = us.repo.GetById(id)

	if err == nil {
		err = us.repo.Delete(id)
	}
	return err
}

func entityAssign[T catalog.HasId](entity T, bindings T, allowedFields map[string]any) error {
	var err error
	var srcMap map[string]any
	var distMap map[string]any

	if srcMap, err = mapstructure.StructToMap(entity); err == nil {
		if distMap, err = mapstructure.StructToMap(bindings); err == nil {
			assignMap := assign(srcMap, distMap, allowedFields)
			err = mapstructure.MapToStruct(assignMap, entity)
		}
	}
	return err
}

func assign(src map[string]any, dist map[string]any, allowed map[string]any) map[string]any {
	result := make(map[string]any)
	for name, oldValue := range src {
		if _, isAllow := allowed[name]; isAllow {
			if newValue, exists := dist[name]; exists {
				result[name] = newValue
			}
		} else {
			result[name] = oldValue
		}
	}
	return result
}
