package episode

import (
	repo "otus/internal/repository/memory/episode"
)

type Usecase struct {
	repo repo.IRepository
}

type CreateParams struct {
	Id             int
	Title          string
	FileId         int
	SeasonId       int
	SerialId       int
	Src            string
	Description    string
	Duration       float64
	Sort           int
	Rating         float64
	ProductionDate string
	Quality        string
	Moderated      bool
	CreatedBy      int
}
