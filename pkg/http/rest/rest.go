package rest

import (
	"net/http"

	"github.com/fiseo/httpsrv"
	"github.com/husobee/vestigo"
)

// New instantiates the wing-chun rest server
func New() *http.Server {
	server := httpsrv.NewWithDefault(newRouter())
	return server
}

func newRouter() http.Handler {
	router := vestigo.NewRouter()

	router.Get("/ping", pingHandler())
	router.Put("/ping", pingHandler())
	router.Post("/ping", pingHandler())
	router.Delete("/ping", pingHandler())

	return router
}

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"response": "pong"}`))
	}
}
