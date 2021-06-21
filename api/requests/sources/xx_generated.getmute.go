// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMuteParams represents the params body for the "GetMute" request.
Get the mute status of a specified source.
Since 4.0.0.
*/
type GetMuteParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "GetMute".
func (o *GetMuteParams) GetSelfName() string {
	return "GetMute"
}

/*
GetMuteResponse represents the response body for the "GetMute" request.
Get the mute status of a specified source.
Since v4.0.0.
*/
type GetMuteResponse struct {
	requests.ResponseBasic

	// Mute status of the source.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name,omitempty"`
}

// GetMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMute(params *GetMuteParams) (*GetMuteResponse, error) {
	data := &GetMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
