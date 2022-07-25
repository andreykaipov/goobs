// This file has been automatically generated. Don't edit it.

package mediainputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TriggerMediaInputActionParams represents the params body for the "TriggerMediaInputAction" request.
Triggers an action on a media input.
*/
type TriggerMediaInputActionParams struct {
	requests.ParamsBasic

	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Identifier of the `ObsMediaInputAction` enum
	MediaAction string `json:"mediaAction,omitempty"`
}

// GetSelfName just returns "TriggerMediaInputAction".
func (o *TriggerMediaInputActionParams) GetSelfName() string {
	return "TriggerMediaInputAction"
}

/*
TriggerMediaInputActionResponse represents the response body for the "TriggerMediaInputAction" request.
Triggers an action on a media input.
*/
type TriggerMediaInputActionResponse struct {
	requests.ResponseBasic
}

// TriggerMediaInputAction sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerMediaInputAction(
	params *TriggerMediaInputActionParams,
) (*TriggerMediaInputActionResponse, error) {
	data := &TriggerMediaInputActionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
