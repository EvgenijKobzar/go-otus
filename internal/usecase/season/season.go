package serial

import (
	e "otus/internal/lib/error"
	"otus/internal/model"
	"otus/internal/model/season"
)

func NewUsecase(repo season.Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (uc *Usecase) Create(params CreateParams, options ...model.SeasonOption) (*model.Season, error) {

	s := model.NewSeason()

	for _, option := range options {
		option(s)
	}

	if params.Title == "" {
		return nil, e.ErrInvalidField("Title")
	}

	s.Id = params.Id
	s.Title = params.Title
	s.CreatedBy = params.CreatedBy

	if err := uc.repo.Save(s); err != nil {
		return nil, err
	}
	return s, nil
}
