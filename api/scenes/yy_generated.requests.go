// This file has been automatically generated. Don't edit it.

package scenes

type SetCurrentSceneParams struct {
	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}
type SetCurrentSceneResponse struct{}
type GetCurrentSceneParams struct{}
type GetCurrentSceneResponse struct {
	// Name of the currently active scene.
	Name string `json:"name"`
	// Ordered list of the current scene's source items.
	Sources []map[string]interface{} `json:"sources"`
}
type GetSceneListParams struct{}
type GetSceneListResponse struct {
	// Name of the currently active scene.
	CurrentScene string `json:"current-scene"`
	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for more information).
	Scenes []map[string]interface{} `json:"scenes"`
}
type ReorderSceneItemsParams struct {
	Items []struct {
		// Id of a specific scene item. Unique on a scene by scene basis.
		Id int `json:"id"`
		// Name of a scene item. Sufficiently unique if no scene items share sources within the scene.
		Name string `json:"name"`
	} `json:"items"`
	// Name of the scene to reorder (defaults to current).
	Scene string `json:"scene"`
}
type ReorderSceneItemsResponse struct{}
