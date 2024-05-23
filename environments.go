package bigpanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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

// GetOrder - Returns a specifc order
func (c *Client) GetOrder(orderID string) (*Order, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/v2.0/environments/%s", c.BigPandaURL, orderID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}


// CreateOrder - Create new order
func (c *Client) CreateOrder(orderItems []OrderItem, ctx Context) (*Order, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	tflog.Info(ctx, "Creating a new environment in the client")
	
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/v2.0/environments", c.BigPandaURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
