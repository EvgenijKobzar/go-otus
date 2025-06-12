package serial

import (
	"errors"
	"otus/internal/model"
)

func NewRepository() *Repository {
	return &Repository{
		serials: make(map[int]model.Serial),
		nextId:  1,
	}
}

func (r *Repository) Save(serial *model.Serial) error {
	if serial.Id == 0 {
		serial.Id = r.nextId
		r.nextId++
	}
	r.serials[serial.Id] = *serial
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.serials, id)
	return nil
}

func (r *Repository) Load(id int) (*model.Serial, error) {
	if serial, ok := r.serials[id]; ok {
		return &serial, nil
	} else {
		return nil, errors.New(`serial not found`)
	}
}

func (r *Repository) GetAll() ([]model.Serial, error) {
	var serials []model.Serial
	for _, serial := range r.serials {
		serials = append(serials, serial)
	}
	return serials, nil
}

//func (r *Repository) GetById(id int) (*catalog.Serial, error) {
//	if serial, ok := r.serials[id]; ok {
//		return &serial, nil
//	} else {
//		return nil, errors.New(`serial not found`)
//	}
//}
