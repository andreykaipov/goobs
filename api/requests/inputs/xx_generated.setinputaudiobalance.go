// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputAudioBalance request.
type SetInputAudioBalanceParams struct {
	// New audio balance value
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`

	// Name of the input to set the audio balance of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *SetInputAudioBalanceParams) GetRequestName() string {
	return "SetInputAudioBalance"
}

// Represents the response body for the SetInputAudioBalance request.
type SetInputAudioBalanceResponse struct{}

// Sets the audio balance of an input.
func (c *Client) SetInputAudioBalance(params *SetInputAudioBalanceParams) (*SetInputAudioBalanceResponse, error) {
	data := &SetInputAudioBalanceResponse{}
	return data, c.SendRequest(params, data)
}
