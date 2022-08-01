// This file has been automatically generated. Don't edit it.

package scenes

/*
CreateSceneParams represents the params body for the "CreateScene" request.
Creates a new scene in OBS.
*/
type CreateSceneParams struct {
	// Name for the new scene
	SceneName string `json:"sceneName,omitempty"`
}

/*
CreateSceneResponse represents the response body for the "CreateScene" request.
Creates a new scene in OBS.
*/
type CreateSceneResponse struct{}

// CreateScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateScene(params *CreateSceneParams) (*CreateSceneResponse, error) {
	resp, err := c.SendRequest("CreateScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateSceneResponse), nil
}
