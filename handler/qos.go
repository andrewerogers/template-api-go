package handler

import (
	"net/http"
	"github.com/andrewerogers/template-api-go/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	gr := response.GeneralResponse{
		Message: "healthy",
	}
	gr.Write(w)
}

func SPing(w http.ResponseWriter, r *http.Request) {
	gr := response.GeneralResponse{
		Message: "healthy",
	}
	gr.Write(w)
}