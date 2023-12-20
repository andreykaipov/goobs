// This file has been automatically generated. Don't edit it.

package general

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'general' requests.
type Client struct {
	client *api.Client
}

// NewGeneral returns a new 'general' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
