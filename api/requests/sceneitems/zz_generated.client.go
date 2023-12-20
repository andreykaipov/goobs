// This file has been automatically generated. Don't edit it.

package sceneitems

import api "github.com/andreykaipov/goobs/api"

type _response = api.ResponseCommon

// Client represents a client for 'scene items' requests.
type Client struct {
	client *api.Client
}

// NewSceneItems returns a new 'scene items' client.
func NewClient(c *api.Client) *Client {
	return &Client{client: c}
}
