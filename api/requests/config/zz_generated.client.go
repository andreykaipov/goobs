// This file has been automatically generated. Don't edit it.

package config

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'config' requests.
type Client struct {
	client *api.Client
}

// NewConfig returns a new 'config' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
