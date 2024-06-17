package app

import (
	appcontext "github.com/EmmanuelStan12/URLShortner/internal/context"
	appmiddleware "github.com/EmmanuelStan12/URLShortner/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Run() {
	ctx, err := appcontext.InitContext()

	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(appmiddleware.ErrorMiddleware())
	r.Use(appmiddleware.ContextMiddleware(*ctx))
	r.Use(appmiddleware.JWTMiddleware(ctx.JWTService, ctx.Routes, ctx.GetUserService()))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":"+ctx.Config.Server.Port, r))
}
