package services

import (
	"errors"
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"time"
)

type IUrlService interface {
	GetById(uint, string) dto.UrlDTO
	GetByShortUrlId(string, string) dto.UrlDTO
	GetByOriginalUrl(uint, string, string) dto.UrlDTO
	Create(uint, dto.CreateShortUrl, string) dto.UrlDTO
	Delete(uint, uint, string) dto.UrlDTO
	VisitUrl(string, string, string, string) dto.UrlDTO
	GetShortLinks(uint, string) []dto.UrlDTO
	GetUrlClicksByUserIdAndUrl(userId uint, shortUrlId string) []dto.UrlClickDTO
}

type UrlService struct {
	DB *gorm.DB
}

func (service UrlService) GetShortLinks(userId uint, hostname string) []dto.UrlDTO {
	var urls []models.Url

	res := service.DB.Where("user_id = ?", userId).Find(&urls)
	if res.Error != nil {
		panic(apperrors.InternalServerError(res.Error))
	}

	urlDTOs := make([]dto.UrlDTO, 0)

	for _, url := range urls {
		urlDTOs = append(urlDTOs, util.ToUrlDTO(url, hostname))
	}

	return urlDTOs
}

func (service UrlService) GetById(urlId uint, hostname string) dto.UrlDTO {
	url := models.Url{}
	res := service.DB.First(&url, urlId)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			panic(apperrors.NotFoundError("user not found"))
		}
		panic(res.Error)
	}

	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) GetByShortUrlId(shortUrlId string, hostname string) dto.UrlDTO {
	if shortUrlId == "" {
		panic(apperrors.BadRequestError("Invalid short url"))
	}
	url := service.getByShortUrlId(shortUrlId)

	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) getByShortUrlId(shortUrlId string) models.Url {
	url := models.Url{}
	result := service.DB.Where("short_url_id = ?", shortUrlId).Find(&url)
	if result.Error != nil {
		panic(result.Error)
	}

	return url
}

func (service UrlService) GetByOriginalUrl(userId uint, originalUrl string, hostname string) dto.UrlDTO {
	url := service.getByOriginalUrl(userId, originalUrl)

	if url.OriginalURL == "" {
		panic(apperrors.NotFoundError(fmt.Sprintf("short url for original url %s for user with id %d does not exist.", originalUrl, userId)))
	}
	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) Create(userId uint, request dto.CreateShortUrl, hostname string) dto.UrlDTO {
	originalUrl := request.OriginalURL
	if originalUrl == "" {
		panic(apperrors.BadRequestError("missing original url"))
	}
	if !util.IsValidUrl(originalUrl) {
		panic(apperrors.BadRequestError(fmt.Sprintf("invalid url %s.", originalUrl)))
	}
	originalUrl, err := util.NormalizeURL(originalUrl)
	if err != nil {
		panic(apperrors.BadRequestError("invalid url."))
	}
	url := service.getByOriginalUrl(userId, originalUrl)
	if url.ID != 0 {
		panic(apperrors.BadRequestError(fmt.Sprintf("shortened url exists for original url %s for user with id %d.", originalUrl, userId)))
	}
	shortUrlId := service.generateUniqueShortUrlId()
	url = models.Url{
		OriginalURL: originalUrl,
		ShortUrlId:  shortUrlId,
		ExpiryDate:  time.Now().AddDate(0, 0, 7),
		UserID:      userId,
	}

	result := service.DB.Create(&url)
	if result.Error != nil {
		panic(result.Error)
	}

	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) generateUniqueShortUrlId() string {
	shortUrlId, err := shortid.Generate()
	if err != nil {
		panic(apperrors.InternalServerError(err))
	}
	return shortUrlId
}

func (service UrlService) Delete(userId uint, urlId uint, hostname string) dto.UrlDTO {
	url := models.Url{}
	result := service.DB.Where("user_id = ? AND id = ?", userId, urlId).Find(&url)

	if result.Error != nil {
		panic(result.Error)
	}
	if url.OriginalURL == "" {
		panic(apperrors.NotFoundError("url is not found for user"))
	}
	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) VisitUrl(shortUrlId string, userAgent string, ipAddress string, hostname string) dto.UrlDTO {
	url := models.Url{}
	result := service.DB.Where("short_url_id = ?", shortUrlId).Find(&url)

	if result.Error != nil {
		panic(result.Error)
	}
	if url.OriginalURL == "" {
		panic(apperrors.NotFoundError("url is not found for user"))
	}

	urlClick := models.UrlClick{
		UserAgent: userAgent,
		IPAddress: ipAddress,
		UrlID:     url.ID,
	}

	result = service.DB.Save(&urlClick)
	if result.Error != nil {
		panic(result.Error)
	}
	return util.ToUrlDTO(url, hostname)
}

func (service UrlService) GetUrlClicksByIPAddress(shortUrlId string, ipAddress string) []dto.UrlClickDTO {
	url := service.getByShortUrlId(shortUrlId)
	if url.ShortUrlId == "" {
		panic(apperrors.NotFoundError("url not found."))
	}

	var urlClicks []models.UrlClick
	result := service.DB.Where("url_id = ? AND ip_address = ?", url.ID, ipAddress).Find(&urlClicks)

	if result.Error != nil {
		panic(result.Error)
	}

	clicks := make([]dto.UrlClickDTO, 0)
	for _, click := range urlClicks {
		clicks = append(clicks, util.ToUrlClickDTO(click))
	}

	return clicks
}

func (service UrlService) GetUrlClicksByUserIdAndUrl(userId uint, shortUrlId string) []dto.UrlClickDTO {
	url := service.getByShortUrlId(shortUrlId)
	if url.OriginalURL == "" || url.UserID != userId {
		panic(apperrors.NotFoundError("url not found."))
	}

	var urlClicks []models.UrlClick
	result := service.DB.Where("url_id = ?", url.ID).Find(&urlClicks)

	if result.Error != nil {
		panic(result.Error)
	}

	clicks := make([]dto.UrlClickDTO, 0)
	for _, click := range urlClicks {
		clicks = append(clicks, util.ToUrlClickDTO(click))
	}

	return clicks
}

func (service UrlService) getByOriginalUrl(userId uint, originalUrl string) models.Url {
	if originalUrl == "" {
		panic(apperrors.BadRequestError("Invalid url"))
	}
	url := models.Url{}
	result := service.DB.Where("user_id = ? AND original_url = ?", userId, originalUrl).Find(&url)

	if result.Error != nil {
		panic(result.Error)
	}

	return url
}
