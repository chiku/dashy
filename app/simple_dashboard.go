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
}
type SimpleDashboard struct {
	Pipelines []SimplePipeline
	Ignores   []string
}

func (dashboard *SimpleDashboard) ToJSON() (output []byte, err error) {
	output, err = json.Marshal(dashboard.Pipelines)
	if err != nil {
		return nil, fmt.Errorf("error marshalling simple dashboard JSON :%s", err.Error())
	}

	return output, nil
}
