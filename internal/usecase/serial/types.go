package serial

import (
	"otus/internal/model/serial"
)

type Usecase struct {
	repo serial.Repository
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
