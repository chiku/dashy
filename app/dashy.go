package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dashy struct {
	URL       string   `json:"url"`
	Interests []string `json:"interests"`
}

func NewDashy(request *http.Request) (*dashy, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %s", err)
	}

	defer request.Body.Close()

	d := &dashy{}
	err = json.Unmarshal(body, d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %s", err)
	}

	return d, nil
}
