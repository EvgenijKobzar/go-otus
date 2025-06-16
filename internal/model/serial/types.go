package serial

import "otus/internal/model/season"

type Entity struct {
	Id               int     `json:"id,omitempty"`
	Title            string  `json:"title,omitempty"`
	FileId           int     `json:"file_id,omitempty"`
	Description      string  `json:"description,omitempty"`
	Rating           float64 `json:"rating,omitempty"`
	Duration         float64 `json:"duration,omitempty"`
	Sort             int     `json:"sort,omitempty"`
	ProductionPeriod string  `json:"production_period,omitempty"`
	Quality          string  `json:"quality,omitempty"`
	seasons          map[int]season.Entity
}

type Option func(*Entity)

func WithSeason(season *season.Entity) Option {
	return func(opts *Entity) {
		opts.seasons[season.Id] = *season
	}
}

func NewSerial() *Entity {
	return &Entity{
		seasons: make(map[int]season.Entity),
	}
}
