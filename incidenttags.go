package bigpanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetCoffees - Returns list of incidenttags (no auth required)
func (c *Client) GetIncidentTags() ([]IncidentTag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/v2.0/incidents/tags/definitions", c.HostURL), nil)
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

	return coffees, nil
}
