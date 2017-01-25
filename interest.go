// interest.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"encoding/json"
	"strings"
)

type Interest struct {
	Name        string
	DisplayName string
}

func (i *Interest) UnmarshalJSON(b []byte) (err error) {
	raw := ""
	if err = json.Unmarshal(b, &raw); err != nil {
		return err
	}
	parts := strings.Split(raw, ":>")
	i.Name = parts[0]
	if len(parts) >= 2 && parts[1] != "" {
		i.DisplayName = parts[1]
	} else {
		i.DisplayName = i.Name
	}

	return nil
}
