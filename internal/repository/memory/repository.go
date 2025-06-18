package memory

import (
	"errors"
	"otus/internal/model/catalog"
)

func NewRepository[T catalog.HasId]() *Repository[T] {
	return &Repository[T]{
		nextId: 1,
		items:  make(map[int]T),
	}
}

func (r *Repository[T]) Save(entity T) error {
	r.imx.Lock()
	defer r.imx.Unlock()
	if entity.GetId() == 0 {
		entity.SetId(r.nextId)
		r.nextId++
	}
	r.items[entity.GetId()] = entity
	return nil
}

func (r *Repository[T]) Delete(id int) error {
	delete(r.items, id)
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
	r.imx.RUnlock()
	var items []T
	for _, entity := range r.items {
		items = append(items, entity)
	}
	return items, nil
}
