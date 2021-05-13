// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileParams struct {
	requests.ParamsBasic

	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}

// Name just returns "SetCurrentProfile".
func (o *SetCurrentProfileParams) Name() string {
	return "SetCurrentProfile"
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileResponse struct {
	requests.ResponseBasic
}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(
	params *SetCurrentProfileParams,
) (*SetCurrentProfileResponse, error) {
	data := &SetCurrentProfileResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
