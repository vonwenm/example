package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// The Api Resource response for [GET] /api
type ApiResp struct {
	Version int
	Welcome string
}

//
// Test [GET] /api
//
func TestApi(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Error(err)
	}

	route.ServeHTTP(w, req)

	var resp ApiResp
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}

	if resp.Version != 1 && resp.Welcome != "This is the REST API for a book store" {
		t.Error("[GET] /api answare is something different!")
	}
}