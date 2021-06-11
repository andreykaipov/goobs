// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetBrowserSourcePropertiesParams represents the params body for the "GetBrowserSourceProperties" request.
Get current properties for a Browser Source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// Name just returns "GetBrowserSourceProperties".
func (o *GetBrowserSourcePropertiesParams) Name() string {
	return "GetBrowserSourceProperties"
}

/*
GetBrowserSourcePropertiesResponse represents the response body for the "GetBrowserSourceProperties" request.
Get current properties for a Browser Source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesResponse struct {
	requests.ResponseBasic

	// CSS to inject.
	Css string `json:"css"`

	// Framerate.
	Fps int `json:"fps"`

	// Height.
	Height int `json:"height"`

	// Indicates that a local file is in use.
	IsLocalFile bool `json:"is_local_file"`

	// file path.
	LocalFile string `json:"local_file"`

	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`

	// Source name.
	Source string `json:"source"`

	// Url.
	Url string `json:"url"`

	// Width.
	Width int `json:"width"`
}

// GetBrowserSourceProperties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) GetBrowserSourceProperties(
	params *GetBrowserSourcePropertiesParams,
) (*GetBrowserSourcePropertiesResponse, error) {
	data := &GetBrowserSourcePropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
