package httpserver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	api	*mux.Router
	v1 *mux.Router
)

type Map map[string]interface{}

func Start() {
	Setup()
	Listen()
}

func Setup() {
	router = mux.NewRouter()
	api = router.PathPrefix("/api").Subrouter()
	v1 = api.PathPrefix("/v1").Subrouter()

	setHandlers()
}

func Listen() {
	srv := &http.Server{
        Handler:      router,
        Addr:         "0.0.0.0",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}