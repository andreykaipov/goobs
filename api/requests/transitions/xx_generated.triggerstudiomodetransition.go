// This file has been automatically generated. Don't edit it.

package transitions

/*
TriggerStudioModeTransitionParams represents the params body for the "TriggerStudioModeTransition" request.
Triggers the current scene transition. Same functionality as the `Transition` button in studio mode.
*/
type TriggerStudioModeTransitionParams struct{}

/*
TriggerStudioModeTransitionResponse represents the response body for the "TriggerStudioModeTransition" request.
Triggers the current scene transition. Same functionality as the `Transition` button in studio mode.
*/
type TriggerStudioModeTransitionResponse struct{}

// TriggerStudioModeTransition sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) TriggerStudioModeTransition(
	paramss ...*TriggerStudioModeTransitionParams,
) (*TriggerStudioModeTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerStudioModeTransitionParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("TriggerStudioModeTransition", params)
	if err != nil {
		return nil, err
	}
	return resp.(*TriggerStudioModeTransitionResponse), nil
}
