package app_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request Gocd", func() {
	Context("When succeess", func() {
		It("Converts external JSON to Godashboard", func() {
			const dashboardJSON = `[{ "pipelines": [{"name": "Pipeline"}] }]`

			body := ioutil.NopCloser(bytes.NewBufferString(dashboardJSON))
			response := &http.Response{StatusCode: http.StatusOK, Body: body}
			dashboard, err := a.ParseHTTPResponse(response)

			Expect(dashboard).To(HaveLen(1))
			Expect(dashboard[0].Pipelines).To(HaveLen(1))
			pipeline := dashboard[0].Pipelines[0]
			Expect(pipeline.Name).To(Equal("Pipeline"))
			Expect(err).To(BeNil())
		})
	})

	Context("When HTTP status code is not 200", func() {
		It("Reports the incorrect HTTP status code", func() {
			response := &http.Response{StatusCode: http.StatusInternalServerError}
			dashboard, err := a.ParseHTTPResponse(response)

			Expect(err).To(Equal(fmt.Errorf("the HTTP status code was 500")))
			Expect(dashboard).To(BeNil())
		})
	})

	Context("When HTTP body is nil", func() {
		It("Reports the absence of body", func() {
			response := &http.Response{StatusCode: http.StatusOK, Body: nil}
			dashboard, err := a.ParseHTTPResponse(response)

			Expect(err).To(Equal(fmt.Errorf("the response didn't have a body")))
			Expect(dashboard).To(BeNil())
		})
	})

	Context("When HTTP body read fails", func() {
		It("Reports the error", func() {
			body := &BadReadCloser{err: errors.New("read error")}
			response := &http.Response{StatusCode: http.StatusOK, Body: body}
			dashboard, err := a.ParseHTTPResponse(response)

			Expect(err).To(Equal(fmt.Errorf("error reading response: read error")))
			Expect(dashboard).To(BeNil())
		})
	})

	Context("When HTTP body JSON parse fails", func() {
		It("Reports the error", func() {
			body := ioutil.NopCloser(bytes.NewBufferString("Bad body"))
			response := &http.Response{StatusCode: http.StatusOK, Body: body}
			dashboard, err := a.ParseHTTPResponse(response)

			Expect(err.Error()).To(ContainSubstring("error in unmarshalling external dashboard JSON"))
			Expect(dashboard).To(BeNil())
		})
	})
})
