package app

import (
	"encoding/json"
	"fmt"
)

type SimpleStage struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
type SimplePipeline struct {
	Name   string        `json:"name"`
	Stages []SimpleStage `json:"stages"`
	order  int
}
type SimpleDashboard struct {
	Pipelines []SimplePipeline
	Ignores   []string
}

func (dashboard SimplePipeline) Order(order int) SimplePipeline {
	dashboard.order = order
	return dashboard
}

func (dashboard *SimpleDashboard) ToJSON() (output []byte, err error) {
	output, err = json.Marshal(dashboard.Pipelines)
	if err != nil {
		return nil, fmt.Errorf("error marshalling simple dashboard JSON :%s", err.Error())
	}

	return output, nil
}

type ByOrder []SimplePipeline

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].order < a[j].order }
