// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputMuteParams represents the params body for the "GetInputMute" request.
Gets the audio mute state of an input.
*/
type GetInputMuteParams struct {
	// Name of input to get the mute state of
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputMuteResponse represents the response body for the "GetInputMute" request.
Gets the audio mute state of an input.
*/
type GetInputMuteResponse struct {
	// Whether the input is muted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// GetInputMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputMute(params *GetInputMuteParams) (*GetInputMuteResponse, error) {
	resp, err := c.SendRequest("GetInputMute", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputMuteResponse), nil
}
