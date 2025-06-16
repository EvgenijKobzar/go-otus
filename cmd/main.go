package main

import (
	"fmt"
	modelSeason "otus/internal/model/season"
	modelSerial "otus/internal/model/serial"
	repoEpisode "otus/internal/repository/memory/episode"
	repoSeason "otus/internal/repository/memory/season"
	repoSerial "otus/internal/repository/memory/serial"
	ue "otus/internal/usecase/episode"
	usn "otus/internal/usecase/season"
	us "otus/internal/usecase/serial"
)

func main() {

	repo := repoSerial.NewRepository()

	episode, _ := ue.NewUsecase(repoEpisode.NewRepository()).Create(ue.CreateParams{
		Title: "Фелина",
	})

	season, _ := usn.NewUsecase(repoSeason.NewRepository()).Create(usn.CreateParams{
		Title: "5 сезон",
	},
		modelSeason.WithEpisode(episode))

	us.NewUsecase(repo).Create(us.CreateParams{
		Title: "Breaking Bad",
	},
		modelSerial.WithSeason(season))

	fmt.Println(repo.GetAll())
}
