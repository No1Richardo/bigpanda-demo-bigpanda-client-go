package bigpanda

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetIncidentTags - Returns list of incidenttags (no auth required)
func (c *Client) GetIncidentTags() ([]IncidentTag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/v2.0/incidents/tags/definitions", c.BigPandaURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	incidenttags := []IncidentTag{}
	err = json.Unmarshal(body, &incidenttags)
	if err != nil {
		return nil, err
	}

	return incidenttags, nil
}
