// This file has been automatically generated. Don't edit it.

package canvas

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'canvas' requests.
type Client struct {
	client *api.Client
}

// NewCanvas returns a new 'canvas' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
