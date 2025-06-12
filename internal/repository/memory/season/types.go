package serial

import "otus/internal/model"

type Repository struct {
	season map[int]model.Season
	nextId int
}
