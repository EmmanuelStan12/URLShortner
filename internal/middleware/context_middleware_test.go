package middleware

import (
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	"github.com/EmmanuelStan12/URLShortner/internal/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextMiddleware(t *testing.T) {
	ctx, err := context.InitRootContext()
	if err != nil {
		t.Errorf("Failed to init context, %s.", err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeCtx, ok := r.Context().Value(constants.AppContextKey).(context.Context)

		if !ok {
			t.Errorf("Unable to infer context key %s, type.", constants.AppContextKey)
			return
		}

		if routeCtx.Config.DB.DBName != ctx.Config.DB.DBName {
			t.Errorf("Invalid context, expected %v, got %v.", *ctx, routeCtx)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	r := httptest.NewRequest(http.MethodGet, "https://example.com", nil)

	w := httptest.NewRecorder()
	handler := ContextMiddleware(*ctx)(testHandler)

	handler.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Invalid status code, expected %d, got %d.", http.StatusOK, w.Code)
	}

}
