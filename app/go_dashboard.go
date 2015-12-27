package app

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const (
	building   = "Building"
	unknown    = "Unknown"
	failed     = "Failed"
	recovering = "Recovering"
)

type GoStage struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
type GoInstance struct {
	Stages []GoStage `json:"stages"`
}
type GoPreviousInstance struct {
	Result string `json:"result"`
}
type GoPipeline struct {
	Name             string             `json:"name"`
	Instances        []GoInstance       `json:"instances"`
	PreviousInstance GoPreviousInstance `json:"previous_instance"`
}
type GoPipelineGroup struct {
	Pipelines []GoPipeline `json:"pipelines"`
}
type GoDashboard struct {
	PipelineGroups []GoPipelineGroup
	Interests      *Interests
}

func NewGoPipelineGroups(body []byte) ([]GoPipelineGroup, error) {
	var dashboard []GoPipelineGroup
	err := json.Unmarshal(body, &dashboard)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling external dashboard JSON: %s", err.Error())
	}
	return dashboard, nil
}

func (goDashboard *GoDashboard) ToSimpleDashboard() *SimpleDashboard {
	dashboard := &SimpleDashboard{
		Pipelines: []SimplePipeline{},
		Ignores:   []string{},
	}

	for _, goPipelineGroup := range goDashboard.PipelineGroups {
		for _, goPipeline := range goPipelineGroup.Pipelines {
			interestOrder, _ := goDashboard.Interests.PipelineName(goPipeline.Name)
			if interestOrder != -1 {
				stages := []SimpleStage{}
				if len(goPipeline.Instances) > 0 {

					goInstance := goPipeline.Instances[len(goPipeline.Instances)-1]
					for _, goStage := range goInstance.Stages {
						status := traverseStatusInInstances(goStage, goPipeline.Instances, goPipeline.PreviousInstance)
						stages = append(stages, SimpleStage{Name: goStage.Name, Status: status})
					}
					if len(stages) > 0 {
						dashboard.Pipelines = append(dashboard.Pipelines, (SimplePipeline{Name: goPipeline.Name, Stages: stages}).Order(interestOrder))
					}
				}
			} else {
				dashboard.Ignores = append(dashboard.Ignores, goPipeline.Name)
			}
		}
	}

	sort.Sort(ByOrder(dashboard.Pipelines))

	return dashboard
}

func traverseStatusInInstances(currentStage GoStage, instances []GoInstance, previousInstance GoPreviousInstance) string {
	selfStatus := currentStage.Status
	previousInstanceResult := previousInstance.Result

	if previousInstanceResult == failed && strings.EqualFold(selfStatus, building) {
		return recovering
	}

	if !strings.EqualFold(selfStatus, unknown) {
		return selfStatus
	}

	olderInstances := instances[0 : len(instances)-1]
	olderInstanceStatus := findKnownStatusInInstances(currentStage, olderInstances)
	if !strings.EqualFold(olderInstanceStatus, unknown) {
		return olderInstanceStatus
	}

	if previousInstanceResult != "" && !strings.EqualFold(previousInstanceResult, unknown) {
		return previousInstanceResult
	}

	return unknown
}

func findKnownStatusInInstances(currentStage GoStage, instances []GoInstance) string {
	for i := len(instances) - 1; i >= 0; i-- {
		instance := instances[i]
		for j := len(instance.Stages) - 1; j >= 0; j-- {
			stage := instance.Stages[j]
			if currentStage.Name == stage.Name && !strings.EqualFold(stage.Status, unknown) {
				return stage.Status
			}
		}
	}

	return unknown
}
