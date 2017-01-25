// interests.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

type Interests []Interest

func (is Interests) DisplayNameMapping() map[string]string {
	mapping := make(map[string]string)
	for _, i := range is {
		mapping[i.Name] = i.DisplayName
	}
	return mapping
}

func (is Interests) NameList() []string {
	var list []string
	for _, i := range is {
		list = append(list, i.Name)
	}
	return list
}
