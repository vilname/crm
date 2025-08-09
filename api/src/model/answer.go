package model

type Answer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PaginationAnswer struct {
	Data       []Answer   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (pagination *PaginationAnswer) GetPaginationAnswer(list []Answer, page uint16, limit uint16, total uint16) {
	pagination.Data = list
	pagination.Pagination = Pagination{
		Page:  page,
		Limit: limit,
		Total: total,
	}
}
