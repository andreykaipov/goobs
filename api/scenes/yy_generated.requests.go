// This file has been automatically generated. Don't edit it.

package scenes

type SetCurrentSceneParams struct {
	SceneName string `json:"scene-name"`
}
type SetCurrentSceneResponse struct{}
type GetCurrentSceneParams struct{}
type GetCurrentSceneResponse struct {
	Name    string                   `json:"name"`
	Sources []map[string]interface{} `json:"sources"`
}
type GetSceneListParams struct{}
type GetSceneListResponse struct {
	CurrentScene string                   `json:"current-scene"`
	Scenes       []map[string]interface{} `json:"scenes"`
}
type ReorderSceneItemsParams struct {
	Items []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
	Scene string `json:"scene"`
}
type ReorderSceneItemsResponse struct{}
