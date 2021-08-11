package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	log.Printf("Chi HTTP server running on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, chiDispatcher))
}
