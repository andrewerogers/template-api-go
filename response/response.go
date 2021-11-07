package response

import (
	"encoding/json"
	"net/http"
)

type GeneralResponse struct {
	Message string
}

func (gr GeneralResponse) Write(w http.ResponseWriter) {
	json, _ := json.Marshal(gr)
	w.WriteHeader(200)
	w.Write(json)
}
