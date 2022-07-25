// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TriggerStudioModeTransitionParams represents the params body for the "TriggerStudioModeTransition" request.
Triggers the current scene transition. Same functionality as the `Transition` button in studio mode.
*/
type TriggerStudioModeTransitionParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "TriggerStudioModeTransition".
func (o *TriggerStudioModeTransitionParams) GetSelfName() string {
	return "TriggerStudioModeTransition"
}

/*
TriggerStudioModeTransitionResponse represents the response body for the "TriggerStudioModeTransition" request.
Triggers the current scene transition. Same functionality as the `Transition` button in studio mode.
*/
type TriggerStudioModeTransitionResponse struct {
	requests.ResponseBasic
}

// TriggerStudioModeTransition sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) TriggerStudioModeTransition(
	paramss ...*TriggerStudioModeTransitionParams,
) (*TriggerStudioModeTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerStudioModeTransitionParams{{}}
	}
	params := paramss[0]
	data := &TriggerStudioModeTransitionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
