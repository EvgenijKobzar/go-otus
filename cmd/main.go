package main

import (
	"fmt"
	c "otus/internal/model/catalog"
	m "otus/internal/repository/memory"
	us "otus/internal/usecase"
)

func main() {
	repo := m.NewRepository[c.Serial]()

	episode, _ := us.NewUsecase(m.NewRepository[c.Episode]()).Create(us.EpisodeCreateParams{
		Title: "Фелина",
	})

	season, _ := us.NewUsecase(m.NewRepository[c.Season]()).Create(us.SeasonCreateParams{
		Title: "5 сезон",
	}, c.WithEpisode(&episode))

	us.NewUsecase(repo).Create(us.SerialCreateParams{
		Title: "Breaking Bad",
	}, c.WithSeason(&season))

	fmt.Println(repo.GetAll())
}
