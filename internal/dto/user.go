package dto

import "time"

type RegisterUserRequest struct {
	Name     string `name:"json"`
	Email    string `email:"json"`
	Password string `password:"json"`
}

type LoginUserRequest struct {
	Email    string `email:"json"`
	Password string `password:"json"`
}

type UpdateUserRequest struct {
	RegisterUserRequest
	OldPassword string `oldPassword:"json"`
}

type UserDTO struct {
	Name      string    `name:"json"`
	Email     string    `email:"json"`
	ID        uint      `id:"json"`
	CreatedAt time.Time `createdAt:"json"`
	UpdatedAt time.Time `updatedAt:"json"`
}
