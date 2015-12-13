package app_test

import (
	"sort"

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

	Context("Sort by order", func() {
		It("Sorts in ascending order", func() {
			p1 := (a.SimplePipeline{
				Name:   "PipelineOne",
				Stages: []a.SimpleStage{a.SimpleStage{Name: "Stage", Status: "Passed"}},
			}).Order(2)
			p2 := (a.SimplePipeline{
				Name:   "PipelineTwo",
				Stages: []a.SimpleStage{a.SimpleStage{Name: "Stage", Status: "Passed"}},
			}).Order(1)
			pipelines := []a.SimplePipeline{p1, p2}

			sort.Sort(a.ByOrder(pipelines))

			Expect(pipelines).To(HaveLen(2))
			Expect(pipelines[0].Name).To(Equal("PipelineTwo"))
			Expect(pipelines[1].Name).To(Equal("PipelineOne"))
		})
	})
})
