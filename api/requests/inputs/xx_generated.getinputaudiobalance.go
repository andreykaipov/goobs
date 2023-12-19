// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputAudioBalance request.
type GetInputAudioBalanceParams struct {
	// Name of the input to get the audio balance of
	InputName *string `json:"inputName,omitempty"`
}

func NewGetInputAudioBalanceParams() *GetInputAudioBalanceParams {
	return &GetInputAudioBalanceParams{}
}
func (o *GetInputAudioBalanceParams) WithInputName(x string) *GetInputAudioBalanceParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *GetInputAudioBalanceParams) GetRequestName() string {
	return "GetInputAudioBalance"
}

// Represents the response body for the GetInputAudioBalance request.
type GetInputAudioBalanceResponse struct {
	_response

	// Audio balance value from 0.0-1.0
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`
}

// Gets the audio balance of an input.
func (c *Client) GetInputAudioBalance(params *GetInputAudioBalanceParams) (*GetInputAudioBalanceResponse, error) {
	data := &GetInputAudioBalanceResponse{}
	return data, c.SendRequest(params, data)
}
