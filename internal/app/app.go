package app

import (
	appcontext "github.com/EmmanuelStan12/URLShortner/internal/context"
	"github.com/EmmanuelStan12/URLShortner/internal/handlers"
	appmiddleware "github.com/EmmanuelStan12/URLShortner/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Run() {
	ctx, err := appcontext.InitRootContext()

	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(appmiddleware.ErrorMiddleware())
	r.Use(appmiddleware.ContextMiddleware(*ctx))
	//r.Use(appmiddleware.JWTMiddleware(ctx.JWTService, ctx.Routes, ctx.GetUserService()))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", handlers.LoginHandler)
		r.Post("/register", handlers.RegisterHandler)
	})

	log.Fatal(http.ListenAndServe(":"+ctx.Config.Server.Port, r))
}
