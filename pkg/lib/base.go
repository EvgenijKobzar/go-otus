package lib

import (
	"otus/internal/model/catalog"
	f "otus/internal/repository/file"
	"otus/pkg/lib/mapstructure"
)

func getInner[T catalog.HasId](id int) (T, error) {
	repo := f.NewRepository[T]()
	return repo.GetById(id)
}
func getListInner[T catalog.HasId]() ([]T, error) {
	repo := f.NewRepository[T]()
	return repo.GetAll()
}
func addInner[T catalog.HasId](binding *T) (*T, error) {
	var err error
	repo := f.NewRepository[T]()
	if err = repo.Save(*binding); err == nil {
		return binding, nil
	}
	return nil, err
}
func updateInner[T catalog.HasId](binding *T) (*T, error) {
	repo := f.NewRepository[T]()
	var err error
	if err = repo.Save(*binding); err == nil {
		return binding, nil
	}
	return nil, err
}
func deleteInner[T catalog.HasId](id int) error {
	var err error

	repo := f.NewRepository[T]()

	_, err = repo.GetById(id)

	if err == nil {
		err = repo.Delete(id)
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
