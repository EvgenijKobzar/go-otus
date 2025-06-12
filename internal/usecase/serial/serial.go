package serial

import (
	e "otus/internal/lib/error"
	"otus/internal/model"
	"otus/internal/model/serial"
)

func NewUsecase(repo serial.Repository) *Usecase {
	return &Usecase{repo: repo}
}

//func (uc *Usecase) NewSerial(options ...model.Option) *model.Serial {
//	opts := &model.Serial{Title: "test"} // значения по умолчанию
//	for _, option := range options {
//		option(opts)
//	}
//	return opts
//}

func (uc *Usecase) Create(params CreateParams, options ...model.SerialOption) (*model.Serial, error) {
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
