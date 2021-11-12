package response

import (
	"encoding/json"
	"net/http"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

func (gr GeneralResponse) Write(w http.ResponseWriter) {
	json, _ := json.Marshal(gr)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (er ErrorResponse) Write(w http.ResponseWriter) {
	json, _ := json.Marshal(er)
	w.WriteHeader(er.Status)
	w.Write(json)
}
