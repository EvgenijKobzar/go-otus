package memory

import "otus/internal/model/catalog"

type IRepository[T catalog.HasId] interface {
	GetAll() ([]T, error)
	Load(id int) (*T, error)
	Save(entity *T) error
	Delete(id int) error
}

type Repository[T catalog.HasId] struct {
	items  map[int]T
	nextId int
}
