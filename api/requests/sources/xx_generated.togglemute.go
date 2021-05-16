// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleMuteParams represents the params body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// Name just returns "ToggleMute".
func (o *ToggleMuteParams) Name() string {
	return "ToggleMute"
}

/*
ToggleMuteResponse represents the response body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteResponse struct {
	requests.ResponseBasic
}

// ToggleMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ToggleMute(params *ToggleMuteParams) (*ToggleMuteResponse, error) {
	data := &ToggleMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
