package dto

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}
