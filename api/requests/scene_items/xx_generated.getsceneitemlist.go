// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemListParams represents the params body for the "GetSceneItemList" request.
Get a list of all scene items in a scene.
Since 4.9.0.
*/
type GetSceneItemListParams struct {
	requests.ParamsBasic

	// Name of the scene to get the list of scene items from. Defaults to the current scene if not specified.
	SceneName string `json:"sceneName"`
}

// GetSelfName just returns "GetSceneItemList".
func (o *GetSceneItemListParams) GetSelfName() string {
	return "GetSceneItemList"
}

/*
GetSceneItemListResponse represents the response body for the "GetSceneItemList" request.
Get a list of all scene items in a scene.
Since v4.9.0.
*/
type GetSceneItemListResponse struct {
	requests.ResponseBasic

	SceneItems []SceneItem `json:"sceneItems"`

	// Name of the requested (or current) scene
	SceneName string `json:"sceneName"`
}

type SceneItem struct {
	// Unique item id of the source item
	ItemId int `json:"itemId"`

	// ID if the scene item's source. For example `vlc_source` or `image_source`
	SourceKind string `json:"sourceKind"`

	// Name of the scene item's source
	SourceName string `json:"sourceName"`

	// Type of the scene item's source. Either `input`, `group`, or `scene`
	SourceType string `json:"sourceType"`
}

// GetSceneItemList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemList(params *GetSceneItemListParams) (*GetSceneItemListResponse, error) {
	data := &GetSceneItemListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
