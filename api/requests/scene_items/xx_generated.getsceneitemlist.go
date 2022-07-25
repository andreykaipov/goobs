// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemListParams represents the params body for the "GetSceneItemList" request.
Gets a list of all scene items in a scene.

Scenes only
*/
type GetSceneItemListParams struct {
	requests.ParamsBasic

	// Name of the scene to get the items of
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneItemList".
func (o *GetSceneItemListParams) GetSelfName() string {
	return "GetSceneItemList"
}

/*
GetSceneItemListResponse represents the response body for the "GetSceneItemList" request.
Gets a list of all scene items in a scene.

Scenes only
*/
type GetSceneItemListResponse struct {
	requests.ResponseBasic

	// Array of scene items in the scene
	SceneItems []interface{} `json:"sceneItems,omitempty"`
}

// GetSceneItemList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemList(params *GetSceneItemListParams) (*GetSceneItemListResponse, error) {
	data := &GetSceneItemListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
