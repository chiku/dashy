// interests_test.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"reflect"
	"testing"
)

func TestInterestsDisplayName(t *testing.T) {
	interests := Interests{
		Interest{Name: "Foo", DisplayName: "Foo1"},
		Interest{Name: "Bar", DisplayName: "Bar1"},
	}

	expectedDisplayNames := map[string]string{"Foo": "Foo1", "Bar": "Bar1"}
	if !reflect.DeepEqual(interests.DisplayNameMapping(), expectedDisplayNames) {
		t.Errorf("Expected display names when none present to equal %v, but was %v", interests.DisplayNameMapping(), expectedDisplayNames)
	}
}

func TestInterestsNameList(t *testing.T) {
	interests := Interests{
		Interest{Name: "Foo", DisplayName: "Foo1"},
		Interest{Name: "Bar", DisplayName: "Bar1"},
	}

	expectedNameList := []string{"Foo", "Bar"}
	if !reflect.DeepEqual(interests.NameList(), expectedNameList) {
		t.Errorf("Expected names to equal %v, but was %v", interests.NameList(), expectedNameList)
	}
}
