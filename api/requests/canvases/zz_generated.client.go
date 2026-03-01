// This file has been automatically generated. Don't edit it.

package canvases

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'canvases' requests.
type Client struct {
	client *api.Client
}

// NewCanvases returns a new 'canvases' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
