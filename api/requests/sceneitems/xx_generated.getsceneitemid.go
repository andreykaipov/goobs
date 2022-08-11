// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemId request.
type GetSceneItemIdParams struct {
	// Name of the scene or group to search in
	SceneName string `json:"sceneName,omitempty"`

	// Number of matches to skip during search. >= 0 means first forward. -1 means last (top) item
	SearchOffset float64 `json:"searchOffset,omitempty"`

	// Name of the source to find
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneItemIdParams) GetRequestName() string {
	return "GetSceneItemId"
}

// Represents the response body for the GetSceneItemId request.
type GetSceneItemIdResponse struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

/*
Searches a scene for a source, and returns its id.

Scenes and Groups
*/
func (c *Client) GetSceneItemId(params *GetSceneItemIdParams) (*GetSceneItemIdResponse, error) {
	data := &GetSceneItemIdResponse{}
	return data, c.SendRequest(params, data)
}
