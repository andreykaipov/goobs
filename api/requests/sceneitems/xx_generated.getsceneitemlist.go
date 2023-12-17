// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSceneItemList request.
type GetSceneItemListParams struct {
	// Name of the scene to get the items of
	SceneName *string `json:"sceneName,omitempty"`
}

func NewGetSceneItemListParams() *GetSceneItemListParams {
	return &GetSceneItemListParams{}
}
func (o *GetSceneItemListParams) WithSceneName(x string) *GetSceneItemListParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemListParams) GetRequestName() string {
	return "GetSceneItemList"
}

// Represents the response body for the GetSceneItemList request.
type GetSceneItemListResponse struct {
	// Array of scene items in the scene
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`
}

/*
Gets a list of all scene items in a scene.

Scenes only
*/
func (c *Client) GetSceneItemList(params *GetSceneItemListParams) (*GetSceneItemListResponse, error) {
	data := &GetSceneItemListResponse{}
	return data, c.SendRequest(params, data)
}
