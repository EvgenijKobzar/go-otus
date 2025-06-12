package season

import (
	"otus/internal/model"
)

type Repository interface {
	Load(id int) (*model.Season, error)
	Save(serial *model.Season) error
	Delete(id int) error
}
