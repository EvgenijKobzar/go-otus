package catalog

type Serial struct {
	Id               int     `json:"id" example:"1"`
	Title            string  `json:"title" binding:"required" form:"fields[title]" example:"Breaking Bad"`
	FileId           int     `json:"file_id" form:"fields[fileId]" example:"0"`
	Description      string  `json:"description" form:"fields[description]" example:"TV series"`
	Rating           float64 `json:"rating" form:"fields[rating]" example:"9.5"`
	Duration         float64 `json:"duration" form:"fields[duration]" example:"40"`
	Sort             int     `json:"sort" form:"fields[sort]" example:"1"`
	ProductionPeriod string  `json:"production_period" form:"fields[productionPeriod]" example:"2008-2013"`
	Quality          string  `json:"quality" form:"fields[quality]" example:"High"`
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
