package main

import (
	"github.com/gorilla/mux"
	"github.com/andrewerogers/template-api-go/route"
	"net/http"
)

func main() {
	router := mux.NewRouter() // Uses mux router
	route.Qos(router)
	http.ListenAndServe(":8080", router)
}