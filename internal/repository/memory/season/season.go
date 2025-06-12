package serial

import (
	"errors"
	"otus/internal/model"
)

func NewRepository() *Repository {
	return &Repository{
		season: make(map[int]model.Season),
		nextId: 1,
	}
}

func (r *Repository) Save(season *model.Season) error {
	if season.Id == 0 {
		season.Id = r.nextId
		r.nextId++
	}
	r.season[season.Id] = *season
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.season, id)
	return nil
}

func (r *Repository) Load(id int) (*model.Season, error) {
	if season, ok := r.season[id]; ok {
		return &season, nil
	} else {
		return nil, errors.New(`season not found`)
	}
}

func (r *Repository) GetAll() ([]model.Season, error) {
	var seasons []model.Season
	for _, season := range r.season {
		seasons = append(seasons, season)
	}
	return seasons, nil
}
