// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentSceneCollectionParams represents the params body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentSceneCollectionParams
func (o *GetCurrentSceneCollectionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentSceneCollectionParams
func (o *GetCurrentSceneCollectionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentSceneCollectionParams
func (o *GetCurrentSceneCollectionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentSceneCollectionResponse represents the response body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionResponse struct {
	requests.Response

	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}

// GetMessageID returns the MessageID of GetCurrentSceneCollectionResponse
func (o *GetCurrentSceneCollectionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentSceneCollectionResponse
func (o *GetCurrentSceneCollectionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentSceneCollectionResponse
func (o *GetCurrentSceneCollectionResponse) GetError() string {
	return o.Error
}

// GetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentSceneCollection(
	paramss ...*GetCurrentSceneCollectionParams,
) (*GetCurrentSceneCollectionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneCollectionParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentSceneCollection"
	data := &GetCurrentSceneCollectionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
