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
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = getSerial[T](id)
	}
	setResponse(result, err, c)
}

func AddAction[T catalog.HasId](c *gin.Context) {
	var result gin.H
	bindings := new(T)

	err := c.ShouldBindQuery(bindings)

	if err == nil {
		result, err = addInner[T](bindings)
	}
	setResponse(result, err, c)
}

func UpdateAction[T catalog.HasId](c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = updateInner[T](id, c)
	}
	setResponse(result, err, c)
}

func DeleteAction[T catalog.HasId](c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = deleteInner[T](id)
	}
	setResponse(result, err, c)
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

func getSerial[T catalog.HasId](id int) (gin.H, error) {
	var err error
	var result gin.H
	var entity T

	repo := f.NewRepository[T]()
	if entity, err = repo.GetById(id); err == nil {
		result = gin.H{Item: entity}
	}

	return result, err
}

func addInner[T catalog.HasId](binding *T) (gin.H, error) {
	var err error
	var result gin.H
	repo := f.NewRepository[T]()
	if err = repo.Save(*binding); err == nil {
		result = gin.H{Item: binding}
	}
	return result, err
}

func updateInner[T catalog.HasId](id int, c *gin.Context) (gin.H, error) {
	var err error
	var result gin.H
	var entity T

	repo := f.NewRepository[T]()

	entity, err = repo.GetById(id)

	if err == nil {
		bindings := new(T)

		if err = c.ShouldBindQuery(bindings); err == nil {
			allowedFields := c.QueryMap("fields")
			if err = entityAssign[T](entity, *bindings, allowedFields); err == nil {
				if err = repo.Save(entity); err == nil {
					result = gin.H{Item: entity}
				}
			}
		}
	}
	return result, err
}

func entityAssign[T catalog.HasId](entity T, bindings T, allowedFields map[string]string) error {
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

func assign(src map[string]any, dist map[string]any, allowed map[string]string) map[string]any {
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

func deleteInner[T catalog.HasId](id int) (gin.H, error) {
	var err error
	var result gin.H

	repo := f.NewRepository[T]()

	_, err = repo.GetById(id)

	if err == nil {
		if err = repo.Delete(id); err == nil {
			result = gin.H{"deleted": true}
		}
	}
	return result, err
}

func setResponse(result gin.H, err error, c *gin.Context) {
	if err == nil {
		c.Set(middleware.KeyResponse, result)
	} else {
		c.Set(middleware.KeyError, err)
	}
}
