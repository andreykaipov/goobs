// This file has been automatically generated. Don't edit it.

package mediainputs

// Represents the request body for the TriggerMediaInputAction request.
type TriggerMediaInputActionParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Identifier of the `ObsMediaInputAction` enum
	MediaAction string `json:"mediaAction,omitempty"`
}

// Returns the associated request.
func (o *TriggerMediaInputActionParams) GetRequestName() string {
	return "TriggerMediaInputAction"
}

// Represents the response body for the TriggerMediaInputAction request.
type TriggerMediaInputActionResponse struct{}

// Triggers an action on a media input.
func (c *Client) TriggerMediaInputAction(
	params *TriggerMediaInputActionParams,
) (*TriggerMediaInputActionResponse, error) {
	data := &TriggerMediaInputActionResponse{}
	return data, c.SendRequest(params, data)
}
