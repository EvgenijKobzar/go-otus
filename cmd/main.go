package main

import (
	"fmt"
	c "otus/internal/model/catalog"
	m "otus/internal/repository/memory"
	us "otus/internal/usecase"
)

func main() {
	repo := m.NewRepository[*c.Serial]()
	repoEpisode := m.NewRepository[*c.Episode]()

	episode, _ := us.NewUsecase(repoEpisode).Create(us.EpisodeCreateParams{
		Title: "Фелина",
	})

	episode2, _ := us.NewUsecase(repoEpisode).Create(us.EpisodeCreateParams{
		Title: "Фелина2",
	})

	season, _ := us.NewUsecase(m.NewRepository[*c.Season]()).Create(us.SeasonCreateParams{
		Title: "5 сезон",
	}, c.WithEpisode(episode), c.WithEpisode(episode2))

	us.NewUsecase(repo).Create(us.SerialCreateParams{
		Title: "Breaking Bad",
	}, c.WithSeason(season))

	items, _ := repo.GetAll()
	for _, item := range items {
		fmt.Println(*item)
	}
	//fmt.Println(items)
}
