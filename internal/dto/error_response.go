package dto

type ErrorResponse struct {
	Status  int    `status:"json"`
	Message string `message:"json"`
}
