package controller

import (
	"github.com/gin-gonic/gin"
	"otus/internal/lib/mapstructure"
	"otus/internal/middleware"
	"otus/internal/model/catalog"
	f "otus/internal/repository/file"
	"strconv"
)

const Item = "item"
const Items = "items"

// Action region
func GetAction[T catalog.HasId](c *gin.Context) {
	var entity T
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		entity, err = getInner[T](id)
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func AddAction[T catalog.HasId](c *gin.Context) {
	var entity *T
	bindings := new(T)

	err := c.ShouldBind(bindings)

	if err == nil {
		entity, err = addInner[T](bindings)
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func UpdateAction[T catalog.HasId](c *gin.Context) {
	var entity T
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		entity, err = updateInner[T](id, c)
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func DeleteAction[T catalog.HasId](c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		err = deleteInner[T](id)
	}
	setResponse(gin.H{"deleted": true}, err, c)
}

func GetListAction[T catalog.HasId](c *gin.Context) {
	var err error
	var result gin.H

	repo := f.NewRepository[T]()
	items, _ := repo.GetAll()

	result = gin.H{Items: items}

	setResponse(result, err, c)
}

// end region

func getInner[T catalog.HasId](id int) (T, error) {
	repo := f.NewRepository[T]()
	return repo.GetById(id)
}

func addInner[T catalog.HasId](binding *T) (*T, error) {
	var err error
	repo := f.NewRepository[T]()
	if err = repo.Save(*binding); err == nil {
		return binding, nil
	}
	return nil, err
}

func updateInner[T catalog.HasId](id int, c *gin.Context) (T, error) {
	var err error
	var entity T

	repo := f.NewRepository[T]()

	entity, err = repo.GetById(id)

	if err == nil {
		bindings := new(T)
		var inputFields map[string]any
		if err = c.ShouldBindJSON(&inputFields); err == nil {
			mapstructure.MapToStruct(inputFields, bindings)
			if err = entityAssign[T](entity, *bindings, inputFields); err == nil {
				err = repo.Save(entity)
			}
		}
	}
	return entity, err
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

func deleteInner[T catalog.HasId](id int) error {
	var err error

	repo := f.NewRepository[T]()

	_, err = repo.GetById(id)

	if err == nil {
		err = repo.Delete(id)
	}
	return err
}

func setResponse(result gin.H, err error, c *gin.Context) {
	if err == nil {
		c.Set(middleware.KeyResponse, result)
	} else {
		c.Set(middleware.KeyError, err)
	}
}
