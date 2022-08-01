// This file has been automatically generated. Don't edit it.

package config

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.
Switches to a scene collection.

Note: This will block until the collection has finished changing.
*/
type SetCurrentSceneCollectionParams struct {
	// Name of the scene collection to switch to
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.
Switches to a scene collection.

Note: This will block until the collection has finished changing.
*/
type SetCurrentSceneCollectionResponse struct{}

// SetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	resp, err := c.SendRequest("SetCurrentSceneCollection", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentSceneCollectionResponse), nil
}
