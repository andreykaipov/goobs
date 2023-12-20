// This file has been automatically generated. Don't edit it.

package filters

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'filters' requests.
type Client struct {
	client *api.Client
}

// NewFilters returns a new 'filters' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
