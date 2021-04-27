// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ReorderSceneItemsParams represents the params body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsParams struct {
	requests.Params

	Items []struct {
		// Id of a specific scene item. Unique on a scene by scene basis.
		Id int `json:"id"`

		// Name of a scene item. Sufficiently unique if no scene items share sources within the
		// scene.
		Name string `json:"name"`
	} `json:"items"`

	// Name of the scene to reorder (defaults to current).
	Scene string `json:"scene"`
}

// GetRequestType returns the RequestType of ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ReorderSceneItemsResponse represents the response body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetError() string {
	return o.Error
}

// ReorderSceneItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSceneItems(
	params *ReorderSceneItemsParams,
) (*ReorderSceneItemsResponse, error) {
	params.RequestType = "ReorderSceneItems"
	data := &ReorderSceneItemsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
