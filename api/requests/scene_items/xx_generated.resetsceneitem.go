// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResetSceneItemParams represents the params body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemParams struct {
	requests.Params

	// Name of the source item.
	Item string `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of ResetSceneItemParams
func (o *ResetSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ResetSceneItemParams
func (o *ResetSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ResetSceneItemParams
func (o *ResetSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ResetSceneItemResponse represents the response body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetError() string {
	return o.Error
}

// ResetSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ResetSceneItem(params *ResetSceneItemParams) (*ResetSceneItemResponse, error) {
	params.RequestType = "ResetSceneItem"
	data := &ResetSceneItemResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
