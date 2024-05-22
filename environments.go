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
