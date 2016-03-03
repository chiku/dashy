// app/dashy.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type dashy struct {
	URL       string    `json:"url"`
	Interests interests `json:"interests"`
}

type interest struct {
	Name        string
	DisplayName string
}

func (i *interest) UnmarshalJSON(b []byte) (err error) {
	raw := ""
	if err = json.Unmarshal(b, &raw); err != nil {
		return err
	}
	parts := strings.Split(raw, ":>")
	i.Name = parts[0]
	if len(parts) >= 2 && parts[1] != "" {
		i.DisplayName = parts[1]
	}

	return nil
}

type interests []interest

func (is interests) DisplayNameMapping() map[string]string {
	mapping := make(map[string]string)
	for _, i := range is {
		if i.DisplayName != "" {
			mapping[i.Name] = i.DisplayName
		}
	}
	return mapping
}

func (is interests) NameList() []string {
	var list []string
	for _, i := range is {
		list = append(list, i.Name)
	}
	return list
}

func NewDashy(request *http.Request) (*dashy, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %s", err)
	}
	defer request.Body.Close()

	d := &dashy{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %s", err)
	}

	return d, nil
}
