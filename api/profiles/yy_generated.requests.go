// This file has been automatically generated. Don't edit it.

package profiles

// SetCurrentProfileParams contains the request body for the
// [SetCurrentProfile](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile)
// request.
type SetCurrentProfileParams struct {
	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}

// SetCurrentProfileResponse contains the request body for the
// [SetCurrentProfile](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile)
// request.
type SetCurrentProfileResponse struct{}

// GetCurrentProfileParams contains the request body for the
// [GetCurrentProfile](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile)
// request.
type GetCurrentProfileParams struct{}

// GetCurrentProfileResponse contains the request body for the
// [GetCurrentProfile](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile)
// request.
type GetCurrentProfileResponse struct {
	// Name of the currently active profile.
	ProfileName string `json:"profile-name"`
}

// ListProfilesParams contains the request body for the
// [ListProfiles](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles)
// request.
type ListProfilesParams struct{}

// ListProfilesResponse contains the request body for the
// [ListProfiles](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles)
// request.
type ListProfilesResponse struct {
	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}
