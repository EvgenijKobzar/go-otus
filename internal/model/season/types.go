package season

import "otus/internal/model/episode"

type Entity struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	SerialId  int    `json:"serial_id,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	Moderated bool   `json:"moderated,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	episodes  map[int]episode.Entity
}

func NewSeason() *Entity {
	return &Entity{
		episodes: make(map[int]episode.Entity),
	}
}

type Option func(*Entity)

func WithEpisode(episode *episode.Entity) Option {
	return func(opts *Entity) {
		opts.episodes[episode.Id] = *episode
	}
}
