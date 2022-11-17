package main

import (
	"github.com/Shinizle/family-backend/Models"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Models.ConnectDatabase()
	r := mux.NewRouter()

	http.ListenAndServe(":8080", r)
}
