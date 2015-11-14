package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseHTTPResponse(response *http.Response, err error) ([]GoPipelineGroup, error) {
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the HTTP status code was %d", response.StatusCode)
	}
	if response.Body == nil {
		return nil, errors.New("the response didn't have a body")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}
	defer response.Body.Close()

	dashboard, err := GoPipelineGroupsFromJSON(body)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}
