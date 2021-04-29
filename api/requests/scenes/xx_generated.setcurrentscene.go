// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneParams represents the params body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneParams struct {
	requests.Params

	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of SetCurrentSceneParams
func (o *SetCurrentSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentSceneParams
func (o *SetCurrentSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentSceneParams
func (o *SetCurrentSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentSceneResponse represents the response body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetError() string {
	return o.Error
}

// SetCurrentScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentScene(params *SetCurrentSceneParams) (*SetCurrentSceneResponse, error) {
	params.RequestType = "SetCurrentScene"
	data := &SetCurrentSceneResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
