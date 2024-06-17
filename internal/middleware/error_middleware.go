package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"net/http"
)

func ErrorMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					var status int
					var message string

					switch e := err.(type) {
					case *errors.Error:
						status = e.Code
						message = fmt.Sprintf("%s", e.Err)
					case error:
						status = http.StatusInternalServerError
						message = e.Error()
					default:
						status = http.StatusInternalServerError
						message = "An unknown error has occurred."
					}

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(status)
					json.NewEncoder(w).Encode(dto.ErrorResponse{
						Status:  status,
						Message: message,
					})
				}
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
