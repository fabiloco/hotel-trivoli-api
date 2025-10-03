package entities

type Pagination struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Pages int         `json:"pages"`
	Data  interface{} `json:"data"`
}
