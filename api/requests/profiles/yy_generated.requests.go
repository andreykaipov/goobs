// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileParams struct {
	requests.Params

	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}

func (o *SetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetCurrentProfileParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileResponse struct {
	requests.Response
}

func (o *SetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetCurrentProfileResponse) GetStatus() string {
	return o.Status
}
func (o *SetCurrentProfileResponse) GetError() string {
	return o.Error
}

/*
GetCurrentProfileParams represents the params body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileParams struct {
	requests.Params
}

func (o *GetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetCurrentProfileParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentProfileResponse represents the response body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileResponse struct {
	requests.Response

	// Name of the currently active profile.
	ProfileName string `json:"profile-name"`
}

func (o *GetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetCurrentProfileResponse) GetStatus() string {
	return o.Status
}
func (o *GetCurrentProfileResponse) GetError() string {
	return o.Error
}

/*
ListProfilesParams represents the params body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesParams struct {
	requests.Params
}

func (o *ListProfilesParams) GetRequestType() string {
	return o.RequestType
}
func (o *ListProfilesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ListProfilesResponse represents the response body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesResponse struct {
	requests.Response

	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}

func (o *ListProfilesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ListProfilesResponse) GetStatus() string {
	return o.Status
}
func (o *ListProfilesResponse) GetError() string {
	return o.Error
}
