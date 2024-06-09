package util

import (
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
)

func ToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		Name:      user.Name,
		Email:     user.Email,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
