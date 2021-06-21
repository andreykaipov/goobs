// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemRenderParams represents the params body for the "SetSceneItemRender" request.
Show or hide a specified source item in a specified scene.
Since 0.3.
*/
type SetSceneItemRenderParams struct {
	requests.ParamsBasic

	// Scene Item id
	Item int `json:"item,omitempty"`

	// true = shown ; false = hidden
	Render bool `json:"render"`

	// Name of the scene the scene item belongs to. Defaults to the currently active scene.
	SceneName string `json:"scene-name,omitempty"`

	// Scene Item name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "SetSceneItemRender".
func (o *SetSceneItemRenderParams) GetSelfName() string {
	return "SetSceneItemRender"
}

/*
SetSceneItemRenderResponse represents the response body for the "SetSceneItemRender" request.
Show or hide a specified source item in a specified scene.
Since v0.3.
*/
type SetSceneItemRenderResponse struct {
	requests.ResponseBasic
}

// SetSceneItemRender sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemRender(params *SetSceneItemRenderParams) (*SetSceneItemRenderResponse, error) {
	data := &SetSceneItemRenderResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
