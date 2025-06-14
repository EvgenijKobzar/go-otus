package season

import (
	repo "otus/internal/repository/memory/season"
)

type Usecase struct {
	repo repo.IRepository
}

type CreateParams struct {
	Id        int
	Title     string
	SerialId  int
	Sort      int
	Moderated bool
	CreatedBy int
}
