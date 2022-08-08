// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the MediaInputActionTriggered event.
type MediaInputActionTriggered struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Action performed on the input. See `ObsMediaInputAction` enum
	MediaAction string `json:"mediaAction,omitempty"`
}
