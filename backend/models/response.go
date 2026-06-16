package models

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count,omitempty"`
}

func Success(data interface{}) APIResponse {
	return APIResponse{Success: true, Data: data}
}

func SuccessWithCount(data interface{}, count int) APIResponse {
	return APIResponse{Success: true, Data: data, Count: count}
}

func Error(msg string) APIResponse {
	return APIResponse{Success: false, Error: msg}
}
