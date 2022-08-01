// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputAudioBalanceParams represents the params body for the "GetInputAudioBalance" request.
Gets the audio balance of an input.
*/
type GetInputAudioBalanceParams struct {
	// Name of the input to get the audio balance of
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputAudioBalanceResponse represents the response body for the "GetInputAudioBalance" request.
Gets the audio balance of an input.
*/
type GetInputAudioBalanceResponse struct {
	// Audio balance value from 0.0-1.0
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`
}

// GetInputAudioBalance sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioBalance(params *GetInputAudioBalanceParams) (*GetInputAudioBalanceResponse, error) {
	resp, err := c.SendRequest("GetInputAudioBalance", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputAudioBalanceResponse), nil
}
