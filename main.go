// main.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chiku/dashy/app"
	"github.com/gorilla/handlers"
)

func main() {
	file, err := os.OpenFile("dashy.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()
	logWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logWriter)

	mux := http.DefaultServeMux
	mux.HandleFunc("/dashy", app.DashyHandler)
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	loggingHandler := handlers.CombinedLoggingHandler(logWriter, mux)
	server := &http.Server{
		Addr:    ":3000",
		Handler: loggingHandler,
	}

	fmt.Println("Starting the application on http://localhost:3000")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to start application: %s\n", err)
	}
}
