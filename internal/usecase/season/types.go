package serial

import (
	"otus/internal/model/season"
)

type Usecase struct {
	repo season.Repository
}

type CreateParams struct {
	Id        int
	Title     string
	SerialId  int
	Sort      int
	Moderated bool
	CreatedBy int
}
