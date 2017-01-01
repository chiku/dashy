// app/dashy_test.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package app_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dashy", func() {
	Context("when accepting a HTTP request without display name", func() {
		body := ioutil.NopCloser(bytes.NewBufferString(`{"url": "http://gocd.com:8153", "interests": ["Foo", "Bar"]}`))
		request := &http.Request{Body: body}
		dashy, err := a.NewDashy(request)

		It("doesn't have an error", func() {
			Expect(err).To(BeNil())
		})

		It("extracts URL", func() {
			Expect(dashy.URL).To(Equal("http://gocd.com:8153"))
		})

		It("extracts interests with names", func() {
			interests := dashy.Interests
			Expect(interests.NameList()).To(Equal([]string{"Foo", "Bar"}))
			Expect(interests.DisplayNameMapping()).To(BeEmpty())
		})
	})

	Context("when accepting a HTTP request with display name", func() {
		body := ioutil.NopCloser(bytes.NewBufferString(`{"url": "http://gocd.com:8153", "interests": ["Foo:>Foo1", "Bar:>Bar1"]}`))
		request := &http.Request{Body: body}
		dashy, err := a.NewDashy(request)

		It("doesn't have an error", func() {
			Expect(err).To(BeNil())
		})

		It("extracts URL", func() {
			Expect(dashy.URL).To(Equal("http://gocd.com:8153"))
		})

		It("extracts interests with names and display names", func() {
			interests := dashy.Interests
			Expect(interests.NameList()).To(Equal([]string{"Foo", "Bar"}))
			Expect(interests.DisplayNameMapping()).To(Equal(map[string]string{"Foo": "Foo1", "Bar": "Bar1"}))
		})
	})

	Context("when reading body fails", func() {
		body := &BadReadCloser{err: errors.New("read error")}
		request := &http.Request{Body: body}
		dashy, err := a.NewDashy(request)

		It("has an error", func() {
			Expect(err.Error()).To(Equal("failed to read request body: read error"))
			Expect(dashy).To(BeNil())
		})
	})

	Context("when JSON parse fails", func() {
		It("has an error", func() {
			body := ioutil.NopCloser(bytes.NewBufferString(`BAD JSON`))
			request := &http.Request{Body: body}
			dashy, err := a.NewDashy(request)

			Expect(err.Error()).To(ContainSubstring("failed to parse JSON: "))
			Expect(dashy).To(BeNil())
		})
	})
})
