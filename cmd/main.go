package main

import (
	"fmt"
	c "otus/internal/model/catalog"
	m "otus/internal/repository/memory"
	us "otus/internal/usecase"
	"sync"
	"time"
)

func main() {
	repo := m.NewRepository[*c.Serial]()

	ch := make(chan *c.Episode)
	var titles = [8]string{"Фелина", "Гранитный штат", "Озимандия", "Тохаджилли", "Бешеный пёс", "Признания", "Зарытое", "Кровавые деньги"}

	go func(titles [8]string) {
		repoEpisode := m.NewRepository[*c.Episode]()

		wg := sync.WaitGroup{}
		wg.Add(len(titles))
		for _, title := range titles {
			go func(name string) {
				defer wg.Done()
				us.NewUsecase(repoEpisode).Create(us.EpisodeCreateParams{
					Title: name,
				})
			}(title)
		}
		wg.Wait()

		episodes, _ := repoEpisode.GetAll()
		time.Sleep(1000 * time.Millisecond)

		for _, episode := range episodes {

			ch <- episode
		}
		close(ch)
	}(titles)

	chSeason := make(chan *c.Season)
	go func() {
		season, _ := us.NewUsecase(m.NewRepository[*c.Season]()).Create(us.SeasonCreateParams{
			Title: "5 сезон",
		})

		for episode := range ch {
			c.WithEpisode(episode)(season)
		}
		chSeason <- season
	}()

	season := <-chSeason

	us.NewUsecase(repo).Create(us.SerialCreateParams{
		Title: "Breaking Bad",
	}, c.WithSeason(season))

	items, _ := repo.GetAll()
	for _, item := range items {
		fmt.Println(*item)
	}
}
