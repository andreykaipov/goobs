// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.
Transitions the currently previewed scene to the main output.
Will return an `error` if Studio Mode is not enabled.
Since 4.1.0.
*/
type TransitionToProgramParams struct {
	requests.ParamsBasic

	WithTransition *WithTransition `json:"with-transition,omitempty"`
}

type WithTransition struct {
	// Transition duration (in milliseconds).
	Duration int `json:"duration,omitempty"`

	// Name of the transition.
	Name string `json:"name,omitempty"`
}

// GetSelfName just returns "TransitionToProgram".
func (o *TransitionToProgramParams) GetSelfName() string {
	return "TransitionToProgram"
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.
Transitions the currently previewed scene to the main output.
Will return an `error` if Studio Mode is not enabled.
Since v4.1.0.
*/
type TransitionToProgramResponse struct {
	requests.ResponseBasic
}

// TransitionToProgram sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TransitionToProgram(params *TransitionToProgramParams) (*TransitionToProgramResponse, error) {
	data := &TransitionToProgramResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
