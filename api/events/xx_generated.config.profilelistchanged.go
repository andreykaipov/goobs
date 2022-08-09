// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the ProfileListChanged event.

The profile list has changed.
*/
type ProfileListChanged struct {
	// Updated list of profiles
	Profiles []string `json:"profiles,omitempty"`
}
