// This file has been automatically generated. Don't edit it.

package mediainputs

/*
TriggerMediaInputActionParams represents the params body for the "TriggerMediaInputAction" request.
Triggers an action on a media input.
*/
type TriggerMediaInputActionParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Identifier of the `ObsMediaInputAction` enum
	MediaAction string `json:"mediaAction,omitempty"`
}

/*
TriggerMediaInputActionResponse represents the response body for the "TriggerMediaInputAction" request.
Triggers an action on a media input.
*/
type TriggerMediaInputActionResponse struct{}

// TriggerMediaInputAction sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerMediaInputAction(
	params *TriggerMediaInputActionParams,
) (*TriggerMediaInputActionResponse, error) {
	resp, err := c.SendRequest("TriggerMediaInputAction", params)
	if err != nil {
		return nil, err
	}
	return resp.(*TriggerMediaInputActionResponse), nil
}
