package main

import (
	"fmt"
	"otus/internal/model"
	repe "otus/internal/repository/memory/episode"
	repsn "otus/internal/repository/memory/season"
	reps "otus/internal/repository/memory/serial"
	ue "otus/internal/usecase/episode"
	usn "otus/internal/usecase/season"
	us "otus/internal/usecase/serial"
)

func main() {

	repo := reps.NewRepository()

	episode, _ := ue.NewUsecase(repe.NewRepository()).Create(ue.CreateParams{
		Title: "Фелина",
	})

	season, _ := usn.NewUsecase(repsn.NewRepository()).Create(usn.CreateParams{
		Title: "5 сезон",
	},
		model.WithEpisode(episode))

	us.NewUsecase(repo).Create(us.CreateParams{
		Title: "Breaking Bad",
	},
		model.WithSeason(season))

	fmt.Println(repo.GetAll())
}
