// app/dashy_handler.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/chiku/gocd"
)

func DashyHandler() http.HandlerFunc {
	fetcher := gocd.Fetch()

	return func(w http.ResponseWriter, r *http.Request) {
		dashy, err := NewDashy(r)
		if err != nil {
			errorMsg := "error reading dashy request"
			log.Printf("%s: %s", errorMsg, err)
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		url := dashy.URL
		interests := dashy.Interests

		output, ignores, err := fetcher(url, interests.NameList(), interests.DisplayNameMapping())
		if err != nil {
			errorMsg := "error fetching data from Gocd"
			log.Printf("%s: %s", errorMsg, err)
			http.Error(w, errorMsg, http.StatusServiceUnavailable)
			return
		}

		if len(output) == 0 {
			log.Printf("not configured to display any pipelines, you could try to include some of these pipelines: %s", strings.Join(ignores, ", "))
		}

		w.Write(output)
	}
}
