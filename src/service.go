package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func sping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func main() {
	router := mux.NewRouter() // Uses mux router
	router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	router.HandleFunc("/sping", sping).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
