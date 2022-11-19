package Api

import (
	"github.com/Shinizle/family-backend/Controllers/NationalityController"
	"github.com/gorilla/mux"
)

func NationalityRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/get-all", NationalityController.Index).Methods("GET")
	router.HandleFunc("/show/{id}", NationalityController.Show).Methods("GET")
	router.HandleFunc("/create", NationalityController.Create).Methods("POST")
	router.HandleFunc("/update/{id}", NationalityController.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", NationalityController.Delete).Methods("DELETE")

	return router
}
