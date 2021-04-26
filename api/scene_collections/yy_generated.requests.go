// This file has been automatically generated. Don't edit it.

package scenecollections

type SetCurrentSceneCollectionParams struct {
	ScName string `json:"sc-name"`
}
type SetCurrentSceneCollectionResponse struct{}
type GetCurrentSceneCollectionParams struct{}
type GetCurrentSceneCollectionResponse struct {
	ScName string `json:"sc-name"`
}
type ListSceneCollectionsParams struct{}
type ListSceneCollectionsResponse struct {
	SceneCollections []string `json:"scene-collections"`
}
