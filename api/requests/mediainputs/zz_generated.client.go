// This file has been automatically generated. Don't edit it.

package mediainputs

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'media inputs' requests.
type Client struct {
	client *api.Client
}

// NewMediaInputs returns a new 'media inputs' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
