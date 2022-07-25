// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneTransitionDurationParams represents the params body for the "SetCurrentSceneTransitionDuration" request.
Sets the duration of the current scene transition, if it is not fixed.
*/
type SetCurrentSceneTransitionDurationParams struct {
	requests.ParamsBasic

	// Duration in milliseconds
	TransitionDuration float64 `json:"transitionDuration,omitempty"`
}

// GetSelfName just returns "SetCurrentSceneTransitionDuration".
func (o *SetCurrentSceneTransitionDurationParams) GetSelfName() string {
	return "SetCurrentSceneTransitionDuration"
}

/*
SetCurrentSceneTransitionDurationResponse represents the response body for the "SetCurrentSceneTransitionDuration" request.
Sets the duration of the current scene transition, if it is not fixed.
*/
type SetCurrentSceneTransitionDurationResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransitionDuration(
	params *SetCurrentSceneTransitionDurationParams,
) (*SetCurrentSceneTransitionDurationResponse, error) {
	data := &SetCurrentSceneTransitionDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
