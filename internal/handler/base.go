package handler

type DeleteResponse struct {
	Result struct {
		Deleted bool
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"entity not found"`
}
