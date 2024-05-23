package bigpanda

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEnvironments - Returns list of environments
func (c *Client) GetEnvironments() ([]Environment, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/v2.0/environments", c.BigPandaURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	environments := []Environment{}
	err = json.Unmarshal(body, &environments)
	if err != nil {
		return nil, err
	}

	return environments, nil
}

// CreateEnvironment - Create new environment
func (c *Client) CreateEnvironment(environmentItems []EnvironmentItem, authToken *string) (*Order, error) {
	rb, err := json.Marshal(environmentItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/v2.0/environments", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	environment := Environment{}
	err = json.Unmarshal(body, &environment)
	if err != nil {
		return nil, err
	}

	return &environment, nil
}
