// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneListParams represents the params body for the "GetSceneList" request.
Gets an array of all scenes in OBS.
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
Gets an array of all scenes in OBS.
*/
type GetSceneListResponse struct {
	requests.ResponseBasic

	// Current preview scene. `null` if not in studio mode
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`

	// Current program scene
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`

	// Array of scenes
	Scenes []interface{} `json:"scenes,omitempty"`
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
