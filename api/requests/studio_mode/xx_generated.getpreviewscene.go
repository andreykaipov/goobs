// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct {
	requests.ParamsBasic
}

// Name just returns "GetPreviewScene".
func (o *GetPreviewSceneParams) Name() string {
	return "GetPreviewScene"
}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	requests.ResponseBasic

	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []map[string]interface{} `json:"sources"`
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
	if err := c.WriteMessage(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
