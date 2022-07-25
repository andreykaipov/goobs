// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneCollectionListParams represents the params body for the "GetSceneCollectionList" request.
Gets an array of all scene collections
*/
type GetSceneCollectionListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSceneCollectionList".
func (o *GetSceneCollectionListParams) GetSelfName() string {
	return "GetSceneCollectionList"
}

/*
GetSceneCollectionListResponse represents the response body for the "GetSceneCollectionList" request.
Gets an array of all scene collections
*/
type GetSceneCollectionListResponse struct {
	requests.ResponseBasic

	// The name of the current scene collection
	CurrentSceneCollectionName string `json:"currentSceneCollectionName,omitempty"`

	// Array of all available scene collections
	SceneCollections []string `json:"sceneCollections,omitempty"`
}

// GetSceneCollectionList sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetSceneCollectionList(
	paramss ...*GetSceneCollectionListParams,
) (*GetSceneCollectionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneCollectionListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneCollectionListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
