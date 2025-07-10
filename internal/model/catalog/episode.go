package catalog

type Episode struct {
	Id             int     `json:"id"`
	Title          string  `json:"title" binding:"required" form:"fields[title]"`
	FileId         int     `json:"fileId" form:"fields[fileId]"`
	SeasonId       int     `json:"seasonId" form:"fields[seasonId]"`
	SerialId       int     `json:"serialId" form:"fields[serialId]"`
	Src            string  `json:"src" form:"fields[src]"`
	Description    string  `json:"description" form:"fields[description]"`
	Duration       float64 `json:"duration" form:"fields[duration]"`
	Sort           int     `json:"sort" form:"fields[sort]"`
	Rating         float64 `json:"rating" form:"fields[rating]"`
	ProductionDate string  `json:"productionDate" form:"fields[productionDate]"`
	Quality        string  `json:"quality" form:"fields[quality]"`
	Moderated      bool    `json:"moderated" form:"fields[moderated]"`
	CreatedBy      int     `json:"created_by" form:"fields[created_by]"`
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
