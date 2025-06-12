package serial

import (
	"otus/internal/model"
)

type Repository interface {
	Load(id int) (*model.Serial, error)
	Save(serial *model.Serial) error
	Delete(id int) error
}
