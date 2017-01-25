// main.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("dashy.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()

	logWriter := io.MultiWriter(file, os.Stdout)
	logger := log.New(logWriter, "", log.LstdFlags)

	mux := NewRouter(logger)
	loggingHandler := NewLoggingHandler(logWriter, mux)
	server := &http.Server{
		Addr:    ":3000",
		Handler: loggingHandler,
	}

	logger.Println("Starting the application on http://localhost:3000")
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Failed to start application: %s", err)
	}
}
