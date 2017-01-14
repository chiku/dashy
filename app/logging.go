// Parts of the logging.go are based on
// https://github.com/gorilla/handlers/blob/master/handlers.go
// gorilla/handlers is released under the BSD license.

// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// LoggingHandler wraps a http.Handler implementation with logging support.
type LoggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

// NewLoggingHandler return a http.Handler that wraps h and logs requests to out in
// JSON log Format.
func NewLoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return LoggingHandler{out, h}
}

func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	logger := NewResponseLogger(w)
	url := *req.URL
	h.handler.ServeHTTP(logger, req)
	logEntry := NewLogEntry(req, url, startTime, logger.Status(), logger.Size())
	io.WriteString(h.writer, logEntry.Output())
}

// ResponseLogger is wrapper of http.ResponseWriter that keeps track of its HTTP
// status code and body size.
type ResponseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

// NewResponseLogger returns an instance of ResponseLogger.
func NewResponseLogger(w http.ResponseWriter) *ResponseLogger {
	return &ResponseLogger{w: w}
}

func (l *ResponseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *ResponseLogger) Write(b []byte) (int, error) {
	if l.status == 0 {
		// The status will be StatusOK if WriteHeader has not been called yet.
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *ResponseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *ResponseLogger) Status() int {
	return l.status
}

func (l *ResponseLogger) Size() int {
	return l.size
}

type LogEntry struct {
	startTime time.Time
	client    string
	protocol  string
	method    string
	uri       string
	status    int
	duration  time.Duration
	size      int
}

// NewLogEntry returns a log entry for request.
// startTime is the timestamp with which the entry is logged.
// status and size are the response HTTP status and size respectively.
func NewLogEntry(req *http.Request, url url.URL, startTime time.Time, status, size int) LogEntry {
	host, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		host = req.RemoteAddr
	}

	uri := req.RequestURI

	// Requests using the CONNECT method over HTTP/2.0 must use
	// the authority field (aka r.Host) to identify the target.
	// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
	if req.ProtoMajor == 2 && req.Method == "CONNECT" {
		uri = req.Host
	}
	if uri == "" {
		uri = url.RequestURI()
	}

	return LogEntry{
		startTime: startTime,
		client:    host,
		protocol:  req.Proto,
		method:    req.Method,
		uri:       uri,
		status:    status,
		size:      size,
		duration:  time.Since(startTime) / time.Millisecond,
	}
}

func (entry LogEntry) Output() string {
	return fmt.Sprintf(`
{"time": %q, "client": %q, "protocol": %q, "method": %q, "uri": %q, "status": %d, "duration": %d, "size": %d}`,
		entry.startTime.Format("2006-01-02 15:04:05.999 Z07:00"),
		entry.client, entry.protocol, entry.method, entry.uri, entry.status, entry.duration, entry.size)
}
