// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneTransitionListParams represents the params body for the "GetSceneTransitionList" request.
Gets an array of all scene transitions in OBS.
*/
type GetSceneTransitionListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSceneTransitionList".
func (o *GetSceneTransitionListParams) GetSelfName() string {
	return "GetSceneTransitionList"
}

/*
GetSceneTransitionListResponse represents the response body for the "GetSceneTransitionList" request.
Gets an array of all scene transitions in OBS.
*/
type GetSceneTransitionListResponse struct {
	requests.ResponseBasic

	// Kind of the current scene transition. Can be null
	CurrentSceneTransitionKind string `json:"currentSceneTransitionKind,omitempty"`

	// Name of the current scene transition. Can be null
	CurrentSceneTransitionName string `json:"currentSceneTransitionName,omitempty"`

	// Array of transitions
	Transitions []interface{} `json:"transitions,omitempty"`
}

// GetSceneTransitionList sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetSceneTransitionList(
	paramss ...*GetSceneTransitionListParams,
) (*GetSceneTransitionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneTransitionListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneTransitionListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
