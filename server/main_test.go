package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/rinosukmandityo/google-search-ext/server/api"
	"github.com/rinosukmandityo/google-search-ext/server/caller"
)

const (
	ContentTypeJson = "application/json"
)

var (
	r http.Handler
)

func init() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/search", api.PostSearch)

	r = mux
	r = api.CorsMiddleware(r)
}

func TestAuthenticateUserSuccess(t *testing.T) {
	testdata := []api.Search{{
		Title:  "Title 1",
		Desc:   "Desc 1",
		URL:    "www.url1.com/url",
		Domain: "www.url1.com",
		Cite:   "Cite 1",
	}, {
		Title:  "Title 2",
		Desc:   "Desc 2",
		URL:    "www.url2.com/url",
		Domain: "www.url2.com",
		Cite:   "Cite 2",
	}}

	dataBytes, err := getBytes(testdata)
	if err != nil {
		t.Errorf("cannot get data byte: %s", err.Error())
	}

	req, e := http.NewRequest("POST", "/search", bytes.NewReader(dataBytes))
	if e != nil {
		t.Errorf("failed to create a mew request: %s ", e.Error())
	}

	resp, respBody, e := caller.New(r).SetRequest(req).SetResponse(&[]api.Search{}).
		SetHeader("Content-Type", ContentTypeJson).Exec()
	if e != nil {
		t.Errorf("failed to call an API: %s ", e.Error())
	}

	if resp.Code != http.StatusOK {
		t.Errorf("status response is not correct, want %d, got %d", http.StatusOK, resp.Code)
	}

	res := respBody.(*[]api.Search)
	if !reflect.DeepEqual(testdata, *res) {
		t.Errorf("response data is not equal")
	}
}

func getBytes(data []api.Search) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
