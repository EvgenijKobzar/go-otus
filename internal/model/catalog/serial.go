package catalog

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

func (s *Serial) GetId() int {
	return s.Id
}

func (s *Serial) SetId(id int) {
	s.Id = id
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
