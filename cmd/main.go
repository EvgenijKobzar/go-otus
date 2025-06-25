package main

import (
	"fmt"
	"math/rand"
	c "otus/internal/model/catalog"
	m "otus/internal/repository/memory"
	us "otus/internal/usecase"
	"slices"
	"sync"
	"time"
)

func main() {
	repo := m.NewRepository[*c.Serial]()
	repoEpisode := m.NewRepository[*c.Episode]()

	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)

	go func() {
		var inxList []int

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				cnt := repoEpisode.Count()

				if cnt > len(inxList) {
					fmt.Println("->Что-то новое")

					items, _ := repoEpisode.GetAll()
					for _, episode := range items {
						if slices.Contains(inxList, episode.GetId()) == false {
							fmt.Println(*episode)
							inxList = append(inxList, episode.GetId())
						}
					}
				} else {
					fmt.Println("->Изменений нет")
				}
			}
		}
	}()

	ch := make(chan *c.Episode)
	var titles = [8]string{"Фелина", "Гранитный штат", "Озимандия", "Тохаджилли", "Бешеный пёс", "Признания", "Зарытое", "Кровавые деньги"}

	go func(titles [8]string) {

		wg := sync.WaitGroup{}
		wg.Add(len(titles))
		for _, title := range titles {
			go func(name string) {
				defer wg.Done()

				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

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

	fmt.Println("Результат")
	items, _ := repo.GetAll()
	for _, item := range items {
		fmt.Println(*item)
	}

	ticker.Stop()
	done <- true
}
