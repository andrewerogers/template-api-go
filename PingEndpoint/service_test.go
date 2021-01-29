package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetPing1(t *testing.T) {
	req, err := http.NewRequest("GET", "/test/ping", nil)
	if err != nil {
		// Don't want to write any error in this case
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getPing)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Error("Handler returned status code " + strconv.Itoa(responseRecorder.Code))
	} else if responseRecorder.Body.String() != "{\"Data\":\"pong\"}" {
		t.Error("Wrong response returned for key: ping, response got " + responseRecorder.Body.String() + " expected pong")
	}
}
