// This file has been automatically generated. Don't edit it.

package scenecollections

import api "github.com/andreykaipov/goobs/api"

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionParams struct {
	api.Params

	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}

func (o *SetCurrentSceneCollectionParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetCurrentSceneCollectionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionResponse struct {
	api.Response
}

func (o *SetCurrentSceneCollectionResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetCurrentSceneCollectionResponse) GetStatus() string {
	return o.Status
}
func (o *SetCurrentSceneCollectionResponse) GetError() string {
	return o.Error
}

/*
GetCurrentSceneCollectionParams represents the params body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionParams struct {
	api.Params
}

func (o *GetCurrentSceneCollectionParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetCurrentSceneCollectionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentSceneCollectionResponse represents the response body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionResponse struct {
	api.Response

	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}

func (o *GetCurrentSceneCollectionResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetCurrentSceneCollectionResponse) GetStatus() string {
	return o.Status
}
func (o *GetCurrentSceneCollectionResponse) GetError() string {
	return o.Error
}

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsParams struct {
	api.Params
}

func (o *ListSceneCollectionsParams) GetRequestType() string {
	return o.RequestType
}
func (o *ListSceneCollectionsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsResponse struct {
	api.Response

	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}

func (o *ListSceneCollectionsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ListSceneCollectionsResponse) GetStatus() string {
	return o.Status
}
func (o *ListSceneCollectionsResponse) GetError() string {
	return o.Error
}
