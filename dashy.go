// dashy.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dashy struct {
	URL       string    `json:"url"`
	Interests Interests `json:"interests"`
}

func NewDashy(request *http.Request) (*Dashy, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %s", err)
	}
	defer request.Body.Close()

	d := &Dashy{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %s", err)
	}

	return d, nil
}
