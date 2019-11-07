package feeds

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router ... Create feeds router
func Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", allFeeds)
	router.Get("/{feedID}", getFeed)
	return router
}

func allFeeds(w http.ResponseWriter, r *http.Request) {

}

func getFeed(w http.ResponseWriter, r *http.Request) {

}
