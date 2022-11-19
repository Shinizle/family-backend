package main

import (
	"github.com/Shinizle/family-backend/Models"
	"github.com/Shinizle/family-backend/Routers/Api"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

func main() {
	Models.ConnectDatabase()
	router := mux.NewRouter()

	mount(router, "/api/customer", Api.CustomerRoute())

	// Use default options
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
