// This file has been automatically generated. Don't edit it.

package outputs

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'outputs' requests.
type Client struct {
	client *api.Client
}

// NewOutputs returns a new 'outputs' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
