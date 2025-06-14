package season

import (
	"errors"
	"otus/internal/model/season"
)

func NewRepository() *Repository {
	return &Repository{
		items:  make(map[int]season.Entity),
		nextId: 1,
	}
}

func (r *Repository) Save(entity *season.Entity) error {
	if entity.Id == 0 {
		entity.Id = r.nextId
		r.nextId++
	}
	r.items[entity.Id] = *entity
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.items, id)
	return nil
}

func (r *Repository) Load(id int) (*season.Entity, error) {
	if entity, ok := r.items[id]; ok {
		return &entity, nil
	} else {
		return nil, errors.New(`entity not found`)
	}
}

func (r *Repository) GetAll() ([]season.Entity, error) {
	var items []season.Entity
	for _, entity := range r.items {
		items = append(items, entity)
	}
	return items, nil
}
