// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputMuteParams represents the params body for the "SetInputMute" request.
Sets the audio mute state of an input.
*/
type SetInputMuteParams struct {
	requests.ParamsBasic

	// Whether to mute the input or not
	InputMuted bool `json:"inputMuted,omitempty"`

	// Name of the input to set the mute state of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "SetInputMute".
func (o *SetInputMuteParams) GetSelfName() string {
	return "SetInputMute"
}

/*
SetInputMuteResponse represents the response body for the "SetInputMute" request.
Sets the audio mute state of an input.
*/
type SetInputMuteResponse struct {
	requests.ResponseBasic
}

// SetInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputMute(params *SetInputMuteParams) (*SetInputMuteResponse, error) {
	data := &SetInputMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
