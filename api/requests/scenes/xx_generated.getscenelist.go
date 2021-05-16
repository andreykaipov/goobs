// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneListParams represents the params body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListParams struct {
	requests.ParamsBasic
}

// Name just returns "GetSceneList".
func (o *GetSceneListParams) Name() string {
	return "GetSceneList"
}

/*
GetSceneListResponse represents the response body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListResponse struct {
	requests.ResponseBasic

	// Name of the currently active scene.
	CurrentScene string `json:"current-scene"`

	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for
	// more information).
	Scenes []map[string]interface{} `json:"scenes"`
}

// GetSceneList sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
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
