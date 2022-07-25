// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleInputMuteParams represents the params body for the "ToggleInputMute" request.
Toggles the audio mute state of an input.
*/
type ToggleInputMuteParams struct {
	requests.ParamsBasic

	// Name of the input to toggle the mute state of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "ToggleInputMute".
func (o *ToggleInputMuteParams) GetSelfName() string {
	return "ToggleInputMute"
}

/*
ToggleInputMuteResponse represents the response body for the "ToggleInputMute" request.
Toggles the audio mute state of an input.
*/
type ToggleInputMuteResponse struct {
	requests.ResponseBasic

	// Whether the input has been muted or unmuted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// ToggleInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ToggleInputMute(params *ToggleInputMuteParams) (*ToggleInputMuteResponse, error) {
	data := &ToggleInputMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
