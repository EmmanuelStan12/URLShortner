package middleware

import (
	"errors"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorMiddleware(t *testing.T) {

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{
			name: "App errors",
			handler: func(w http.ResponseWriter, r *http.Request) {
				panic(apperrors.UnauthorizedError("testing app errors"))
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "In-built errors",
			handler: func(w http.ResponseWriter, r *http.Request) {
				panic(errors.New("testing in-go errors"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Unknown errors",
			handler: func(w http.ResponseWriter, r *http.Request) {
				panic("Unknown error handler")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "https://example.com", nil)
			w := httptest.NewRecorder()

			handler := ErrorMiddleware()(tt.handler)

			handler.ServeHTTP(w, r)

			if tt.expectedStatus != w.Code {
				t.Errorf("Invalid status code, expected %d, got %d.", tt.expectedStatus, w.Code)
			}

			t.Logf("Response %v", w.Body)
		})
	}
}
