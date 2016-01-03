// main.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chiku/dashy/app"
	"github.com/gorilla/handlers"
)

const maxRetries = 3

func fetchWithRetries(url string) (response *http.Response, err error) {
	retries := 0

	for response == nil && retries < maxRetries {
		response, err = http.Get(url)
		if err != nil {
			log.Printf("error fetching data from Gocd (retry #%d): %s", retries+1, err)
		}
		retries++
	}
	return
}

func dashyHandler(w http.ResponseWriter, r *http.Request) {
	dashy, err := app.NewDashy(r)
	if err != nil {
		errorMsg := "error reading dashy request"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	response, err := fetchWithRetries(dashy.URL)
	if err != nil {
		errorMsg := "error fetching data from Gocd"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusServiceUnavailable)
		return
	}

	goPipelineGroups, err := app.ParseHTTPResponse(response)
	if err != nil {
		errorMsg := "error parsing Gocd response"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusServiceUnavailable)
		return
	}

	goDashboard := app.GoDashboard{
		PipelineGroups: goPipelineGroups,
		Interests:      dashy.Interests,
	}
	simpleDashboard := goDashboard.ToSimpleDashboard()
	if len(simpleDashboard.Pipelines) == 0 {
		log.Printf("not configured to display any pipelines, you could try to include some of these pipelines: %s", strings.Join(simpleDashboard.Ignores, ", "))
	}

	output, err := json.Marshal(simpleDashboard.Pipelines)
	if err != nil {
		errorMsg := "error marshalling simple dashboard JSON"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusServiceUnavailable)
		return
	}

	w.Write(output)
}

func main() {
	file, err := os.OpenFile("dashy.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()
	logWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logWriter)

	mux := http.DefaultServeMux
	mux.HandleFunc("/dashy", dashyHandler)
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
