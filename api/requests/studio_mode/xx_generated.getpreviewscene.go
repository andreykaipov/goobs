// This file has been automatically generated. Don't edit it.

package studiomode

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.
Get the name of the currently previewed scene and its list of sources.
Will return an `error` if Studio Mode is not enabled.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetPreviewScene".
func (o *GetPreviewSceneParams) GetSelfName() string {
	return "GetPreviewScene"
}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.
Get the name of the currently previewed scene and its list of sources.
Will return an `error` if Studio Mode is not enabled.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	requests.ResponseBasic

	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []typedefs.SceneItem `json:"sources"`
}

// GetPreviewScene sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetPreviewScene(
	paramss ...*GetPreviewSceneParams,
) (*GetPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetPreviewSceneParams{{}}
	}
	params := paramss[0]
	data := &GetPreviewSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
