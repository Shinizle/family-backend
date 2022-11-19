package Api

import (
	"github.com/Shinizle/family-backend/Controllers/CustomerController"
	"github.com/gorilla/mux"
)

func CustomerRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/get-all", CustomerController.Index).Methods("GET")
	router.HandleFunc("/show/{id}", CustomerController.Show).Methods("GET")
	router.HandleFunc("/create", CustomerController.Create).Methods("POST")
	router.HandleFunc("/update/{id}", CustomerController.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", CustomerController.Delete).Methods("DELETE", "OPTIONS")

	return router
}
