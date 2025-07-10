package controller

import (
	"github.com/gin-gonic/gin"
	"otus/internal/lib/mapstructure"
	"otus/internal/middleware"
	"otus/internal/model/catalog"
	f "otus/internal/repository/file"
	"strconv"
)

const Item = "serial"
const Items = "serials"

// http://localhost:8080/v1/otus.serial.update?id=53&fields[title]=123
// http://localhost:8080/v1/otus.serial.get?id=1
// http://localhost:8080/v1/otus.serial.add?fields[id]=test1&fields[title]=test4&fields[sort]=4
// http://localhost:8080/v1/otus.serial.delete?id=57
// http://localhost:8080/v1/otus.serial.list

// Action region
func GetSerialAction(c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = getSerial(id)
	}
	setResponse(result, err, c)
}

func AddSerialAction(c *gin.Context) {
	var result gin.H
	var binding catalog.Serial
	err := c.ShouldBindQuery(&binding)

	if err == nil {
		result, err = addSerial(binding)
	}
	setResponse(result, err, c)
}

func UpdateSerialAction(c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = updateSerial(id, c)
	}
	setResponse(result, err, c)
}

func DeleteSerialAction(c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		result, err = deleteSerial(id)
	}
	setResponse(result, err, c)
}

func GetListSerialAction(c *gin.Context) {
	var err error
	var result gin.H

	repo := f.NewRepository[*catalog.Serial]()
	items, _ := repo.GetAll()

	result = gin.H{Items: items}

	setResponse(result, err, c)
}

// end region
func getSerial(id int) (gin.H, error) {
	var err error
	var result gin.H
	var entity *catalog.Serial

	repo := f.NewRepository[*catalog.Serial]()
	if entity, err = repo.GetById(id); err == nil {
		result = gin.H{Item: entity}
	}

	return result, err
}

func addSerial(binding catalog.Serial) (gin.H, error) {
	var err error
	var result gin.H
	repo := f.NewRepository[*catalog.Serial]()
	if err = repo.Save(&binding); err == nil {
		result = gin.H{Item: binding}
	}
	return result, err
}

func updateSerial(id int, c *gin.Context) (gin.H, error) {
	var err error
	var result gin.H
	var entity *catalog.Serial

	repo := f.NewRepository[*catalog.Serial]()

	entity, err = repo.GetById(id)

	if err == nil {
		var bindings catalog.Serial

		if err = c.ShouldBindQuery(&bindings); err == nil {
			allowedFields := c.QueryMap("fields")
			if err = entityAssign(entity, bindings, allowedFields); err == nil {
				if err = repo.Save(entity); err == nil {
					result = gin.H{Item: entity}
				}
			}
		}
	}
	return result, err
}

func entityAssign(entity *catalog.Serial, bindings catalog.Serial, allowedFields map[string]string) error {
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

func deleteSerial(id int) (gin.H, error) {
	var err error
	var result gin.H
	var _ *catalog.Serial

	repo := f.NewRepository[*catalog.Serial]()

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
