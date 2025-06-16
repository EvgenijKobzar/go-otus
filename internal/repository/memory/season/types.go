package season

import "otus/internal/model/season"

type IRepository interface {
	GetAll() ([]season.Entity, error)
	Load(id int) (*season.Entity, error)
	Save(entity *season.Entity) error
	Delete(id int) error
}

type Repository struct {
	items  map[int]season.Entity
	nextId int
}
