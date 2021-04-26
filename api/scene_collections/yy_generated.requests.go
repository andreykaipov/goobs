// This file has been automatically generated. Don't edit it.

package scenecollections

// SetCurrentSceneCollectionParams contains the request body for the
// [SetCurrentSceneCollection](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection)
// request.
type SetCurrentSceneCollectionParams struct {
	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}

// SetCurrentSceneCollectionResponse contains the request body for the
// [SetCurrentSceneCollection](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentSceneCollection)
// request.
type SetCurrentSceneCollectionResponse struct{}

// GetCurrentSceneCollectionParams contains the request body for the
// [GetCurrentSceneCollection](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection)
// request.
type GetCurrentSceneCollectionParams struct{}

// GetCurrentSceneCollectionResponse contains the request body for the
// [GetCurrentSceneCollection](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentSceneCollection)
// request.
type GetCurrentSceneCollectionResponse struct {
	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}

// ListSceneCollectionsParams contains the request body for the
// [ListSceneCollections](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections)
// request.
type ListSceneCollectionsParams struct{}

// ListSceneCollectionsResponse contains the request body for the
// [ListSceneCollections](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections)
// request.
type ListSceneCollectionsResponse struct {
	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}
