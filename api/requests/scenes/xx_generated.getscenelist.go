// This file has been automatically generated. Don't edit it.

package scenes

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSceneList request.
type GetSceneListParams struct {
	// UUID of the canvas the scenes are in
	CanvasUuid *string `json:"canvasUuid,omitempty"`
}

func NewGetSceneListParams() *GetSceneListParams {
	return &GetSceneListParams{}
}
func (o *GetSceneListParams) WithCanvasUuid(x string) *GetSceneListParams {
	o.CanvasUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneListParams) GetRequestName() string {
	return "GetSceneList"
}

// Represents the response body for the GetSceneList request.
type GetSceneListResponse struct {
	_response

	// Current preview scene name. `null` if not in studio mode or non-main canvas
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`

	// Current preview scene UUID. `null` if not in studio mode or non-main canvas
	CurrentPreviewSceneUuid string `json:"currentPreviewSceneUuid,omitempty"`

	// Current program scene name. Can be `null` if non-main canvas or internal state desync
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`

	// Current program scene UUID. Can be `null` if non-main canvas or internal state desync
	CurrentProgramSceneUuid string `json:"currentProgramSceneUuid,omitempty"`

	// Array of scenes
	Scenes []*typedefs.Scene `json:"scenes,omitempty"`
}

// Gets an array of scenes in OBS.
func (c *Client) GetSceneList(paramss ...*GetSceneListParams) (*GetSceneListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneListResponse{}
	return data, c.client.SendRequest(params, data)
}
