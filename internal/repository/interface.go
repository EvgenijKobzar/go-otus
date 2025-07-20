package repository

import "otus/internal/model/catalog"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mocks

type IRepository[T catalog.HasId] interface {
	GetAll() ([]T, error)
	GetById(id int) (T, error)
	Save(entity T) error
	Delete(id int) error
	Count() int
}
