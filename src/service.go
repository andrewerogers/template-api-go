package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type result struct {
	Data string
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World! GoLang is pretty fun so far!"))
}

func getPing(w http.ResponseWriter, r *http.Request) {
	var param string = mux.Vars(r)["key"]

	if param == "ping" {

		data := result{ // Create an instance of result struct
			Data: "pong",
		}
		resp, err := json.Marshal(data)
		if err != nil {
			fmt.Println("An error occurred marshaling json w/ key " + param)
		} else { // Return the new marshaled json object
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, string(resp))
		}

	} else {
		w.WriteHeader(http.StatusOK) // Still want to set status 200 s
	}
}

func main() {
	router := mux.NewRouter() // Create gorilla mux router
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/ping/{key}", getPing).Methods("GET") // Pass key into test endpoint
	http.ListenAndServe(":8080", router)
}
