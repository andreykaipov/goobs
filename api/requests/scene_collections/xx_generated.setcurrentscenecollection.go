// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionParams struct {
	requests.Params

	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}

// GetRequestType returns the RequestType of SetCurrentSceneCollectionParams
func (o *SetCurrentSceneCollectionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentSceneCollectionParams
func (o *SetCurrentSceneCollectionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentSceneCollectionParams
func (o *SetCurrentSceneCollectionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetCurrentSceneCollectionResponse
func (o *SetCurrentSceneCollectionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentSceneCollectionResponse
func (o *SetCurrentSceneCollectionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentSceneCollectionResponse
func (o *SetCurrentSceneCollectionResponse) GetError() string {
	return o.Error
}

// SetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	params.RequestType = "SetCurrentSceneCollection"
	data := &SetCurrentSceneCollectionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
