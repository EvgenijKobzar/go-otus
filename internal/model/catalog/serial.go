package catalog

type Serial struct {
	Id               int     `json:"id"`
	Title            string  `json:"title" binding:"required" form:"fields[title]"`
	FileId           int     `json:"file_id" form:"fields[fileId]"`
	Description      string  `json:"description" form:"fields[description]"`
	Rating           float64 `json:"rating" form:"fields[rating]"`
	Duration         float64 `json:"duration" form:"fields[duration]"`
	Sort             int     `json:"sort" form:"fields[sort]"`
	ProductionPeriod string  `json:"production_period" form:"fields[productionPeriod]"`
	Quality          string  `json:"quality" form:"fields[quality]"`
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
