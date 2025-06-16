package episode

import (
	e "otus/internal/lib/error"
	model "otus/internal/model/episode"
	repo "otus/internal/repository/memory/episode"
)

func NewUsecase(repo repo.IRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (uc *Usecase) Create(params CreateParams) (*model.Entity, error) {

	if params.Title == "" {
		return nil, e.ErrInvalidField("Title")
	}

	s := model.NewEpisode()
	s.Title = params.Title

	if err := uc.repo.Save(s); err != nil {
		return nil, err
	}
	return s, nil
}
