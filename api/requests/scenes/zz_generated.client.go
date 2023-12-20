// This file has been automatically generated. Don't edit it.

package scenes

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'scenes' requests.
type Client struct {
	client *api.Client
}

// NewScenes returns a new 'scenes' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
