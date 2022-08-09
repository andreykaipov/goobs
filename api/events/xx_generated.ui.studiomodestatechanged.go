// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the StudioModeStateChanged event.

Studio mode has been enabled or disabled.
*/
type StudioModeStateChanged struct {
	// True == Enabled, False == Disabled
	StudioModeEnabled bool `json:"studioModeEnabled,omitempty"`
}
