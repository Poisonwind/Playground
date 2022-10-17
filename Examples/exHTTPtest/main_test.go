package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// код писать тут

var answ = "dummy data"

func SearchServer(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(answ))

}

func TestServer(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	defer server.Close()

	result := Ping(server.URL)

	if result != answ {
		t.Errorf("expexted: %s, got: %s", answ, result)
	}

}
