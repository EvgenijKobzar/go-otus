package sqlc

import (
	"otus/internal/model/catalog"
)

type Repository[T catalog.HasId] struct {
	repo *Queries
}
