// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetGroupItemList request.
type GetGroupItemListParams struct {
	// Name of the group to get the items of
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetGroupItemListParams) GetRequestName() string {
	return "GetGroupItemList"
}

// Represents the response body for the GetGroupItemList request.
type GetGroupItemListResponse struct {
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`
}

/*
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood.

Groups only
*/
func (c *Client) GetGroupItemList(params *GetGroupItemListParams) (*GetGroupItemListResponse, error) {
	data := &GetGroupItemListResponse{}
	return data, c.SendRequest(params, data)
}
