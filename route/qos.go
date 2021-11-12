package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/andrewerogers/template-api-go/handler"
)

type Route interface { 
	Register(router *mux.Router)
}

type Qos struct {}

func (q Qos) Register(router *mux.Router) {
	router.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)
	router.HandleFunc("/sping", handler.SPing).Methods(http.MethodGet)
}