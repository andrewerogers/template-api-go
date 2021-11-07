package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/andrewerogers/template-api-go/handler"
)

func Qos(router *mux.Router) {
	router.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)
	router.HandleFunc("/sping", handler.SPing).Methods(http.MethodGet)
}