package dto

type APIResponse struct {
	StatusCode int32       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
