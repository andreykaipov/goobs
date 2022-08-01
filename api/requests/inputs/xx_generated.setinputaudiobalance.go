// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputAudioBalanceParams represents the params body for the "SetInputAudioBalance" request.
Sets the audio balance of an input.
*/
type SetInputAudioBalanceParams struct {
	// New audio balance value
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`

	// Name of the input to set the audio balance of
	InputName string `json:"inputName,omitempty"`
}

/*
SetInputAudioBalanceResponse represents the response body for the "SetInputAudioBalance" request.
Sets the audio balance of an input.
*/
type SetInputAudioBalanceResponse struct{}

// SetInputAudioBalance sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioBalance(params *SetInputAudioBalanceParams) (*SetInputAudioBalanceResponse, error) {
	resp, err := c.SendRequest("SetInputAudioBalance", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputAudioBalanceResponse), nil
}
