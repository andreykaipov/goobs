// This file has been automatically generated. Don't edit it.

package scenes

/*
SetCurrentSceneParams represents the params body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneParams struct {
	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}

/*
SetCurrentSceneResponse represents the response body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneResponse struct{}

/*
GetCurrentSceneParams represents the params body for the "GetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneParams struct{}

/*
GetCurrentSceneResponse represents the response body for the "GetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneResponse struct {
	// Name of the currently active scene.
	Name string `json:"name"`

	// Ordered list of the current scene's source items.
	Sources []map[string]interface{} `json:"sources"`
}

/*
GetSceneListParams represents the params body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListParams struct{}

/*
GetSceneListResponse represents the response body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListResponse struct {
	// Name of the currently active scene.
	CurrentScene string `json:"current-scene"`

	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for
	// more information).
	Scenes []map[string]interface{} `json:"scenes"`
}

/*
ReorderSceneItemsParams represents the params body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
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

/*
ReorderSceneItemsResponse represents the response body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsResponse struct{}
