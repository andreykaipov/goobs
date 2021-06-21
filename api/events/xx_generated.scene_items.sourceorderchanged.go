// This file has been automatically generated. Don't edit it.

package events

/*
SourceOrderChanged represents the event body for the "SourceOrderChanged" event.
Since v4.0.0.
*/
type SourceOrderChanged struct {
	EventBasic

	SceneItems []*SceneItem `json:"scene-items,omitempty"`

	// Name of the scene where items have been reordered.
	SceneName string `json:"scene-name,omitempty"`
}

type SceneItem struct {
	// Scene item unique ID
	ItemId int `json:"item-id,omitempty"`

	// Item source name
	SourceName string `json:"source-name,omitempty"`
}
