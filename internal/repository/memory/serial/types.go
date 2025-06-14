package serial

import "otus/internal/model/serial"

type IRepository interface {
	GetAll() ([]serial.Entity, error)
	Load(id int) (*serial.Entity, error)
	Save(entity *serial.Entity) error
	Delete(id int) error
}
type Repository struct {
	items  map[int]serial.Entity
	nextId int
}
