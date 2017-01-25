// handler_test.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chiku/gocd"
)

const validRemoteResponse string = `[{
    "name": "Group",
    "pipelines": [{
        "name": "Pipeline",
        "instances": [{
            "stages": [{
                "name": "StageOne",
                "status": "Passed"
            }, {
                "name": "StageTwo",
                "status": "Building"
            }]
          }
        ],
        "previous_instance": {
          "result": "Passed"
        }
      }
    ]
}]`

func TestDashyHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, validRemoteResponse)
	}))
	defer ts.Close()

	logger := log.New(os.Stdout, "test-log: ", 0)
	router := NewRouter(logger)

	body := bytes.NewBufferString(fmt.Sprintf(`{"url": "%s", "interests": ["Pipeline"]}`, ts.URL))
	r, err := http.NewRequest("POST", "/dashy", body)
	if err != nil {
		t.Fatalf("failed to create test HTTP request: %s", err)
	}
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("HTTP status was %d instead of %d, body was: %s", w.Code, http.StatusOK, w.Body.String())
	}

	expectedContentType := "application/json; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Fatalf("Expected HTTP content type to be %q, but was %q", w.Header().Get("Content-Type"), expectedContentType)
	}

	var dashboard gocd.Dashboard
	err = json.Unmarshal(w.Body.Bytes(), &dashboard)
	if err != nil {
		t.Fatalf("Expected response body to be well formatted JSON, but wasn't :%s", err)
	}

	if len(dashboard) != 1 {
		t.Fatalf("Expected dashboard to contain 1 item, but it had %d items: dashboard: %#v", len(dashboard), dashboard)
	}

	pipeline := dashboard[0]
	if pipeline.Name != "Pipeline" {
		t.Errorf("Expected proper pipeline name, but was: %v", pipeline.Name)
	}

	stages := pipeline.Stages
	if len(stages) != 2 {
		t.Fatalf("Expected stages to contain 2 items, but it had %d items: stages: %#v", len(stages), stages)
	}

	stage0 := stages[0]
	if stage0.Name != "StageOne" || stage0.Status != "Passed" {
		t.Errorf("Expected first stage to be proper, but was: %#v", stage0)
	}

	stage1 := stages[1]
	if stage1.Name != "StageTwo" || stage1.Status != "Building" {
		t.Errorf("Expected second stage to be proper, but was: %#v", stage1)
	}
}

func TestDashyHandlerWhenClientError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, validRemoteResponse)
	}))
	defer ts.Close()

	logger := log.New(os.Stdout, "test-log: ", 0)
	router := NewRouter(logger)

	body := bytes.NewBufferString(fmt.Sprintf(`{"url": "%s", "interests": MALFORMED}`, ts.URL))
	r, err := http.NewRequest("POST", "/dashy", body)
	if err != nil {
		t.Fatalf("failed to create test HTTP request: %s", err)
	}
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("HTTP status was %d instead of %d, body was: %s", w.Code, http.StatusBadRequest, w.Body.String())
	}

	expectedContentType := "text/plain; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Fatalf("Expected HTTP content type to be %q, but was %q", w.Header().Get("Content-Type"), expectedContentType)
	}

	expectedBody := "error reading dashy request\n"
	if w.Body.String() != expectedBody {
		t.Fatalf("HTTP body was %q instead of %q", w.Body.String(), expectedBody)
	}
}

func TestDashyHandlerWhenRemoteError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request!", http.StatusInternalServerError)
	}))
	defer ts.Close()

	logger := log.New(os.Stdout, "test-log: ", 0)
	router := NewRouter(logger)

	body := bytes.NewBufferString(fmt.Sprintf(`{"url": "%s", "interests": ["Pipeline"]}`, ts.URL))
	r, err := http.NewRequest("POST", "/dashy", body)
	if err != nil {
		t.Fatalf("failed to create test HTTP request: %s", err)
	}
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("HTTP status was %d instead of %d, body was: %s", w.Code, http.StatusServiceUnavailable, w.Body.String())
	}

	expectedContentType := "text/plain; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Fatalf("Expected HTTP content type to be %q, but was %q", w.Header().Get("Content-Type"), expectedContentType)
	}

	expectedBody := "error fetching data from Gocd\n"
	if w.Body.String() != expectedBody {
		t.Fatalf("HTTP body was %q instead of %q", w.Body.String(), expectedBody)
	}
}

func TestDashyHandlerWhenNoPipelineMatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, validRemoteResponse)
	}))
	defer ts.Close()

	logger := log.New(os.Stdout, "test-log: ", 0)
	router := NewRouter(logger)

	body := bytes.NewBufferString(fmt.Sprintf(`{"url": "%s", "interests": ["SomeThingElse"]}`, ts.URL))
	r, err := http.NewRequest("POST", "/dashy", body)
	if err != nil {
		t.Fatalf("failed to create test HTTP request: %s", err)
	}
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("HTTP status was %d instead of %d, body was: %s", w.Code, http.StatusOK, w.Body.String())
	}

	expectedContentType := "application/json; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Fatalf("Expected HTTP content type to be %q, but was %q", w.Header().Get("Content-Type"), expectedContentType)
	}

	expectedBody := "null"
	if w.Body.String() != expectedBody {
		t.Fatalf("HTTP body was %q instead of %q", w.Body.String(), expectedBody)
	}
}
