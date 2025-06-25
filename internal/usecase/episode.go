package usecase

import (
	e "otus/internal/lib/error"
	"otus/internal/model/catalog"
)

func EpisodeCreate(params EpisodeCreateParams) (*catalog.Episode, error) {

	s := catalog.NewEpisode()

	if params.Title == "" {
		return s, e.ErrInvalidField("Title")
	}

	s.Title = params.Title

	return s, nil
}
