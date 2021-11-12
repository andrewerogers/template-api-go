package main

import (
	"net/http"
	"github.com/andrewerogers/template-api-go/route"
	"github.com/gorilla/mux"
)

var routes []route.Route = []route.Route{ 
	route.Qos{}, 
}

func main() {
	router := createRouter()
	http.ListenAndServe(":8080", router)
}

// createRouter registers all service routes 
func createRouter() *mux.Router {
	router := mux.NewRouter()
	for _, r := range routes {
		r.Register(router)
	}
	return router
}