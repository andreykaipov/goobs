// This file has been automatically generated. Don't edit it.

package sources

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'sources' requests.
type Client struct {
	client *api.Client
}

// NewSources returns a new 'sources' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
