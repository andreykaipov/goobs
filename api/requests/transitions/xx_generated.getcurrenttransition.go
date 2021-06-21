// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentTransitionParams represents the params body for the "GetCurrentTransition" request.
Get the name of the currently selected transition in the frontend's dropdown menu.
Since 0.3.
*/
type GetCurrentTransitionParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentTransition".
func (o *GetCurrentTransitionParams) GetSelfName() string {
	return "GetCurrentTransition"
}

/*
GetCurrentTransitionResponse represents the response body for the "GetCurrentTransition" request.
Get the name of the currently selected transition in the frontend's dropdown menu.
Since v0.3.
*/
type GetCurrentTransitionResponse struct {
	requests.ResponseBasic

	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration,omitempty"`

	// Name of the selected transition.
	Name string `json:"name,omitempty"`
}

// GetCurrentTransition sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentTransition(paramss ...*GetCurrentTransitionParams) (*GetCurrentTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentTransitionParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentTransitionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
