package catalog

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

func NewEpisode() *Episode {
	return &Episode{}
}

func (e *Episode) GetId() int {
	return e.Id
}

func (e *Episode) SetId(id int) {
	e.Id = id
}
