// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetBrowserSourcePropertiesParams represents the params body for the "SetBrowserSourceProperties" request.
Set current properties for a Browser Source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesParams struct {
	requests.ParamsBasic

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

	// Visibility of the scene item.
	Render bool `json:"render"`

	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`

	// Name of the source.
	Source string `json:"source"`

	// Url.
	Url string `json:"url"`

	// Width.
	Width int `json:"width"`
}

// Name just returns "SetBrowserSourceProperties".
func (o *SetBrowserSourcePropertiesParams) Name() string {
	return "SetBrowserSourceProperties"
}

/*
SetBrowserSourcePropertiesResponse represents the response body for the "SetBrowserSourceProperties" request.
Set current properties for a Browser Source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesResponse struct {
	requests.ResponseBasic
}

// SetBrowserSourceProperties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) SetBrowserSourceProperties(
	params *SetBrowserSourcePropertiesParams,
) (*SetBrowserSourcePropertiesResponse, error) {
	data := &SetBrowserSourcePropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
