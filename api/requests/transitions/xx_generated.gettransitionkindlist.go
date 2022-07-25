// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionKindListParams represents the params body for the "GetTransitionKindList" request.
Gets an array of all available transition kinds.

Similar to `GetInputKindList`
*/
type GetTransitionKindListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetTransitionKindList".
func (o *GetTransitionKindListParams) GetSelfName() string {
	return "GetTransitionKindList"
}

/*
GetTransitionKindListResponse represents the response body for the "GetTransitionKindList" request.
Gets an array of all available transition kinds.

Similar to `GetInputKindList`
*/
type GetTransitionKindListResponse struct {
	requests.ResponseBasic

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
	data := &GetTransitionKindListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
