// This file has been automatically generated. Don't edit it.

package profiles

type SetCurrentProfileParams struct {
	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}
type SetCurrentProfileResponse struct{}
type GetCurrentProfileParams struct{}
type GetCurrentProfileResponse struct {
	// Name of the currently active profile.
	ProfileName string `json:"profile-name"`
}
type ListProfilesParams struct{}
type ListProfilesResponse struct {
	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}
