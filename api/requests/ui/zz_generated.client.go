// This file has been automatically generated. Don't edit it.

package ui

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'ui' requests.
type Client struct {
	client *api.Client
}

// NewUi returns a new 'ui' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
