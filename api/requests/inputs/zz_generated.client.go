// This file has been automatically generated. Don't edit it.

package inputs

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'inputs' requests.
type Client struct {
	client *api.Client
}

// NewInputs returns a new 'inputs' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
