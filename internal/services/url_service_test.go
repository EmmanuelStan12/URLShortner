package services

import (
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"testing"
	"time"
)

func initTestUrl(db *gorm.DB, userId uint) *models.Url {
	shortUrlId, _ := shortid.Generate()
	url := models.Url{
		OriginalURL: "https://test.com",
		ShortUrlId:  shortUrlId,
		ExpiryDate:  time.Now().AddDate(0, 0, 7),
		UserID:      userId,
	}
	db.Create(&url)
	return &url
}

func TestUrlService_Create(t *testing.T) {
	db, conf := initTestDB(t)
	user := initTestUser(db)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}

	t.Run("test create url panic tests", func(t *testing.T) {
		panicTests := []struct {
			request dto.CreateShortUrl
			name    string
		}{
			{
				request: dto.CreateShortUrl{},
				name:    "create without original url",
			},
			{
				request: dto.CreateShortUrl{
					OriginalURL: "http://invalid.com/",
				},
				name: "create with invalid original url",
			},
		}

		for _, tt := range panicTests {
			t.Run(tt.name, func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					urlService.Create(user.ID, tt.request, conf.Server.Hostname)
				})
			})
		}
	})

	t.Run("create url", func(t *testing.T) {
		request := dto.CreateShortUrl{
			OriginalURL: "https://www.youtube.com/",
		}

		urlService.Create(user.ID, request, conf.Server.Hostname)
	})
}

func TestUrlService_CreateForDiverseUrls(t *testing.T) {
	db, conf := initTestDB(t)
	user := initTestUser(db)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}

	t.Run("panic when there is an existing url", func(t *testing.T) {
		request := dto.CreateShortUrl{
			OriginalURL: "https://www.youtube.com/",
		}

		urlService.Create(user.ID, request, conf.Server.Hostname)

		urls := []string{
			"https://www.youtube.com",
			"https://www.youtube.com/",
			"https://youtube.com/",
			"https://youtube.com",
			"www.youtube.com",
			"https://YouTube.com",
		}

		for _, u := range urls {
			t.Run(fmt.Sprintf("panic for already existing url %s", u), func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					request := dto.CreateShortUrl{
						OriginalURL: u,
					}

					urlService.Create(user.ID, request, conf.Server.Hostname)
				})
			})
		}
	})
}

func TestUrlService_GetByOriginalUrl(t *testing.T) {
	db, conf := initTestDB(t)
	user := initTestUser(db)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}

	t.Run("test get original url panic tests", func(t *testing.T) {
		panicTests := []struct {
			originalUrl string
			name        string
		}{
			{
				originalUrl: "",
				name:        "empty original url",
			},
			{
				originalUrl: "https://test2.com",
				name:        "an original url that is not found",
			},
		}

		for _, tt := range panicTests {
			t.Run(tt.name, func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					urlService.GetByOriginalUrl(user.ID, tt.originalUrl, conf.Server.Hostname)
				})
			})
		}
	})

	t.Run("get original url", func(t *testing.T) {
		url := initTestUrl(db, user.ID)
		urlService.getByOriginalUrl(user.ID, url.OriginalURL)
	})
}

func TestUrlService_VisitUrl(t *testing.T) {
	db, _ := initTestDB(t)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}

	t.Run("test visit short url panic tests", func(t *testing.T) {
		panicTests := []struct {
			shortUrlId string
			userAgent  string
			ipAddress  string
			name       string
		}{
			{
				shortUrlId: "",
				name:       "empty short url",
			},
			{
				shortUrlId: "1234",
				name:       "a short url that is not found",
			},
		}

		for _, tt := range panicTests {
			t.Run(tt.name, func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					urlService.VisitUrl(tt.shortUrlId, tt.userAgent, tt.ipAddress, "my host name")
				})
			})
		}
	})

	t.Run("visit short url", func(t *testing.T) {
		user := initTestUser(db)
		url := initTestUrl(db, user.ID)
		for i := 0; i < 5; i++ {
			urlService.VisitUrl(url.ShortUrlId, "test user agent", "test ip address", "my host name")
		}

		clicks := urlService.GetUrlClicksByIPAddress(url.ShortUrlId, "test ip address")
		if len(clicks) != 5 {
			t.Errorf("Invalid number of clicks, expected 5 got %d", len(clicks))
		}
	})
}

func TestUrlService_GetShortLinks(t *testing.T) {
	db, _ := initTestDB(t)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}
	user := initTestUser(db)
	for i := 0; i < 5; i++ {
		initTestUrl(db, user.ID)
	}

	clicks := urlService.GetShortLinks(user.ID, "test ip address")
	if len(clicks) != 5 {
		t.Errorf("Invalid number of short links, expected 5 got %d", len(clicks))
	}
}

func TestUrlService_GetUrlClicksByUserIdAndUrl(t *testing.T) {
	db, _ := initTestDB(t)
	defer teardownTestDB(db)
	urlService := UrlService{DB: db}
	user := initTestUser(db)
	url := initTestUrl(db, user.ID)
	for i := 0; i < 5; i++ {
		urlService.VisitUrl(url.ShortUrlId, "test user agent", "test ip address", "my host name")
	}

	clicks := urlService.GetUrlClicksByUserIdAndUrl(user.ID, url.ShortUrlId)
	if len(clicks) != 5 {
		t.Errorf("Invalid number of short links, expected 5 got %d", len(clicks))
	}
}
