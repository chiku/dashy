package app_test

import (
	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoDashboard", func() {
	Context("without pipeline-groups", func() {
		goDashboard := a.GoDashboard{}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has no simple-pipelines", func() {
			Expect(simpleDashboard.Pipelines).To(BeEmpty())
		})
	})

	Context("with pipeline-group without pipelines", func() {
		goPipelineGroups := []a.GoPipelineGroup{}
		interests := a.NewInterests().Add("Pipeline")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has no simple-pipelines", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(BeEmpty())
		})
	})

	Context("with pipeline-group, pipeline without instances", func() {
		goInstances := []a.GoInstance{}
		goPipelines := []a.GoPipeline{{Instances: goInstances}}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has no simple-pipelines", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(BeEmpty())
		})
	})

	Context("with pipeline-group, pipeline, instance without stages", func() {
		goStages := []a.GoStage{}
		goInstances := []a.GoInstance{{Stages: goStages}}
		goPipelines := []a.GoPipeline{{Instances: goInstances}}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has no simple-pipelines", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(BeEmpty())
		})
	})

	Context("with pipeline-group, pipeline, instance and stage", func() {
		goStages := []a.GoStage{{Name: "Stage One", Status: "Unknown"}}
		goInstances := []a.GoInstance{{Stages: goStages}}
		goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances}}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline One")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has a simple-pipeline", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(HaveLen(1))
			Expect(pipelines[0].Name).To(Equal("Pipeline One"))
			stages := pipelines[0].Stages
			Expect(stages).To(HaveLen(1))
			Expect(stages[0].Name).To(Equal("Stage One"))
			Expect(stages[0].Status).To(Equal("Unknown"))
		})

		It("has no ignores", func() {
			Expect(simpleDashboard.Ignores).To(BeEmpty())
		})
	})

	Context("with pipeline-group, pipeline, multiple instances and stage", func() {
		goStagesForOldInstance := []a.GoStage{{Name: "Stage Old", Status: "Failed"}}
		goStagesForNewInstance := []a.GoStage{{Name: "Stage New", Status: "Passed"}}
		goOldInstance := a.GoInstance{Stages: goStagesForOldInstance}
		goNewInstance := a.GoInstance{Stages: goStagesForNewInstance}
		goInstances := []a.GoInstance{goOldInstance, goNewInstance}
		goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances}}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline One")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("ignores older instances", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(HaveLen(1))
			stages := pipelines[0].Stages
			Expect(stages).To(HaveLen(1))
			Expect(stages[0].Name).To(Equal("Stage New"))
			Expect(stages[0].Status).To(Equal("Passed"))
		})

		Context("with the current status as unknown", func() {
			goStagesForLatestInstance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
			goStagesForMinus1Instance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
			goStagesForMinus2Instance := []a.GoStage{{Name: "Stage X", Status: "Passed"}}
			goLatestInstance := a.GoInstance{Stages: goStagesForLatestInstance}
			goMinus1Instance := a.GoInstance{Stages: goStagesForMinus1Instance}
			goMinus2Instance := a.GoInstance{Stages: goStagesForMinus2Instance}
			goInstances := []a.GoInstance{goMinus2Instance, goMinus1Instance, goLatestInstance}
			goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances}}
			goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
			interests := a.NewInterests().Add("Pipeline One")
			goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
			simpleDashboard := goDashboard.ToSimpleDashboard()

			It("uses the status of the older build", func() {
				pipelines := simpleDashboard.Pipelines
				Expect(pipelines).To(HaveLen(1))
				stages := pipelines[0].Stages
				Expect(stages).To(HaveLen(1))
				Expect(stages[0].Name).To(Equal("Stage X"))
				Expect(stages[0].Status).To(Equal("Passed"))
			})
		})

		Context("with current and older statuses as unknown", func() {
			goStagesForLatestInstance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
			goStagesForMinus1Instance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
			goLatestInstance := a.GoInstance{Stages: goStagesForLatestInstance}
			goMinus1Instance := a.GoInstance{Stages: goStagesForMinus1Instance}
			goInstances := []a.GoInstance{goMinus1Instance, goLatestInstance}
			goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances}}
			goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
			interests := a.NewInterests().Add("Pipeline One")
			goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
			simpleDashboard := goDashboard.ToSimpleDashboard()

			It("has unknown status", func() {
				pipelines := simpleDashboard.Pipelines
				Expect(pipelines).To(HaveLen(1))
				stages := pipelines[0].Stages
				Expect(stages).To(HaveLen(1))
				Expect(stages[0].Name).To(Equal("Stage X"))
				Expect(stages[0].Status).To(Equal("Unknown"))
			})

			Context("with known previous instance status", func() {
				goStagesForLatestInstance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
				goStagesForMinus1Instance := []a.GoStage{{Name: "Stage X", Status: "Unknown"}}
				goLatestInstance := a.GoInstance{Stages: goStagesForLatestInstance}
				goMinus1Instance := a.GoInstance{Stages: goStagesForMinus1Instance}
				goPreviousInstance := a.GoPreviousInstance{Result: "Passed"}
				goInstances := []a.GoInstance{goMinus1Instance, goLatestInstance}
				goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances, PreviousInstance: goPreviousInstance}}
				goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
				interests := a.NewInterests().Add("Pipeline One")
				goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
				simpleDashboard := goDashboard.ToSimpleDashboard()

				It("uses the status of previous instance", func() {
					pipelines := simpleDashboard.Pipelines
					Expect(pipelines).To(HaveLen(1))
					stages := pipelines[0].Stages
					Expect(stages).To(HaveLen(1))
					Expect(stages[0].Name).To(Equal("Stage X"))
					Expect(stages[0].Status).To(Equal("Passed"))
				})
			})
		})

		Context("with previous result as failed and current status as building", func() {
			goStagesForLatestInstance := []a.GoStage{{Name: "Stage X", Status: "Building"}}
			goLatestInstance := a.GoInstance{Stages: goStagesForLatestInstance}
			goPreviousInstance := a.GoPreviousInstance{Result: "Failed"}
			goInstances := []a.GoInstance{goLatestInstance}
			goPipelines := []a.GoPipeline{{Name: "Pipeline One", Instances: goInstances, PreviousInstance: goPreviousInstance}}
			goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
			interests := a.NewInterests().Add("Pipeline One")
			goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
			simpleDashboard := goDashboard.ToSimpleDashboard()

			It("uses marks the status as recovering", func() {
				pipelines := simpleDashboard.Pipelines
				Expect(pipelines).To(HaveLen(1))
				stages := pipelines[0].Stages
				Expect(stages).To(HaveLen(1))
				Expect(stages[0].Name).To(Equal("Stage X"))
				Expect(stages[0].Status).To(Equal("Recovering"))
			})
		})
	})

	Context("with multiple pipeline-group, pipelines, instances and stages", func() {
		goStage_1_old_1 := a.GoStage{Name: "Stage 1.1.1", Status: "Passed"}
		goStage_1_old_2 := a.GoStage{Name: "Stage 1.1.2", Status: "Failed"}
		goStage_1_new_1 := a.GoStage{Name: "Stage 1.2.1", Status: "Cancelled"}
		goStage_1_new_2 := a.GoStage{Name: "Stage 1.2.2", Status: "Failing"}
		goStage_2_old_1 := a.GoStage{Name: "Stage 2.1.1", Status: "Building"}
		goStage_2_old_2 := a.GoStage{Name: "Stage 2.1.2", Status: "Unknown"}
		goStage_2_new_1 := a.GoStage{Name: "Stage 2.2.1", Status: "Passed"}
		goStage_2_new_2 := a.GoStage{Name: "Stage 2.2.2", Status: "Failed"}
		goStages_1_old := []a.GoStage{goStage_1_old_1, goStage_1_old_2}
		goStages_1_new := []a.GoStage{goStage_1_new_1, goStage_1_new_2}
		goStages_2_old := []a.GoStage{goStage_2_old_1, goStage_2_old_2}
		goStages_2_new := []a.GoStage{goStage_2_new_1, goStage_2_new_2}
		goInstance_1_old := a.GoInstance{Stages: goStages_1_old}
		goInstance_1_new := a.GoInstance{Stages: goStages_1_new}
		goInstance_2_old := a.GoInstance{Stages: goStages_2_old}
		goInstance_2_new := a.GoInstance{Stages: goStages_2_new}
		goInstances_1 := []a.GoInstance{goInstance_1_old, goInstance_1_new}
		goInstances_2 := []a.GoInstance{goInstance_2_old, goInstance_2_new}
		goPipeline_1 := a.GoPipeline{Instances: goInstances_1, Name: "Pipeline One"}
		goPipeline_2 := a.GoPipeline{Instances: goInstances_2, Name: "Pipeline Two"}
		goPipelines := []a.GoPipeline{goPipeline_1, goPipeline_2}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline One:>Pipeline 1").Add("Pipeline Two")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("has simple-pipelines", func() {
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(HaveLen(2))
			pipeline_1 := pipelines[0]
			pipeline_2 := pipelines[1]
			Expect(pipeline_1.Name).To(Equal("Pipeline 1"))
			Expect(pipeline_2.Name).To(Equal("Pipeline Two"))
			stages_1 := pipeline_1.Stages
			stages_2 := pipeline_2.Stages
			Expect(stages_1).To(HaveLen(2))
			Expect(stages_2).To(HaveLen(2))
			Expect(stages_1[0]).To(Equal(a.SimpleStage{Name: "Stage 1.2.1", Status: "Cancelled"}))
			Expect(stages_1[1]).To(Equal(a.SimpleStage{Name: "Stage 1.2.2", Status: "Failing"}))
			Expect(stages_2[0]).To(Equal(a.SimpleStage{Name: "Stage 2.2.1", Status: "Passed"}))
			Expect(stages_2[1]).To(Equal(a.SimpleStage{Name: "Stage 2.2.2", Status: "Failed"}))
		})

		It("is sorted based on the interest", func() {
			interests := a.NewInterests().Add("Pipeline Two").Add("Pipeline One")
			goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
			simpleDashboard := goDashboard.ToSimpleDashboard()
			pipelines := simpleDashboard.Pipelines
			Expect(pipelines).To(HaveLen(2))
			pipeline_1 := pipelines[0]
			pipeline_2 := pipelines[1]
			Expect(pipeline_1.Name).To(Equal("Pipeline Two"))
			Expect(pipeline_2.Name).To(Equal("Pipeline One"))
			stages_1 := pipeline_1.Stages
			stages_2 := pipeline_2.Stages
			Expect(stages_1).To(HaveLen(2))
			Expect(stages_2).To(HaveLen(2))
			Expect(stages_1[0]).To(Equal(a.SimpleStage{Name: "Stage 2.2.1", Status: "Passed"}))
			Expect(stages_1[1]).To(Equal(a.SimpleStage{Name: "Stage 2.2.2", Status: "Failed"}))
			Expect(stages_2[0]).To(Equal(a.SimpleStage{Name: "Stage 1.2.1", Status: "Cancelled"}))
			Expect(stages_2[1]).To(Equal(a.SimpleStage{Name: "Stage 1.2.2", Status: "Failing"}))
		})
	})

	Context("with pipeline-group, non-matching pipeline, instance and stage", func() {
		goStages := []a.GoStage{{Name: "Stage One", Status: "Passed"}}
		goInstances := []a.GoInstance{{Stages: goStages}}
		goPipelines := []a.GoPipeline{{Instances: goInstances, Name: "Pipeline One"}}
		goPipelineGroups := []a.GoPipelineGroup{{Pipelines: goPipelines}}
		interests := a.NewInterests().Add("Pipeline")
		goDashboard := a.GoDashboard{PipelineGroups: goPipelineGroups, Interests: interests}
		simpleDashboard := goDashboard.ToSimpleDashboard()

		It("ignores non-matching pipelines", func() {
			Expect(simpleDashboard.Pipelines).To(BeEmpty())
		})

		It("gathers the names of the ignored pipelines", func() {
			Expect(simpleDashboard.Ignores).To(HaveLen(1))
			Expect(simpleDashboard.Ignores[0]).To(Equal("Pipeline One"))
		})
	})

	Context("unmarshal from JSON", func() {
		const dashboardJSON = `[{
		  "name": "Group",
		  "pipelines": [
		    {
		      "name": "Pipeline",
		      "instances": [
		        {
		          "stages": [
		            { "name": "StageOne", "status": "Passed" },
		            { "name": "StageTwo", "status": "Building" }
		          ]
		        }
		      ],
		      "previous_instance": {
		        "result": "Passed"
		      }
		    }
		  ]
		}
		]`

		It("is created from byte array", func() {
			goDashboard, err := a.NewGoPipelineGroups([]byte(dashboardJSON))
			Expect(err).To(BeNil())
			Expect(goDashboard).To(HaveLen(1))
			Expect(goDashboard[0].Pipelines).To(HaveLen(1))
			pipeline := goDashboard[0].Pipelines[0]
			Expect(pipeline.Name).To(Equal("Pipeline"))
			Expect(pipeline.Instances).To(HaveLen(1))
			instance := pipeline.Instances[0]
			Expect(instance.Stages).To(HaveLen(2))
			stages := instance.Stages
			Expect(stages[0].Name).To(Equal("StageOne"))
			Expect(stages[0].Status).To(Equal("Passed"))
			Expect(stages[1].Name).To(Equal("StageTwo"))
			Expect(stages[1].Status).To(Equal("Building"))
			previousInstance := pipeline.PreviousInstance
			Expect(previousInstance.Result).To(Equal("Passed"))
		})

		Context("on failure", func() {
			It("has error", func() {
				goDashboard, err := a.NewGoPipelineGroups([]byte(`Random`))
				Expect(err).ToNot(BeNil())
				Expect(goDashboard).To(BeNil())
			})
		})
	})
})
