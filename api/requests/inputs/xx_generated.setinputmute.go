// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputMuteParams represents the params body for the "SetInputMute" request.
Sets the audio mute state of an input.
*/
type SetInputMuteParams struct {
	// Whether to mute the input or not
	InputMuted *bool `json:"inputMuted,omitempty"`

	// Name of the input to set the mute state of
	InputName string `json:"inputName,omitempty"`
}

/*
SetInputMuteResponse represents the response body for the "SetInputMute" request.
Sets the audio mute state of an input.
*/
type SetInputMuteResponse struct{}

// SetInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputMute(params *SetInputMuteParams) (*SetInputMuteResponse, error) {
	resp, err := c.SendRequest("SetInputMute", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputMuteResponse), nil
}
