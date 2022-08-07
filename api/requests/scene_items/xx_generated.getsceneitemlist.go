// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
GetSceneItemListParams represents the params body for the "GetSceneItemList" request.
Gets a list of all scene items in a scene.

Scenes only
*/
type GetSceneItemListParams struct {
	// Name of the scene to get the items of
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemListResponse represents the response body for the "GetSceneItemList" request.
Gets a list of all scene items in a scene.

Scenes only
*/
type GetSceneItemListResponse struct {
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`
}

// GetSceneItemList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemList(params *GetSceneItemListParams) (*GetSceneItemListResponse, error) {
	resp, err := c.SendRequest("GetSceneItemList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemListResponse), nil
}
