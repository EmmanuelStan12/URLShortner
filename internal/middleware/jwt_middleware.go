package middleware

import (
	"context"
	"errors"
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/EmmanuelStan12/URLShortner/pkg/jwt"
	"gorm.io/gorm"
	"net/http"
)

func JWTMiddleware(service jwt.JWTService, routes util.Routes, db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if routes.IsPublic(path) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			if token == "" {
				panic(apperrors.UnauthorizedError(errors.New("invalid token")))
			}
			token = token[len("Bearer "):]
			userId, err := service.ParseToken(token)
			if err != nil {
				panic(err)
			}
			user := models.User{}
			res := db.First(&user, userId)
			if res.Error != nil {
				if errors.Is(res.Error, gorm.ErrRecordNotFound) {
					panic(errors.New("user not found"))
				}
				panic(res.Error)
			}
			routeCtx := context.WithValue(r.Context(), constants.UserKey, user)
			r = r.WithContext(routeCtx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
