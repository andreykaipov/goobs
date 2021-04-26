// This file has been automatically generated. Don't edit it.

package scenecollections

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionParams struct {
	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection.
*/
type SetCurrentSceneCollectionResponse struct{}

/*
GetCurrentSceneCollectionParams represents the params body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionParams struct{}

/*
GetCurrentSceneCollectionResponse represents the response body for the "GetCurrentSceneCollection" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionResponse struct {
	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsParams struct{}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsResponse struct {
	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}
