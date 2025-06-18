package catalog

type Season struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	SerialId  int    `json:"serial_id,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	Moderated bool   `json:"moderated,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	episodes  map[int]Episode
}

func (s *Season) GetId() int {
	return s.Id
}

func (s *Season) SetId(id int) {
	s.Id = id
}

func NewSeason() *Season {
	return &Season{
		episodes: make(map[int]Episode),
	}
}

type SeasonOption func(*Season)

func WithEpisode(episode *Episode) SeasonOption {
	return func(opts *Season) {
		opts.episodes[episode.Id] = *episode
	}
}
