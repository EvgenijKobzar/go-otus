package episode

import (
	"otus/internal/model"
)

type Repository interface {
	Load(id int) (*model.Episode, error)
	Save(serial *model.Episode) error
	Delete(id int) error
}
