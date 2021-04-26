// This file has been automatically generated. Don't edit it.

package scenes

// SetCurrentSceneParams contains the request body for the
// [SetCurrentScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene)
// request.
type SetCurrentSceneParams struct {
	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}

// SetCurrentSceneResponse contains the request body for the
// [SetCurrentScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene)
// request.
type SetCurrentSceneResponse struct{}

// GetCurrentSceneResponse contains the request body for the
// [GetCurrentScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene)
// request.
type GetCurrentSceneResponse struct {
	// Name of the currently active scene.
	Name string `json:"name"`
	// Ordered list of the current scene's source items.
	Sources []map[string]interface{} `json:"sources"`
}

// GetCurrentSceneParams contains the request body for the
// [GetCurrentScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene)
// request.
type GetCurrentSceneParams struct{}

// GetSceneListParams contains the request body for the
// [GetSceneList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList)
// request.
type GetSceneListParams struct{}

// GetSceneListResponse contains the request body for the
// [GetSceneList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList)
// request.
type GetSceneListResponse struct {
	// Name of the currently active scene.
	CurrentScene string `json:"current-scene"`
	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for
	// more information).
	Scenes []map[string]interface{} `json:"scenes"`
}

// ReorderSceneItemsResponse contains the request body for the
// [ReorderSceneItems](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems)
// request.
type ReorderSceneItemsResponse struct{}

// ReorderSceneItemsParams contains the request body for the
// [ReorderSceneItems](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems)
// request.
type ReorderSceneItemsParams struct {
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
