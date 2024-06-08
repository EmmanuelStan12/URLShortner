package middleware

import (
	"context"
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	appcontext "github.com/EmmanuelStan12/URLShortner/internal/context"
	"net/http"
)

func ContextMiddleware(ctx appcontext.Context) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			routeCtx := context.WithValue(r.Context(), constants.AppContextKey, ctx)
			r = r.WithContext(routeCtx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
