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
	Context("when accepting a HTTP request", func() {
		body := ioutil.NopCloser(bytes.NewBufferString(`{"url": "http://gocd.com:8153", "interests": ["Foo", "Bar"]}`))
		request := &http.Request{Body: body}
		dashy, err := a.NewDashy(request)

		It("doesn't have an error", func() {
			Expect(err).To(BeNil())
		})

		It("extracts URL", func() {
			Expect(dashy.URL).To(Equal("http://gocd.com:8153"))
		})

		It("extracts interests by name", func() {
			interests := dashy.Interests
			position, displayName := interests.PipelineName("Bar")
			Expect(position).To(Equal(1))
			Expect(displayName).To(Equal("Bar"))
		})

		It("doesn't extract not-existing interest name", func() {
			interests := dashy.Interests
			position, displayName := interests.PipelineName("NotExisting")
			Expect(position).To(Equal(-1))
			Expect(displayName).To(BeEmpty())
		})

		Context("when reading body fails", func() {
			It("has an error", func() {
				body := &BadReadCloser{err: errors.New("read error")}
				request := &http.Request{Body: body}
				dashy, err := a.NewDashy(request)

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
})
