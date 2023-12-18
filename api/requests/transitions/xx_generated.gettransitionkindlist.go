// This file has been automatically generated. Don't edit it.

package transitions

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the GetTransitionKindList request.
type GetTransitionKindListParams struct{}

// Returns the associated request.
func (o *GetTransitionKindListParams) GetRequestName() string {
	return "GetTransitionKindList"
}

// Represents the response body for the GetTransitionKindList request.
type GetTransitionKindListResponse struct {
	api.ResponseCommon

	// Array of transition kinds
	TransitionKinds []string `json:"transitionKinds,omitempty"`
}

/*
Gets an array of all available transition kinds.

Similar to `GetInputKindList`
*/
func (c *Client) GetTransitionKindList(
	paramss ...*GetTransitionKindListParams,
) (*GetTransitionKindListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionKindListParams{{}}
	}
	params := paramss[0]
	data := &GetTransitionKindListResponse{}
	return data, c.SendRequest(params, data)
}
