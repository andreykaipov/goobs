// This file has been automatically generated. Don't edit it.

package transitions

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'transitions' requests.
type Client struct {
	client *api.Client
}

// NewTransitions returns a new 'transitions' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
