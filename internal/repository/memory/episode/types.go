package serial

import "otus/internal/model"

type Repository struct {
	episode map[int]model.Episode
	nextId  int
}
