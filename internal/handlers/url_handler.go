package handlers

import (
	"encoding/json"
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	appcontext "github.com/EmmanuelStan12/URLShortner/internal/context"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CreateShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(appcontext.Context)
	urlService := ctx.GetUrlService()
	request := dto.CreateShortUrl{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(apperrors.BadRequestError(err))
	}

	user := r.Context().Value(constants.UserKey).(dto.UserDTO)

	url := urlService.Create(user.ID, request, ctx.Config.Server.Hostname)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(dto.SuccessResponse[dto.UrlDTO]{
		Status: http.StatusOK,
		Data:   url,
	})
}

func GetShortenedLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(appcontext.Context)
	urlService := ctx.GetUrlService()

	user := r.Context().Value(constants.UserKey).(dto.UserDTO)

	urls := urlService.GetShortLinks(user.ID, ctx.Config.Server.Hostname)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(dto.SuccessResponse[[]dto.UrlDTO]{
		Status: http.StatusOK,
		Data:   urls,
	})
}

func VisitShortenedUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(appcontext.Context)
	urlService := ctx.GetUrlService()
	path := chi.URLParam(r, constants.ShortLinkKey)

	url := urlService.VisitUrl(path, r.UserAgent(), util.GetIPAddress(*r), ctx.Config.Server.Hostname)

	http.Redirect(w, r, url.OriginalURL, http.StatusSeeOther)
}

func GetShortenedLinkClicks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(appcontext.Context)
	urlService := ctx.GetUrlService()
	path := chi.URLParam(r, constants.ShortLinkKey)

	user := r.Context().Value(constants.UserKey).(dto.UserDTO)

	urlClicks := urlService.GetUrlClicksByUserIdAndUrl(user.ID, path)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(dto.SuccessResponse[[]dto.UrlClickDTO]{
		Status: http.StatusOK,
		Data:   urlClicks,
	})
}
