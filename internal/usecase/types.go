package usecase

import (
	"otus/internal/model/catalog"
	"otus/internal/repository/memory"
)

type Usecase[T catalog.HasId] struct {
	repo *memory.Repository[T]
}

type SerialCreateParams struct {
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

type SeasonCreateParams struct {
	Id        int
	Title     string
	SerialId  int
	Sort      int
	Moderated bool
	CreatedBy int
}

type EpisodeCreateParams struct {
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
