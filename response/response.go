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
