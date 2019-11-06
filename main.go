package main

import (
	"fmt"
	"goland/api"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
	)

	api := api.NewRouter()

	router.Mount("/v1/api", api)

	return router
}

func main() {
	router := newRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Running server...")
	http.ListenAndServe(":3000", router)
}
