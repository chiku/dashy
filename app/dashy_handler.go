// app/dashy_handler.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/chiku/gocd"
)

func DashyHandler(client *gocd.Client) http.HandlerFunc {
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

		dashboard, err := client.Fetch(url)
		if err != nil {
			errorMsg := "error fetching data from Gocd"
			log.Printf("%s: %s", errorMsg, err)
			http.Error(w, errorMsg, http.StatusServiceUnavailable)
			return
		}

		dashboard, ignores := dashboard.FilteredSort(interests.NameList())
		dashboard = dashboard.MapNames(interests.DisplayNameMapping())
		if len(dashboard) == 0 {
			log.Printf("not configured to display any pipelines, you could try to include some of these pipelines: %s", strings.Join(ignores, ", "))
		}

		output, err := json.Marshal(dashboard)
		if err != nil {
			errorMsg := "error marshalling simple dashboard JSON"
			log.Printf("%s: %s", errorMsg, err)
			http.Error(w, errorMsg, http.StatusServiceUnavailable)
			return
		}

		w.Write(output)
	}
}
