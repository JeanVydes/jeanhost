package httpserver

import (
	"net/http"
)

func setHandlers() {
	router.HandleFunc("/", routerHandler)
	
	router.HandleFunc("/jean", jeanHandler)
	router.HandleFunc("/chat", chatHandler)
	router.HandleFunc("/short", shortenerHandler)
	router.HandleFunc("/shortener", shortenerHandler)

	api.HandleFunc("/", apiHandler)
	v1.HandleFunc("/", v1Handler)
}

func routerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"services\": \"https://chat.jean.host https://short.jean.host \",\"domains\": \"jean.host jeanvides.me jeanservices.co cagaca.ga vides.tech jeann.studio elemental.ga\",\"contact\": \"contact@jean.host\"}"))
}

func jeanHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://jeanvides.me", http.StatusMovedPermanently)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://chat.jean.host", http.StatusMovedPermanently)
}

func shortenerHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://short.jean.host", http.StatusMovedPermanently)
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
