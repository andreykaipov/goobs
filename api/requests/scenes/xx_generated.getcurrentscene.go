// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentSceneParams represents the params body for the "GetCurrentScene" request.
Get the current scene's name and source items.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneParams struct {
	requests.ParamsBasic
}

// Name just returns "GetCurrentScene".
func (o *GetCurrentSceneParams) Name() string {
	return "GetCurrentScene"
}

/*
GetCurrentSceneResponse represents the response body for the "GetCurrentScene" request.
Get the current scene's name and source items.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneResponse struct {
	requests.ResponseBasic

	// Name of the currently active scene.
	Name string `json:"name"`

	// Ordered list of the current scene's source items.
	Sources []map[string]interface{} `json:"sources"`
}

// GetCurrentScene sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentScene(
	paramss ...*GetCurrentSceneParams,
) (*GetCurrentSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
