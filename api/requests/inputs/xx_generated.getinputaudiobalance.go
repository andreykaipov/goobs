// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputAudioBalanceParams represents the params body for the "GetInputAudioBalance" request.
Gets the audio balance of an input.
*/
type GetInputAudioBalanceParams struct {
	requests.ParamsBasic

	// Name of the input to get the audio balance of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputAudioBalance".
func (o *GetInputAudioBalanceParams) GetSelfName() string {
	return "GetInputAudioBalance"
}

/*
GetInputAudioBalanceResponse represents the response body for the "GetInputAudioBalance" request.
Gets the audio balance of an input.
*/
type GetInputAudioBalanceResponse struct {
	requests.ResponseBasic

	// Audio balance value from 0.0-1.0
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`
}

// GetInputAudioBalance sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioBalance(params *GetInputAudioBalanceParams) (*GetInputAudioBalanceResponse, error) {
	data := &GetInputAudioBalanceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
