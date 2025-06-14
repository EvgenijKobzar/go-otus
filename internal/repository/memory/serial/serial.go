package serial

import (
	"errors"
	"otus/internal/model/serial"
)

func NewRepository() *Repository {
	return &Repository{
		items:  make(map[int]serial.Entity),
		nextId: 1,
	}
}

func (r *Repository) Save(serial *serial.Entity) error {
	if serial.Id == 0 {
		serial.Id = r.nextId
		r.nextId++
	}
	r.items[serial.Id] = *serial
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.items, id)
	return nil
}

func (r *Repository) Load(id int) (*serial.Entity, error) {
	if entity, ok := r.items[id]; ok {
		return &entity, nil
	} else {
		return nil, errors.New(`entity not found`)
	}
}

func (r *Repository) GetAll() ([]serial.Entity, error) {
	var items []serial.Entity
	for _, entity := range r.items {
		items = append(items, entity)
	}
	return items, nil
}
