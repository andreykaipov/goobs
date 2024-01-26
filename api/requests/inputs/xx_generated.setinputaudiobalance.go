// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputAudioBalance request.
type SetInputAudioBalanceParams struct {
	// New audio balance value
	InputAudioBalance *float64 `json:"inputAudioBalance,omitempty"`

	// Name of the input to set the audio balance of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to set the audio balance of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputAudioBalanceParams() *SetInputAudioBalanceParams {
	return &SetInputAudioBalanceParams{}
}
func (o *SetInputAudioBalanceParams) WithInputAudioBalance(x float64) *SetInputAudioBalanceParams {
	o.InputAudioBalance = &x
	return o
}
func (o *SetInputAudioBalanceParams) WithInputName(x string) *SetInputAudioBalanceParams {
	o.InputName = &x
	return o
}
func (o *SetInputAudioBalanceParams) WithInputUuid(x string) *SetInputAudioBalanceParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputAudioBalanceParams) GetRequestName() string {
	return "SetInputAudioBalance"
}

// Represents the response body for the SetInputAudioBalance request.
type SetInputAudioBalanceResponse struct {
	_response
}

// Sets the audio balance of an input.
func (c *Client) SetInputAudioBalance(params *SetInputAudioBalanceParams) (*SetInputAudioBalanceResponse, error) {
	data := &SetInputAudioBalanceResponse{}
	return data, c.client.SendRequest(params, data)
}
