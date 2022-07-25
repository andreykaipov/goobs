// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputMuteParams represents the params body for the "GetInputMute" request.
Gets the audio mute state of an input.
*/
type GetInputMuteParams struct {
	requests.ParamsBasic

	// Name of input to get the mute state of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputMute".
func (o *GetInputMuteParams) GetSelfName() string {
	return "GetInputMute"
}

/*
GetInputMuteResponse represents the response body for the "GetInputMute" request.
Gets the audio mute state of an input.
*/
type GetInputMuteResponse struct {
	requests.ResponseBasic

	// Whether the input is muted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// GetInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputMute(params *GetInputMuteParams) (*GetInputMuteResponse, error) {
	data := &GetInputMuteResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
