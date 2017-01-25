// dashy_handler.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/chiku/gocd"
)

func DashyHandler(logger *log.Logger) http.HandlerFunc {
	fetcher := gocd.Fetch()

	return func(w http.ResponseWriter, r *http.Request) {
		dashy, err := NewDashy(r)
		if err != nil {
			errorMsg := "error reading dashy request"
			logger.Printf("%s: %s", errorMsg, err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		url := dashy.URL
		interests := dashy.Interests

		output, ignores, err := fetcher(url, interests.NameList(), interests.DisplayNameMapping())
		if err != nil {
			errorMsg := "error fetching data from Gocd"
			logger.Printf("%s: %s", errorMsg, err)
			http.Error(w, errorMsg, http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if string(output) == "null" {
			logger.Printf("not configured to display any pipelines, you could try to include some of these pipelines: %s", strings.Join(ignores, ", "))
		}

		w.Write(output)
	}
}
