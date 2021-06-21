// This file has been automatically generated. Don't edit it.

package events

/*
ProfileListChanged represents the event body for the "ProfileListChanged" event.
Since v4.0.0.
*/
type ProfileListChanged struct {
	EventBasic

	Profiles []*Profile `json:"profiles,omitempty"`
}

type Profile struct {
	// Profile name.
	Name string `json:"name,omitempty"`
}
