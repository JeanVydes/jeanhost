package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	api	*mux.Router
	v1 *mux.Router
)

type Map map[string]interface{}

func Start() {
	router = mux.NewRouter()
	api = router.PathPrefix("/api").Subrouter()
	v1 = api.PathPrefix("/v1").Subrouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi darling"))
	})

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working."))
	})

	v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(Map{"message": "working."})

		if err != nil {
			w.Write([]byte("error."))
			return
		}

		w.Write(bytes)
	})
}