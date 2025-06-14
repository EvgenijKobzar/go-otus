package episode

import "otus/internal/model/episode"

type IRepository interface {
	GetAll() ([]episode.Entity, error)
	Load(id int) (*episode.Entity, error)
	Save(entity *episode.Entity) error
	Delete(id int) error
}

type Repository struct {
	entity map[int]episode.Entity
	nextId int
}
