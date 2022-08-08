// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the ToggleInputMute request.
type ToggleInputMuteParams struct {
	// Name of the input to toggle the mute state of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *ToggleInputMuteParams) GetRequestName() string {
	return "ToggleInputMute"
}

// Represents the response body for the ToggleInputMute request.
type ToggleInputMuteResponse struct {
	// Whether the input has been muted or unmuted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// Toggles the audio mute state of an input.
func (c *Client) ToggleInputMute(params *ToggleInputMuteParams) (*ToggleInputMuteResponse, error) {
	data := &ToggleInputMuteResponse{}
	return data, c.SendRequest(params, data)
}
