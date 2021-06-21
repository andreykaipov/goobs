// This file has been automatically generated. Don't edit it.

package scenes

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetSceneListParams represents the params body for the "GetSceneList" request.
Get a list of scenes in the currently active profile.
Since 0.3.
*/
type GetSceneListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSceneList".
func (o *GetSceneListParams) GetSelfName() string {
	return "GetSceneList"
}

/*
GetSceneListResponse represents the response body for the "GetSceneList" request.
Get a list of scenes in the currently active profile.
Since v0.3.
*/
type GetSceneListResponse struct {
	requests.ResponseBasic

	// Name of the currently active scene.
	CurrentScene string `json:"current-scene,omitempty"`

	// Ordered list of the current profile's scenes (See [GetCurrentScene](#getcurrentscene) for more information).
	Scenes []typedefs.Scene `json:"scenes,omitempty"`
}

// GetSceneList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetSceneList(paramss ...*GetSceneListParams) (*GetSceneListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
