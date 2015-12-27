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
		It("extracts information", func() {
			body := ioutil.NopCloser(bytes.NewBufferString(`{"url": "http://gocd.com:8153", "interests": ["Foo", "Bar"]}`))
			request := &http.Request{Body: body}
			dashy, err := a.NewDashy(request)

			Expect(err).To(BeNil())
			Expect(dashy.URL).To(Equal("http://gocd.com:8153"))
			Expect(dashy.Interests).To(Equal([]string{"Foo", "Bar"}))
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
