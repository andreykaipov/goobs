// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the GetSceneCollectionList request.
type GetSceneCollectionListParams struct{}

// Returns the associated request.
func (o *GetSceneCollectionListParams) GetRequestName() string {
	return "GetSceneCollectionList"
}

// Represents the response body for the GetSceneCollectionList request.
type GetSceneCollectionListResponse struct {
	_response

	// The name of the current scene collection
	CurrentSceneCollectionName string `json:"currentSceneCollectionName,omitempty"`

	// Array of all available scene collections
	SceneCollections []string `json:"sceneCollections,omitempty"`
}

// Gets an array of all scene collections
func (c *Client) GetSceneCollectionList(
	paramss ...*GetSceneCollectionListParams,
) (*GetSceneCollectionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneCollectionListParams{{}}
	}
	params := paramss[0]
	data := &GetSceneCollectionListResponse{}
	return data, c.client.SendRequest(params, data)
}
