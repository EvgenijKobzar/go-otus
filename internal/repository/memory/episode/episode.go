package serial

import (
	"errors"
	"otus/internal/model"
)

func NewRepository() *Repository {
	return &Repository{
		episode: make(map[int]model.Episode),
		nextId:  1,
	}
}

func (r *Repository) Save(episode *model.Episode) error {
	if episode.Id == 0 {
		episode.Id = r.nextId
		r.nextId++
	}
	r.episode[episode.Id] = *episode
	return nil
}

func (r *Repository) Delete(id int) error {
	delete(r.episode, id)
	return nil
}

func (r *Repository) Load(id int) (*model.Episode, error) {
	if episode, ok := r.episode[id]; ok {
		return &episode, nil
	} else {
		return nil, errors.New(`episode not found`)
	}
}

func (r *Repository) GetAll() ([]model.Episode, error) {
	var episodes []model.Episode
	for _, episode := range r.episode {
		episodes = append(episodes, episode)
	}
	return episodes, nil
}
