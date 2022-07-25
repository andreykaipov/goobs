// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemIdParams represents the params body for the "GetSceneItemId" request.
Searches a scene for a source, and returns its id.

Scenes and Groups
*/
type GetSceneItemIdParams struct {
	requests.ParamsBasic

	// Name of the scene or group to search in
	SceneName string `json:"sceneName,omitempty"`

	// Number of matches to skip during search. >= 0 means first forward. -1 means last (top) item
	SearchOffset float64 `json:"searchOffset,omitempty"`

	// Name of the source to find
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSceneItemId".
func (o *GetSceneItemIdParams) GetSelfName() string {
	return "GetSceneItemId"
}

/*
GetSceneItemIdResponse represents the response body for the "GetSceneItemId" request.
Searches a scene for a source, and returns its id.

Scenes and Groups
*/
type GetSceneItemIdResponse struct {
	requests.ResponseBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// GetSceneItemId sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemId(params *GetSceneItemIdParams) (*GetSceneItemIdResponse, error) {
	data := &GetSceneItemIdResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
