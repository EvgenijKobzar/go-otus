package serial

import (
	e "otus/internal/lib/error"
	"otus/internal/model"
	"otus/internal/model/episode"
)

func NewUsecase(repo episode.Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (uc *Usecase) Create(params CreateParams) (*model.Episode, error) {
	if params.Title == "" {
		return nil, e.ErrInvalidField("Title")
	}

	s := &model.Episode{
		Title: params.Title,
	}

	if err := uc.repo.Save(s); err != nil {
		return nil, err
	}
	return s, nil
}
