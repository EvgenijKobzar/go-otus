package lib

import (
	"otus/internal/model/catalog"
	"otus/pkg/lib/mapstructure"
)

func Get(id int) (*catalog.Serial, error) {
	return getInner[*catalog.Serial](id)
}
func GetList() ([]*catalog.Serial, error) {
	return getListInner[*catalog.Serial]()
}
func Add(fields map[string]interface{}) (map[string]interface{}, error) {
	s := catalog.NewSerial()

	mapstructure.MapToStruct(fields, &s)

	_, err := addInner[*catalog.Serial](&s)

	m, _ := mapstructure.StructToMap(s)

	return m, err
}
func Update(id int, inputFields map[string]any) (map[string]interface{}, error) {
	entity, err := Get(id)
	if err != nil {
		return nil, err
	}

	s := catalog.NewSerial()
	mapstructure.MapToStruct(inputFields, &s)

	if err = entityAssign[*catalog.Serial](entity, s, inputFields); err == nil {
		updateInner(&entity)
		m, _ := mapstructure.StructToMap(entity)

		return m, err
	}
	return nil, err
}
func Delete(id int) error {
	return deleteInner[*catalog.Serial](id)
}
