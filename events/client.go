package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goinapp/checkout-sdk-go"
	"github.com/goinapp/checkout-sdk-go/httpclient"
	"github.com/google/go-querystring/query"
)

const path = "events"

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

// RetrieveEventTypes -
func (c *Client) RetrieveEventTypes(request *Request) (*Response, error) {

	value, _ := query.Values(request.EventTypeRequest)
	var query string = value.Encode()
	var urlPath string = "/event-types" + "?"
	resp, err := c.API.Get(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var eventTypes []EventType
		err = json.Unmarshal(resp.ResponseBody, &eventTypes)
		response.EventTypes = eventTypes
		return response, err
	}
	return response, err
}

// RetrieveEvents -
func (c *Client) RetrieveEvents(request *Request) (*Response, error) {

	value, _ := query.Values(request.QueryParameter)
	var query string = value.Encode()
	var urlPath string = "/" + path + "?"
	resp, err := c.API.Get(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var events Events
		err = json.Unmarshal(resp.ResponseBody, &events)
		response.Events = &events
		return response, err
	}
	return response, err

}

// RetrieveEvent -
func (c *Client) RetrieveEvent(eventID string) (*Response, error) {

	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, eventID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var event Event
		err = json.Unmarshal(resp.ResponseBody, &event)
		response.Event = &event
		return response, err
	}
	return response, err
}

// RetrieveEventNotification -
func (c *Client) RetrieveEventNotification(eventID string, notificationID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/%v/%v/notifications/%v", path, eventID, notificationID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var notification Notification
		err = json.Unmarshal(resp.ResponseBody, &notification)
		response.Notification = &notification
		return response, err
	}
	return response, err
}

// Retry -
func (c *Client) Retry(eventID string, webhookID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/%v/%v/webhooks/%v/retry", path, eventID, webhookID), nil, nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		return response, err
	}
	return response, err
}

// RetryAll -
func (c *Client) RetryAll(eventID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/%v/%v/webhooks/retry", path, eventID), nil, nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		return response, err
	}
	return response, err
}
