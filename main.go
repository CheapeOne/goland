package main

import (
	"fmt"
	"net/http"

	feeds "github.com/cheapeone/goland/api/feeds/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1/api", func(r chi.Router) {
		r.Mount("/feeds", feeds.Router())
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}

func main() {
	appRouter := router()

	fmt.Println("Running server...")
	http.ListenAndServe(":3000", appRouter)
}
