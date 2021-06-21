// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleMuteParams represents the params body for the "ToggleMute" request.
Inverts the mute status of a specified source.
Since 4.0.0.
*/
type ToggleMuteParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "ToggleMute".
func (o *ToggleMuteParams) GetSelfName() string {
	return "ToggleMute"
}

/*
ToggleMuteResponse represents the response body for the "ToggleMute" request.
Inverts the mute status of a specified source.
Since v4.0.0.
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
