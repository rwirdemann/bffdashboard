package main

import (
	"github.com/gorilla/mux"
	"github.com/rwirdemann/bffdashboard/marketplace/usecase"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", usecase.LoginHandler).Methods("POST")
	router.HandleFunc("/seller/{id}", usecase.SellerHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
