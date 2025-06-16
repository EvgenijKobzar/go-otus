package usecase

import (
	e "otus/internal/lib/error"
	"otus/internal/model/catalog"
)

func SeasonCreate(params SeasonCreateParams, options ...catalog.SeasonOption) (catalog.Season, error) {

	s := catalog.NewSeason()

	for _, option := range options {
		option(&s)
	}

	if params.Title == "" {
		return s, e.ErrInvalidField("Title")
	}

	s.Title = params.Title

	return s, nil
}
