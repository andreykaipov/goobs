// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the TriggerStudioModeTransition request.
type TriggerStudioModeTransitionParams struct{}

// Returns the associated request.
func (o *TriggerStudioModeTransitionParams) GetRequestName() string {
	return "TriggerStudioModeTransition"
}

// Represents the response body for the TriggerStudioModeTransition request.
type TriggerStudioModeTransitionResponse struct {
	_response
}

// Triggers the current scene transition. Same functionality as the `Transition` button in studio mode.
func (c *Client) TriggerStudioModeTransition(
	paramss ...*TriggerStudioModeTransitionParams,
) (*TriggerStudioModeTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerStudioModeTransitionParams{{}}
	}
	params := paramss[0]
	data := &TriggerStudioModeTransitionResponse{}
	return data, c.client.SendRequest(params, data)
}
