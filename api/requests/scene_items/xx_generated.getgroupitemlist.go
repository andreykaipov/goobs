// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetGroupItemListParams represents the params body for the "GetGroupItemList" request.
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood.

Groups only
*/
type GetGroupItemListParams struct {
	requests.ParamsBasic

	// Name of the group to get the items of
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetGroupItemList".
func (o *GetGroupItemListParams) GetSelfName() string {
	return "GetGroupItemList"
}

/*
GetGroupItemListResponse represents the response body for the "GetGroupItemList" request.
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood.

Groups only
*/
type GetGroupItemListResponse struct {
	requests.ResponseBasic

	// Array of scene items in the group
	SceneItems []interface{} `json:"sceneItems,omitempty"`
}

// GetGroupItemList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetGroupItemList(params *GetGroupItemListParams) (*GetGroupItemListResponse, error) {
	data := &GetGroupItemListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
