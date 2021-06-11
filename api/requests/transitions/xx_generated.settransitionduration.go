// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTransitionDurationParams represents the params body for the "SetTransitionDuration" request.
Set the duration of the currently selected transition if supported.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationParams struct {
	requests.ParamsBasic

	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}

// Name just returns "SetTransitionDuration".
func (o *SetTransitionDurationParams) Name() string {
	return "SetTransitionDuration"
}

/*
SetTransitionDurationResponse represents the response body for the "SetTransitionDuration" request.
Set the duration of the currently selected transition if supported.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationResponse struct {
	requests.ResponseBasic
}

// SetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTransitionDuration(
	params *SetTransitionDurationParams,
) (*SetTransitionDurationResponse, error) {
	data := &SetTransitionDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
