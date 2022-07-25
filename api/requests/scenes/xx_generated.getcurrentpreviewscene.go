// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentPreviewSceneParams represents the params body for the "GetCurrentPreviewScene" request.
Gets the current preview scene.

Only available when studio mode is enabled.
*/
type GetCurrentPreviewSceneParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentPreviewScene".
func (o *GetCurrentPreviewSceneParams) GetSelfName() string {
	return "GetCurrentPreviewScene"
}

/*
GetCurrentPreviewSceneResponse represents the response body for the "GetCurrentPreviewScene" request.
Gets the current preview scene.

Only available when studio mode is enabled.
*/
type GetCurrentPreviewSceneResponse struct {
	requests.ResponseBasic

	// Current preview scene
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`
}

// GetCurrentPreviewScene sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentPreviewScene(
	paramss ...*GetCurrentPreviewSceneParams,
) (*GetCurrentPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentPreviewSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentPreviewSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
