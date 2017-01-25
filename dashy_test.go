// dashy_test.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type BadReader struct{ err error }

func (rc BadReader) Read([]byte) (int, error) { return 0, rc.err }

func TestDashy(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBufferString(`{"url": "http://gocd.com:8153", "interests": ["Foo", "Bar:>Bar1"]}`))
	request := &http.Request{Body: body}
	dashy, err := NewDashy(request)

	if err != nil {
		t.Fatalf("Expected no error creating a dashy from HTTP request: %s", err)
	}

	expectedURL := "http://gocd.com:8153"
	if dashy.URL != expectedURL {
		t.Errorf("Expected dashy.URL to equal %s, but was %s", dashy.URL, expectedURL)
	}

	interests := dashy.Interests

	expectedNameList := []string{"Foo", "Bar"}
	if !reflect.DeepEqual(interests.NameList(), expectedNameList) {
		t.Errorf("Expected names to equal %v, but was %v", interests.NameList(), expectedNameList)
	}

	expectedDisplayNames := map[string]string{"Foo": "Foo", "Bar": "Bar1"}
	if !reflect.DeepEqual(interests.DisplayNameMapping(), expectedDisplayNames) {
		t.Errorf("Expected display names when none present to equal %v, but was %v", interests.DisplayNameMapping(), expectedDisplayNames)
	}
}

func TestDashyWhenRequestBodyReadFails(t *testing.T) {
	body := ioutil.NopCloser(BadReader{err: errors.New("read error")})
	request := &http.Request{Body: body}
	dashy, err := NewDashy(request)

	if err == nil {
		t.Fatal("Expected error creating a dashy from HTTP request when body read fails but wasn't")
	}

	expectedErrMsg := "failed to read request body: read error"
	if err.Error() != expectedErrMsg {
		t.Fatalf("Expected error message %q to equal %q, but wasn't", err.Error(), expectedErrMsg)
	}

	if dashy != nil {
		t.Fatalf("Expected no invalid dashy, but was: %v", dashy)
	}
}

func TestDashyWhenRequestBodyJSONParseFails(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBufferString(`BAD JSON`))
	request := &http.Request{Body: body}
	dashy, err := NewDashy(request)

	if err == nil {
		t.Fatal("Expected error creating a dashy from HTTP request when body read fails but wasn't")
	}

	expectedErrPartMsg := "failed to parse JSON: "
	if !strings.Contains(err.Error(), expectedErrPartMsg) {
		t.Fatalf("Expected error message %q to contain sub-string %q, but didn't", err.Error(), expectedErrPartMsg)
	}

	if dashy != nil {
		t.Fatalf("Expected no invalid dashy, but was: %v", dashy)
	}
}
