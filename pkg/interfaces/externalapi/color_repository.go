package externalapi

import (
	"encoding/json"
	"fmt"
	"go-issue-tracker/pkg/domain"
	"io/ioutil"
)

// ColorRepository is a repository
type ColorRepository struct {
	Endpoint   string
	HTTPClient HTTPClient
}

// NewColorRepository to create ColorRepository
func NewColorRepository(endpoint string, httpClient HTTPClient) *ColorRepository {
	return &ColorRepository{
		Endpoint:   endpoint,
		HTTPClient: httpClient,
	}
}

// GetColor to get color
func (r *ColorRepository) GetColor() (domain.Color, error) {
	var color domain.Color

	response, err := r.HTTPClient.Get(r.Endpoint)
	if err != nil {
		return color, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return color, fmt.Errorf("endpoint %s not found", r.Endpoint)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return color, err
	}

	err = json.Unmarshal(body, &color)
	if err != nil {
		return color, err
	}

	return color, nil
}
