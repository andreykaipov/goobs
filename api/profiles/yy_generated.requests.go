// This file has been automatically generated. Don't edit it.

package profiles

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileParams struct {
	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileResponse struct{}

/*
GetCurrentProfileParams represents the params body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileParams struct{}

/*
GetCurrentProfileResponse represents the response body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileResponse struct {
	// Name of the currently active profile.
	ProfileName string `json:"profile-name"`
}

/*
ListProfilesParams represents the params body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesParams struct{}

/*
ListProfilesResponse represents the response body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesResponse struct {
	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}
