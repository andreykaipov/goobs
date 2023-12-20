// This file has been automatically generated. Don't edit it.

package stream

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'stream' requests.
type Client struct {
	client *api.Client
}

// NewStream returns a new 'stream' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
