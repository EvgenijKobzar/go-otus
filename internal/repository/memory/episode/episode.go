package episode

import (
	"errors"
	"otus/internal/model/episode"
)

func NewRepository() *Repository {
	return &Repository{
		entity: make(map[int]episode.Entity),
		nextId: 1,
	}
}

func (r *Repository) Save(entity *episode.Entity) error {
	if entity.Id == 0 {
		entity.Id = r.nextId
		r.nextId++
	}
	r.entity[entity.Id] = *entity
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.entity, id)
	return nil
}

func (r *Repository) Load(id int) (*episode.Entity, error) {
	if entity, ok := r.entity[id]; ok {
		return &entity, nil
	} else {
		return nil, errors.New(`entity not found`)
	}
}

func (r *Repository) GetAll() ([]episode.Entity, error) {
	var items []episode.Entity
	for _, entity := range r.entity {
		items = append(items, entity)
	}
	return items, nil
}
