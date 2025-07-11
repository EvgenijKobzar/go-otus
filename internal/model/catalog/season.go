package catalog

type Season struct {
	Id        int    `json:"id"`
	Title     string `json:"title" binding:"required" form:"fields[title]"`
	SerialId  int    `json:"serial_id" form:"fields[serial_id]"`
	Sort      int    `json:"sort" form:"fields[sort]"`
	Moderated bool   `json:"moderated" form:"fields[moderated]"`
	CreatedBy int    `json:"created_by" form:"fields[created_by]"`
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
