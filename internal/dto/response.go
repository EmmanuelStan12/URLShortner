package dto

type ErrorResponse struct {
	Status  int    `status:"json"`
	Message string `message:"json"`
}

type SuccessResponse[T any] struct {
	Status int `status:"json"`
	Data   T   `data:"json"`
}
