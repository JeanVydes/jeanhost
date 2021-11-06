package httpserver

import (
	"net/http"
)

func setHandlers() {
	router.HandleFunc("/", routerHandler)
	api.HandleFunc("/", apiHandler)
	v1.HandleFunc("/", v1Handler)
}

func routerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"working.\"}"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"working.\"}"))
}

func v1Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"working.\"}"))
}
