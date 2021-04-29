// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DeleteSceneItemParams represents the params body for the "DeleteSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem.
*/
type DeleteSceneItemParams struct {
	requests.Params

	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	Scene string `json:"scene"`
}

// GetRequestType returns the RequestType of DeleteSceneItemParams
func (o *DeleteSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DeleteSceneItemParams
func (o *DeleteSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DeleteSceneItemParams
func (o *DeleteSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DeleteSceneItemResponse represents the response body for the "DeleteSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem.
*/
type DeleteSceneItemResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetError() string {
	return o.Error
}

// DeleteSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DeleteSceneItem(params *DeleteSceneItemParams) (*DeleteSceneItemResponse, error) {
	params.RequestType = "DeleteSceneItem"
	data := &DeleteSceneItemResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
