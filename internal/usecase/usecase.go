package usecase

import (
	"otus/internal/model/catalog"
	"otus/internal/repository/memory"
)

func NewUsecase[T catalog.HasId](repo *memory.Repository[T]) *Usecase[T] {
	return &Usecase[T]{
		repo: repo,
	}
}

func (uc *Usecase[T]) Create(params any, options ...any) (T, error) {
	var entity T

	switch any(entity).(type) {
	case catalog.Serial:
		var serialOptions []catalog.SerialOption

		for _, option := range options {
			if o, ok := option.(catalog.SerialOption); ok {
				serialOptions = append(serialOptions, o)
			}
		}

		if createParams, ok := params.(SerialCreateParams); ok {
			serial, _ := SerialCreate(createParams, serialOptions...)
			entity = any(serial).(T)
		}

	case catalog.Season:
		var seasonOptions []catalog.SeasonOption

		for _, option := range options {
			if o, ok := option.(catalog.SeasonOption); ok {
				seasonOptions = append(seasonOptions, o)
			}
		}

		if createParams, ok := params.(SeasonCreateParams); ok {
			season, _ := SeasonCreate(createParams, seasonOptions...)
			entity = any(season).(T)
		}
	case catalog.Episode:
		if createParams, ok := params.(EpisodeCreateParams); ok {
			episode, _ := EpisodeCreate(createParams)
			entity = any(episode).(T)
		}
	}

	err := uc.repo.Save(entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}
