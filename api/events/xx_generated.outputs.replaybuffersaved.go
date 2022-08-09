// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the ReplayBufferSaved event.

The replay buffer has been saved.
*/
type ReplayBufferSaved struct {
	// Path of the saved replay file
	SavedReplayPath string `json:"savedReplayPath,omitempty"`
}
