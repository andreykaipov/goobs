// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTransitionDurationParams represents the params body for the "SetTransitionDuration" request.
Set the duration of the currently selected transition if supported.
Since 4.0.0.
*/
type SetTransitionDurationParams struct {
	requests.ParamsBasic

	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration,omitempty"`
}

// GetSelfName just returns "SetTransitionDuration".
func (o *SetTransitionDurationParams) GetSelfName() string {
	return "SetTransitionDuration"
}

/*
SetTransitionDurationResponse represents the response body for the "SetTransitionDuration" request.
Set the duration of the currently selected transition if supported.
Since v4.0.0.
*/
type SetTransitionDurationResponse struct {
	requests.ResponseBasic
}

// SetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTransitionDuration(params *SetTransitionDurationParams) (*SetTransitionDurationResponse, error) {
	data := &SetTransitionDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
