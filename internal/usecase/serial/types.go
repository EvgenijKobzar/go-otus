package serial

import (
	repo "otus/internal/repository/memory/serial"
)

type Usecase struct {
	repo repo.IRepository
}

type CreateParams struct {
	Id               int
	Title            string
	FileId           int
	Description      string
	Rating           float64
	Duration         float64
	Sort             int
	ProductionPeriod string
	Quality          string
}
