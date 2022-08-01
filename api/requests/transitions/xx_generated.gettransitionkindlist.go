// This file has been automatically generated. Don't edit it.

package transitions

/*
GetTransitionKindListParams represents the params body for the "GetTransitionKindList" request.
Gets an array of all available transition kinds.

Similar to `GetInputKindList`
*/
type GetTransitionKindListParams struct{}

/*
GetTransitionKindListResponse represents the response body for the "GetTransitionKindList" request.
Gets an array of all available transition kinds.

Similar to `GetInputKindList`
*/
type GetTransitionKindListResponse struct {
	// Array of transition kinds
	TransitionKinds []string `json:"transitionKinds,omitempty"`
}

// GetTransitionKindList sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionKindList(
	paramss ...*GetTransitionKindListParams,
) (*GetTransitionKindListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionKindListParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetTransitionKindList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetTransitionKindListResponse), nil
}
