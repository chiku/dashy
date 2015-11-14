package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chiku/dashy/app"
	"github.com/gorilla/handlers"
)

// AJAX Request
type dashy struct {
	URL       string   `json:"url"`
	Interests []string `json:"interests"`
}

func dashyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "error reading dashy request"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	dashy := &dashy{}
	if err = json.Unmarshal(body, dashy); err != nil {
		errorMsg := "error unmarshalling dashy JSON"
		log.Printf("%s: %s (%s)", errorMsg, err, string(body))
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	goPipelineGroups, err := app.ParseHTTPResponse(http.Get(dashy.URL))
	if err != nil {
		errorMsg := "error fetching data from Gocd"
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
	mux := http.DefaultServeMux
	mux.HandleFunc("/dashy", dashyHandler)
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, mux)
	server := &http.Server{
		Addr:    ":3000",
		Handler: loggingHandler,
	}
	fmt.Println("Starting the application on http://localhost:3000")
	server.ListenAndServe()
}
