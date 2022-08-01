// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemIdParams represents the params body for the "GetSceneItemId" request.
Searches a scene for a source, and returns its id.

Scenes and Groups
*/
type GetSceneItemIdParams struct {
	// Name of the scene or group to search in
	SceneName string `json:"sceneName,omitempty"`

	// Number of matches to skip during search. >= 0 means first forward. -1 means last (top) item
	SearchOffset float64 `json:"searchOffset,omitempty"`

	// Name of the source to find
	SourceName string `json:"sourceName,omitempty"`
}

/*
GetSceneItemIdResponse represents the response body for the "GetSceneItemId" request.
Searches a scene for a source, and returns its id.

Scenes and Groups
*/
type GetSceneItemIdResponse struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// GetSceneItemId sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemId(params *GetSceneItemIdParams) (*GetSceneItemIdResponse, error) {
	resp, err := c.SendRequest("GetSceneItemId", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemIdResponse), nil
}
