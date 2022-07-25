// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputAudioBalanceParams represents the params body for the "SetInputAudioBalance" request.
Sets the audio balance of an input.
*/
type SetInputAudioBalanceParams struct {
	requests.ParamsBasic

	// New audio balance value
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`

	// Name of the input to set the audio balance of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "SetInputAudioBalance".
func (o *SetInputAudioBalanceParams) GetSelfName() string {
	return "SetInputAudioBalance"
}

/*
SetInputAudioBalanceResponse represents the response body for the "SetInputAudioBalance" request.
Sets the audio balance of an input.
*/
type SetInputAudioBalanceResponse struct {
	requests.ResponseBasic
}

// SetInputAudioBalance sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioBalance(params *SetInputAudioBalanceParams) (*SetInputAudioBalanceResponse, error) {
	data := &SetInputAudioBalanceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
