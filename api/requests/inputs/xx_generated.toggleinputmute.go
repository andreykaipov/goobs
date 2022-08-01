// This file has been automatically generated. Don't edit it.

package inputs

/*
ToggleInputMuteParams represents the params body for the "ToggleInputMute" request.
Toggles the audio mute state of an input.
*/
type ToggleInputMuteParams struct {
	// Name of the input to toggle the mute state of
	InputName string `json:"inputName,omitempty"`
}

/*
ToggleInputMuteResponse represents the response body for the "ToggleInputMute" request.
Toggles the audio mute state of an input.
*/
type ToggleInputMuteResponse struct {
	// Whether the input has been muted or unmuted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// ToggleInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ToggleInputMute(params *ToggleInputMuteParams) (*ToggleInputMuteResponse, error) {
	resp, err := c.SendRequest("ToggleInputMute", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ToggleInputMuteResponse), nil
}
