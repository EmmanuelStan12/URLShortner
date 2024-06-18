package middleware

import (
	"context"
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	"github.com/EmmanuelStan12/URLShortner/internal/services"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/EmmanuelStan12/URLShortner/pkg/jwt"
	"net/http"
	"strings"
)

func JWTMiddleware(service jwt.JWTService, routes util.Routes, userService services.IUserService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			path := strings.Split(r.URL.Path, "/")
			if routes.IsPublic(path[len(path)-1]) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			if token == "" {
				panic(apperrors.UnauthorizedError("invalid token"))
			}
			token = token[len("Bearer "):]
			userId, err := service.ParseToken(token)
			if err != nil {
				panic(apperrors.UnauthorizedError("invalid token"))
			}
			user := userService.GetById(userId)
			routeCtx := context.WithValue(r.Context(), constants.UserKey, user)
			r = r.WithContext(routeCtx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
