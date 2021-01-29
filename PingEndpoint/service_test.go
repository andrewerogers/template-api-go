package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPing1(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/test/ping")
	if err != nil {
		// Don't want to write any error in this case
	}

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Don't want to write any error in this case
	} else if string(body) != `{"Data":"pong"}` {
		t.Error("Wrong response for ping, got " + string(body) + "expected {\"Data\":\"pong\"}")
	}
}

func TestGetPing2(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/test/123")
	if err != nil {
		// Don't want to write any error in this case
	}

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Don't want to write any error in this case
	} else if string(body) == `{"Data":"pong"}` {
		t.Error("Wrong response for ping, got " + string(body) + "expected ")
	}
}
