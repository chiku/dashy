// interest_test.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestInterestUnmarshalJSONWithDisplayNameAlias(t *testing.T) {
	input := []byte(`"Foo:>Foo1"`)
	interest := &Interest{}
	err := interest.UnmarshalJSON(input)

	if err != nil {
		t.Fatalf("Expected no error creating interests from valid JSON: %s", err)
	}

	expectedInterest := &Interest{Name: "Foo", DisplayName: "Foo1"}
	if !reflect.DeepEqual(interest, expectedInterest) {
		t.Errorf("Expected interest name to equal the given display name %+v, but was %+v", expectedInterest, interest)
	}
}

func TestInterestUnmarshalJSONWithoutDisplayNameAlias(t *testing.T) {
	input := []byte(`"Foo"`)
	interest := &Interest{}
	err := interest.UnmarshalJSON(input)

	if err != nil {
		t.Fatalf("Expected no error creating interests from valid JSON: %s", err)
	}

	expectedInterest := &Interest{Name: "Foo", DisplayName: "Foo"}
	if !reflect.DeepEqual(interest, expectedInterest) {
		t.Errorf("Expected interest name to equal name %+v, but was %+v", expectedInterest, interest)
	}
}

func TestInterestUnmarshalJSONWithMultipleAliases(t *testing.T) {
	input := []byte(`"Foo:>Foo1:>Foo2"`)
	interest := &Interest{}
	err := interest.UnmarshalJSON(input)

	if err != nil {
		t.Fatalf("Expected no error creating interests from valid JSON: %s", err)
	}

	expectedInterest := &Interest{Name: "Foo", DisplayName: "Foo1"}
	if !reflect.DeepEqual(interest, expectedInterest) {
		t.Errorf("Expected interest name to use the first display name %+v, but was %+v", expectedInterest, interest)
	}
}

func TestInterestUnmarshalJSONWithEmptyAlias(t *testing.T) {
	input := []byte(`"Foo:>"`)
	interest := &Interest{}
	err := interest.UnmarshalJSON(input)

	if err != nil {
		t.Fatalf("Expected no error creating interests from valid JSON: %s", err)
	}

	expectedInterest := &Interest{Name: "Foo", DisplayName: "Foo"}
	if !reflect.DeepEqual(interest, expectedInterest) {
		t.Errorf("Expected interest name to equal use the name %+v, but was %+v", expectedInterest, interest)
	}
}

func TestInterestUnmarshalInvalidJSON(t *testing.T) {
	input := []byte{}
	interest := &Interest{}
	err := interest.UnmarshalJSON(input)

	if err == nil {
		t.Fatal("Expected error creating interests from empty JSON")
	}

	expectedErrMsgPart := ""
	if !strings.Contains(err.Error(), expectedErrMsgPart) {
		t.Errorf(`Expected error "%s" to contain "%s", but didn't`, err.Error(), expectedErrMsgPart)
	}
}
