package dto

import "time"

type RegisterUserRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type UpdateUserRequest struct {
	RegisterUserRequest
	OldPassword string
}

type UserDTO struct {
	Name      string
	Email     string
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
