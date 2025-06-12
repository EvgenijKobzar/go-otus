package model

type Serial struct {
	Id               int     `json:"id,omitempty"`
	Title            string  `json:"title,omitempty"`
	FileId           int     `json:"file_id,omitempty"`
	Description      string  `json:"description,omitempty"`
	Rating           float64 `json:"rating,omitempty"`
	Duration         float64 `json:"duration,omitempty"`
	Sort             int     `json:"sort,omitempty"`
	ProductionPeriod string  `json:"production_period,omitempty"`
	Quality          string  `json:"quality,omitempty"`
	seasons          map[int]Season
}

type SerialOption func(*Serial)

func WithSeason(season *Season) SerialOption {
	return func(opts *Serial) {
		opts.seasons[season.Id] = *season
	}
}

func NewSerial() *Serial {
	return &Serial{
		seasons: make(map[int]Season),
	}
}

type Season struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	SerialId  int    `json:"serial_id,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	Moderated bool   `json:"moderated,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	episodes  map[int]Episode
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

type Episode struct {
	Id             int     `json:"id,omitempty"`
	Title          string  `json:"title,omitempty"`
	FileId         int     `json:"fileId,omitempty"`
	SeasonId       int     `json:"seasonId,omitempty"`
	SerialId       int     `json:"serialId,omitempty"`
	Src            string  `json:"src,omitempty"`
	Description    string  `json:"description,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Sort           int     `json:"sort,omitempty"`
	Rating         float64 `json:"rating,omitempty"`
	ProductionDate string  `json:"productionDate,omitempty"`
	Quality        string  `json:"quality,omitempty"`
	Moderated      bool    `json:"moderated,omitempty"`
	CreatedBy      int     `json:"created_by,omitempty"`
}

type Account struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
