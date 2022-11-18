package main

import (
	"github.com/Shinizle/family-backend/Models"
	"github.com/Shinizle/family-backend/Routers/Api"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func main() {
	Models.ConnectDatabase()
	router := mux.NewRouter()

	mount(router, "/customer", Api.CustomerRoute())

	http.ListenAndServe(":8080", router)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
