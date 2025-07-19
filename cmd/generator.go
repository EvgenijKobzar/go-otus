package generator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	c "otus/internal/model/catalog"
	"otus/internal/repository"
	f "otus/internal/repository/file"
	us "otus/internal/usecase"
	"slices"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func startLogger(done <-chan bool, ticker *time.Ticker, repoEpisode *f.Repository[*c.Episode]) {
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
}

func createEpisode(ch chan<- *c.Episode, titles [8]string, repoEpisode repository.IRepository[*c.Episode]) {
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
}

func createSeason(ch <-chan *c.Episode, chSeason chan<- *c.Season) {
	season, _ := us.NewUsecase(f.NewRepository[*c.Season]()).Create(us.SeasonCreateParams{
		Title: "5 сезон",
	})

	for episode := range ch {
		c.WithEpisode(episode)(season)
	}
	chSeason <- season
}

func createSerial(season *c.Season, title string, repo repository.IRepository[*c.Serial]) (*c.Serial, error) {
	serial, err := us.NewUsecase(repo).Create(us.SerialCreateParams{
		Title: title,
	}, c.WithSeason(season))
	return serial, err
}

func generateSerial(ctx context.Context, repo repository.IRepository[*c.Serial], wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down...")
			time.Sleep(500 * time.Millisecond)
			return
		default:
			rand.Seed(time.Now().UnixNano())
			title := strconv.Itoa(int(time.Duration(rand.Intn(1000))))

			serial, _ := createSerial(c.NewSeason(), title, repo)
			log.Printf("Worker %d working...", serial.GetId())
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	repo := f.NewRepository[*c.Serial]()
	repoEpisode := f.NewRepository[*c.Episode]()

	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)
	go startLogger(done, ticker, repoEpisode)

	chEpisode := make(chan *c.Episode)
	chSeason := make(chan *c.Season)
	var titles = [8]string{"Фелина", "Гранитный штат", "Озимандия", "Тохаджилли", "Бешеный пёс", "Признания", "Зарытое", "Кровавые деньги"}

	go createEpisode(chEpisode, titles, repoEpisode)
	go createSeason(chEpisode, chSeason)

	season := <-chSeason

	createSerial(season, "Breaking Bad", repo)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go generateSerial(ctx, repo, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	log.Println("Received shutdown signal")

	cancel()

	shutdownSuccess := make(chan struct{})
	go func() {
		wg.Wait()
		close(shutdownSuccess)
	}()

	select {
	case <-shutdownSuccess:
		log.Println("All workers stopped gracefully")
	case <-time.After(5 * time.Second):
		log.Println("Shutdown timeout, some workers may not have stopped cleanly")
	}

	fmt.Println("Результат")
	items, _ := repo.GetAll()
	for _, item := range items {
		fmt.Println(*item)
	}

	ticker.Stop()
	done <- true
}
