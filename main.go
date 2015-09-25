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

func toDashboard(dashy *dashy) ([]app.GoPipelineGroup, error) {
	response, err := http.Get(dashy.URL)
	if response == nil {
		return nil, fmt.Errorf("no response from external dashboard")
	}

	if err != nil {
		return nil, fmt.Errorf("bad response external dashboard: %d", response.StatusCode)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response code %d for external dashboard", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading external dashboard response body: %s", err.Error())
	}
	defer response.Body.Close()

	dashboard, err := app.GoPipelineGroupsFromJSON(body)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

func dashyHandler(w http.ResponseWriter, r *http.Request) {
	dashy := &dashy{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "error reading dashy request"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(body, dashy); err != nil {
		errorMsg := "error unmarshalling dashy JSON"
		log.Printf("%s: %s", errorMsg, err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	goPipelineGroups, err := toDashboard(dashy)
	if err != nil {
		errorMsg := "error fetching dashboard"
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
		log.Printf("not configued to display any pipeline, however these pipelines are not included: %s", strings.Join(simpleDashboard.Ignores, ", "))
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
	server.ListenAndServe()
}
