// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetGroupSceneItemList request.
type GetGroupSceneItemListParams struct {
	// Name of the group to get the items of
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the group to get the items of
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetGroupSceneItemListParams() *GetGroupSceneItemListParams {
	return &GetGroupSceneItemListParams{}
}
func (o *GetGroupSceneItemListParams) WithSceneName(x string) *GetGroupSceneItemListParams {
	o.SceneName = &x
	return o
}
func (o *GetGroupSceneItemListParams) WithSceneUuid(x string) *GetGroupSceneItemListParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetGroupSceneItemListParams) GetRequestName() string {
	return "GetGroupSceneItemList"
}

// Represents the response body for the GetGroupSceneItemList request.
type GetGroupSceneItemListResponse struct {
	_response

	// Array of scene items in the group
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`
}

/*
Basically GetSceneItemList, but for groups.

Using groups at all in OBS is discouraged, as they are very broken under the hood. Please use nested scenes instead.

Groups only
*/
func (c *Client) GetGroupSceneItemList(
	paramss ...*GetGroupSceneItemListParams,
) (*GetGroupSceneItemListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetGroupSceneItemListParams{{}}
	}
	params := paramss[0]
	data := &GetGroupSceneItemListResponse{}
	return data, c.client.SendRequest(params, data)
}
