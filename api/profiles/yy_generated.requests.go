// This file has been automatically generated. Don't edit it.

package profiles

type SetCurrentProfileParams struct {
	ProfileName string `json:"profile-name"`
}
type SetCurrentProfileResponse struct{}
type GetCurrentProfileParams struct{}
type GetCurrentProfileResponse struct {
	ProfileName string `json:"profile-name"`
}
type ListProfilesParams struct{}
type ListProfilesResponse struct {
	Profiles []map[string]interface{} `json:"profiles"`
}
