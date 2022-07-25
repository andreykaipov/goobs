// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentSceneTransitionParams represents the params body for the "GetCurrentSceneTransition" request.
Gets information about the current scene transition.
*/
type GetCurrentSceneTransitionParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentSceneTransition".
func (o *GetCurrentSceneTransitionParams) GetSelfName() string {
	return "GetCurrentSceneTransition"
}

/*
GetCurrentSceneTransitionResponse represents the response body for the "GetCurrentSceneTransition" request.
Gets information about the current scene transition.
*/
type GetCurrentSceneTransitionResponse struct {
	requests.ResponseBasic

	// Whether the transition supports being configured
	TransitionConfigurable bool `json:"transitionConfigurable,omitempty"`

	// Configured transition duration in milliseconds. `null` if transition is fixed
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Whether the transition uses a fixed (unconfigurable) duration
	TransitionFixed bool `json:"transitionFixed,omitempty"`

	// Kind of the transition
	TransitionKind string `json:"transitionKind,omitempty"`

	// Name of the transition
	TransitionName string `json:"transitionName,omitempty"`

	// Object of settings for the transition. `null` if transition is not configurable
	TransitionSettings interface{} `json:"transitionSettings,omitempty"`
}

// GetCurrentSceneTransition sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentSceneTransition(
	paramss ...*GetCurrentSceneTransitionParams,
) (*GetCurrentSceneTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneTransitionParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentSceneTransitionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
