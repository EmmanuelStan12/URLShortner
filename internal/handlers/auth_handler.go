package handlers

import (
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	"github.com/EmmanuelStan12/URLShortner/internal/context"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(context.Context)

}
