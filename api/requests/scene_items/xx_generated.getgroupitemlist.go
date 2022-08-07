// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
GetGroupItemListParams represents the params body for the "GetGroupItemList" request.
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood.

Groups only
*/
type GetGroupItemListParams struct {
	// Name of the group to get the items of
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetGroupItemListResponse represents the response body for the "GetGroupItemList" request.
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood.

Groups only
*/
type GetGroupItemListResponse struct {
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`
}

// GetGroupItemList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetGroupItemList(params *GetGroupItemListParams) (*GetGroupItemListResponse, error) {
	resp, err := c.SendRequest("GetGroupItemList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetGroupItemListResponse), nil
}
