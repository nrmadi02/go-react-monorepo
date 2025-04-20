package dto

type SuccessResponse struct {
	Data    any    `json:"data"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}
