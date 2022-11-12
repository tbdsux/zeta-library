package lib

type APIResponse struct {
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
