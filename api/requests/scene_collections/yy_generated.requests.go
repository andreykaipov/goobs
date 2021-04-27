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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsResponse struct {
	requests.Response

	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}

// GetMessageID returns the MessageID of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetError() string {
	return o.Error
}

// ListSceneCollections sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) ListSceneCollections(
	paramss ...*ListSceneCollectionsParams,
) (*ListSceneCollectionsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListSceneCollectionsParams{{}}
	}
	params := paramss[0]
	params.RequestType = "ListSceneCollections"
	data := &ListSceneCollectionsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}