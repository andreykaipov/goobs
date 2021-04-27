// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemParams struct {
	requests.Params

	// Name of the scene to copy the item from. Defaults to the current scene.
	FromScene string `json:"fromScene"`

	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene to create the item in. Defaults to the current scene.
	ToScene string `json:"toScene"`
}

// GetRequestType returns the RequestType of DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemResponse struct {
	requests.Response

	Item struct {
		// New item ID
		Id int `json:"id"`

		// New item name
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene where the new item was created
	Scene string `json:"scene"`
}

// GetMessageID returns the MessageID of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetError() string {
	return o.Error
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(
	params *DuplicateSceneItemParams,
) (*DuplicateSceneItemResponse, error) {
	params.RequestType = "DuplicateSceneItem"
	data := &DuplicateSceneItemResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
