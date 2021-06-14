// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMuteParams represents the params body for the "GetMute" request.
Get the mute status of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// Name just returns "GetMute".
func (o *GetMuteParams) Name() string {
	return "GetMute"
}

/*
GetMuteResponse represents the response body for the "GetMute" request.
Get the mute status of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteResponse struct {
	requests.ResponseBasic

	// Mute status of the source.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`
}

// GetMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMute(params *GetMuteParams) (*GetMuteResponse, error) {
	data := &GetMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
