package main

import (
	"github.com/gorilla/mux"
	"github.com/andrewerogers/template-api-go/response"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	gr := response.GeneralResponse {
		Message: "healthy",
	}
	gr.Write(w)
}

func sping(w http.ResponseWriter, r *http.Request) {
	gr := response.GeneralResponse {
		Message: "healthy",
	}
	gr.Write(w)
}

func main() {
	router := mux.NewRouter() // Uses mux router
	router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	router.HandleFunc("/sping", sping).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}