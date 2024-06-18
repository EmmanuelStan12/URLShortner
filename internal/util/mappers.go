package util

import (
	"fmt"
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

func ToUrlDTO(url models.Url, hostname string) dto.UrlDTO {
	return dto.UrlDTO{
		OriginalURL: url.OriginalURL,
		ShortURL:    fmt.Sprintf("%s/short/%s", hostname, url.ShortUrlId),
		ExpiryDate:  url.ExpiryDate,
		ID:          url.ID,
		UserID:      url.UserID,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
	}
}

func ToUrlClickDTO(urlClick models.UrlClick) dto.UrlClickDTO {
	return dto.UrlClickDTO{
		ID:        urlClick.ID,
		CreatedAt: urlClick.CreatedAt,
		UpdatedAt: urlClick.UpdatedAt,
		UserAgent: urlClick.UserAgent,
		IPAddress: urlClick.IPAddress,
		UrlID:     urlClick.UrlID,
	}
}
