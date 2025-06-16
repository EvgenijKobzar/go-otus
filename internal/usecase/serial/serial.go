package serial

import (
	e "otus/internal/lib/error"
	model "otus/internal/model/serial"
	repo "otus/internal/repository/memory/serial"
)

func NewUsecase(repo repo.IRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (uc *Usecase) Create(params CreateParams, options ...model.Option) (*model.Entity, error) {

	s := model.NewSerial()

	for _, option := range options {
		option(s)
	}

	if params.Title == "" {
		return nil, e.ErrInvalidField("Title")
	}

	s.Id = params.Id
	s.Title = params.Title

	if err := uc.repo.Save(s); err != nil {
		return nil, err
	}
	return s, nil
}
