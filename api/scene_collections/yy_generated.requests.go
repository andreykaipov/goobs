// This file has been automatically generated. Don't edit it.

package scenecollections

type SetCurrentSceneCollectionParams struct {
	// Name of the desired scene collection.
	ScName string `json:"sc-name"`
}
type SetCurrentSceneCollectionResponse struct{}
type GetCurrentSceneCollectionParams struct{}
type GetCurrentSceneCollectionResponse struct {
	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}
type ListSceneCollectionsParams struct{}
type ListSceneCollectionsResponse struct {
	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}
