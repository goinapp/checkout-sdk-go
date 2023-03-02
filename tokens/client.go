package tokens

import (
	"encoding/json"
	"net/http"

	"github.com/goinapp/checkout-sdk-go"
	"github.com/goinapp/checkout-sdk-go/httpclient"
)

const path = "tokens"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Request ...
func (c *Client) Request(request *Request) (*Response, error) {
	response, err := c.API.Post("/"+path, request, nil)
	resp := &Response{
		StatusResponse: response,
	}
	if err != nil {
		return resp, err
	}
	if response.StatusCode == http.StatusCreated {
		var created Created
		err = json.Unmarshal(response.ResponseBody, &created)
		resp.Created = &created
	}
	return resp, err
}
