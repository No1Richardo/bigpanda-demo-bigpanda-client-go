package bigpanda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const BigPandaURL string = "https://api.bigpanda.io/"

// Client -
type Client struct {
	BigPandaURL	string
	HTTPClient *http.Client
	Token      string 
}



// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default BigPanda URL
		BigPandaURL: BigPandaURL,
	}

	if host != nil {
		c.BigPandaURL = *host
	}

	if token != nil {
		c.Token = *token
	}
	
	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", "Bearer " + token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
