package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"otus/internal/config"
	"otus/internal/model/catalog"
	"slices"
)

func NewRepository[T catalog.HasId]() *Repository[T] {
	r := &Repository[T]{
		items: make(map[int]T),
	}
	return r.refresh()
}

func (r *Repository[T]) Save(entity T) error {
	r.imx.Lock()
	defer r.imx.Unlock()
	if entity.GetId() == 0 {
		entity.SetId(r.nextId)
		r.nextId++
	}
	r.items[entity.GetId()] = entity

	r.saveToFile(entity)

	return nil
}

func (r *Repository[T]) Delete(id int) error {
	delete(r.items, id)

	var entity T
	r.saveToFile(entity)
	return nil
}

func (r *Repository[T]) Load(id int) (*T, error) {
	if entity, ok := r.items[id]; ok {
		return &entity, nil
	} else {
		return nil, errors.New(`entity not found`)
	}
}

func (r *Repository[T]) GetAll() ([]T, error) {
	r.imx.RLock()
	defer r.imx.RUnlock()

	return slices.Collect(maps.Values(r.items)), nil
}

func (r *Repository[T]) GetById(id int) (T, error) {
	items, _ := r.GetAll()
	for _, entity := range items {
		if entity.GetId() == id {
			return entity, nil
		}
	}
	var entity T
	return entity, errors.New(`entity not found`)
}

func (r *Repository[T]) Count() int {
	items, _ := r.GetAll()
	return len(items)
}

func (r *Repository[T]) refresh() *Repository[T] {
	items, _ := r.loadFromFile()

	for _, item := range items {
		r.items[item.GetId()] = item
	}
	r.nextId = r.Count() + 1
	return r
}

func (r *Repository[T]) loadFromFile() ([]T, error) {
	var entity T
	path, _ := config.ResolvePathByEntityType(entity)

	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return nil, err
	}

	var items []T
	err = json.Unmarshal(fileContent, &items)
	if err != nil {
		fmt.Printf("Ошибка декодирования JSON: %v\n", err)
		return nil, err
	}

	return items, nil
}

func (r *Repository[T]) saveToFile(entity T) error {
	path, _ := config.ResolvePathByEntityType(entity)

	items := slices.Collect(maps.Values(r.items))
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка кодирования JSON: %v\n", err)
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи файла: %v\n", err)
		return err
	}
	return nil
}
