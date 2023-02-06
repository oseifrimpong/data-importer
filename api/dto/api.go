package dto

type APIResponse struct {
	StatusCode int64       `json:"status_code,string"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
