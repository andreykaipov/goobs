// This file has been automatically generated. Don't edit it.

package record

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'record' requests.
type Client struct {
	client *api.Client
}

// NewRecord returns a new 'record' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
