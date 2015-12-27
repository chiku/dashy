package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dashyRequest struct {
	URL       string   `json:"url"`
	Interests []string `json:"interests"`
}

type dashy struct {
	URL       string
	Interests *Interests
}

func NewDashy(request *http.Request) (*dashy, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %s", err)
	}

	defer request.Body.Close()

	d := &dashyRequest{}
	err = json.Unmarshal(body, d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %s", err)
	}

	interests := NewInterests()
	for _, interest := range d.Interests {
		interests.Add(interest)
	}

	return &dashy{
		URL:       d.URL,
		Interests: interests,
	}, nil
}
