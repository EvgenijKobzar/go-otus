package usecase

import (
	e "otus/internal/lib/error"
	"otus/internal/model/catalog"
)

func SerialCreate(params SerialCreateParams, options ...catalog.SerialOption) (catalog.Serial, error) {

	s := catalog.NewSerial()

	for _, option := range options {
		option(&s)
	}

	if params.Title == "" {
		return s, e.ErrInvalidField("Title")
	}

	s.Id = params.Id
	s.Title = params.Title

	return s, nil
}
