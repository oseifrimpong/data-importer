package dto

type APIResponse struct {
	StatusCode int32       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type SearchParams struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Sort     string `json:"sort"`
	Unix     int64  `json:"unix"`
	Symbol   string `json:"symbol"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
}
