package serial

import "otus/internal/model"

type Repository struct {
	serials map[int]model.Serial
	nextId  int
}
