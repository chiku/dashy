// Parts of the logging.go are based on
// https://github.com/gorilla/handlers/blob/master/handlers_test.go
// gorilla/handlers is released under the BSD license.

// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func newTimeStamp() time.Time {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	return time.Date(1983, 05, 26, 3, 30, 45, 0, loc)
}

func TestLogEntryOutput(t *testing.T) {
	ts := newTimeStamp()
	req := newRequest("GET", "http://example.com")
	req.RemoteAddr = "192.168.100.5"

	entry := NewLogEntry(req, *req.URL, ts, http.StatusOK, 100)
	expectedStartTime := ts
	if entry.startTime != expectedStartTime {
		t.Errorf("wrong startTime, got %s expected %s", entry.startTime, expectedStartTime)
	}

	expectedClient := "192.168.100.5"
	if entry.client != expectedClient {
		t.Errorf("wrong client, got %s expected %s", entry.client, expectedClient)
	}

	expectedProtocol := "HTTP/1.1"
	if entry.protocol != expectedProtocol {
		t.Errorf("wrong protocol, got %s expected %s", entry.protocol, expectedProtocol)
	}

	expectedMethod := "GET"
	if entry.method != expectedMethod {
		t.Errorf("wrong method, got %s expected %s", entry.method, expectedMethod)
	}

	expectedURI := "/"
	if entry.uri != expectedURI {
		t.Errorf("wrong URI, got %s expected %s", entry.uri, expectedURI)
	}

	expectedStatus := 200
	if entry.status != expectedStatus {
		t.Errorf("wrong HTTP status, got %d expected %d", entry.status, expectedStatus)
	}

	expectedSize := 100
	if entry.size != expectedSize {
		t.Errorf("wrong size, got %d expected %d", entry.size, expectedSize)
	}

	if entry.duration <= 0 {
		t.Errorf("wrong duration, got %d expected greater than or equal to 0", entry.duration)
	}
}

func TestLogEntryOutputForHTTP2Connect(t *testing.T) {
	ts := newTimeStamp()
	req := &http.Request{
		Method:     "CONNECT",
		Host:       "www.example.com:443",
		Proto:      "HTTP/2.0",
		ProtoMajor: 2,
		ProtoMinor: 0,
		RemoteAddr: "192.168.100.5",
		Header:     http.Header{},
		URL:        &url.URL{Host: "www.example.com:443"},
	}

	entry := NewLogEntry(req, *req.URL, ts, http.StatusOK, 100)
	expectedURI := "www.example.com:443"
	if entry.uri != expectedURI {
		t.Errorf("wrong URI, got %s expected %s", entry.uri, expectedURI)
	}
}

func TestLogEntryOutputForIPv6RemoteAddr(t *testing.T) {
	ts := newTimeStamp()
	req := newRequest("GET", "http://example.com")
	req.RemoteAddr = "::1"

	entry := NewLogEntry(req, *req.URL, ts, http.StatusOK, 100)

	expectedClient := "::1"
	if entry.client != expectedClient {
		t.Errorf("wrong client, got %s expected %s", entry.client, expectedClient)
	}

	expectedURI := "/"
	if entry.uri != expectedURI {
		t.Errorf("wrong URI, got %s expected %s", entry.uri, expectedURI)
	}
}

func TestLogEntryOutputForIPv6RemoteAddrWithPort(t *testing.T) {
	ts := newTimeStamp()
	req := newRequest("GET", "http://example.com")
	req.RemoteAddr = net.JoinHostPort("::1", "65000")

	entry := NewLogEntry(req, *req.URL, ts, http.StatusOK, 100)

	expectedClient := "::1"
	if entry.client != expectedClient {
		t.Errorf("wrong client, got %s expected %s", entry.client, expectedClient)
	}

	expectedURI := "/"
	if entry.uri != expectedURI {
		t.Errorf("wrong URI, got %s expected %s", entry.uri, expectedURI)
	}
}

func TestResponseLoggerStatusAndSize(t *testing.T) {
	const hello = "Hello"
	const world = "World!"
	var buf bytes.Buffer

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello")
		fmt.Fprintln(w, "World!")
		w.WriteHeader(205)
	})

	logger := NewLoggingHandler(&buf, handler)

	logger.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/foo/2"))

	expectedSizePart := fmt.Sprintf(`"size": %d`, len(hello)+len(world)+2) // 2 new-lines
	if !strings.Contains(buf.String(), expectedSizePart) {
		t.Fatalf("Got log \n%s\n, wanted substring \n%s", buf.String(), expectedSizePart)
	}

	expectedStatusPart := `"status": 205`
	if !strings.Contains(buf.String(), expectedStatusPart) {
		t.Fatalf("Got log \n%s\n, wanted substring \n%s", buf.String(), expectedStatusPart)
	}
}

func TestLogPathRewrites(t *testing.T) {
	var buf bytes.Buffer

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/" // simulate http.StripPrefix and friends
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
	})
	logger := NewLoggingHandler(&buf, handler)

	logger.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/subdir/asdf"))

	expectedURIPart := `"uri": "/subdir/asdf"`
	if !strings.Contains(buf.String(), expectedURIPart) {
		t.Fatalf("Got log \n%s\n, wanted substring \n%s", buf.String(), expectedURIPart)
	}
}

func BenchmarkLogEntryOutput(b *testing.B) {
	ts := newTimeStamp()

	req := newRequest("GET", "http://example.com")
	req.RemoteAddr = "192.168.100.5"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		entry := NewLogEntry(req, *req.URL, ts, http.StatusUnauthorized, 500).Output()
		io.WriteString(ioutil.Discard, entry)
	}
}
