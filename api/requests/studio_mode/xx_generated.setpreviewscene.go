// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneParams struct {
	requests.Params

	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of SetPreviewSceneParams
func (o *SetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetPreviewSceneParams
func (o *SetPreviewSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetPreviewSceneParams
func (o *SetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetError() string {
	return o.Error
}

// SetPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPreviewScene(params *SetPreviewSceneParams) (*SetPreviewSceneResponse, error) {
	params.RequestType = "SetPreviewScene"
	data := &SetPreviewSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
