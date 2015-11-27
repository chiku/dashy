package app_test

import (
	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SimpleDashboard", func() {
	Context("Marshal to JSON", func() {
		It("Has key names of pipelines starting with lower-case", func() {
			simpleDashboard := a.SimpleDashboard{
				Pipelines: []a.SimplePipeline{
					a.SimplePipeline{
						Name: "Pipeline",
						Stages: []a.SimpleStage{
							a.SimpleStage{
								Name:   "Stage",
								Status: "Passed",
							},
						},
					},
				},
			}

			body, err := simpleDashboard.ToJSON()
			Expect(err).To(BeNil())
			Expect(string(body)).To(Equal(`[{"name":"Pipeline","stages":[{"name":"Stage","status":"Passed"}]}]`))
		})
	})
})
