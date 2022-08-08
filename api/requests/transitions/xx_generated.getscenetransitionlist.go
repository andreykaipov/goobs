// This file has been automatically generated. Don't edit it.

package transitions

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSceneTransitionList request.
type GetSceneTransitionListParams struct{}

// Returns the associated request.
func (o *GetSceneTransitionListParams) GetRequestName() string {
	return "GetSceneTransitionList"
}

// Represents the response body for the GetSceneTransitionList request.
type GetSceneTransitionListResponse struct {
	// Kind of the current scene transition. Can be null
	CurrentSceneTransitionKind string `json:"currentSceneTransitionKind,omitempty"`

	// Name of the current scene transition. Can be null
	CurrentSceneTransitionName string `json:"currentSceneTransitionName,omitempty"`

	Transitions []*typedefs.Transition `json:"transitions,omitempty"`
}

// Gets an array of all scene transitions in OBS.
func (c *Client) GetSceneTransitionList(
	paramss ...*GetSceneTransitionListParams,
) (*GetSceneTransitionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneTransitionListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneTransitionListResponse{}
	return data, c.SendRequest(params, data)
}
