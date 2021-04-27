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

// GetRequestType returns the RequestType of SetCurrentProfileParams
func (o *SetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentProfileParams
func (o *SetCurrentProfileParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentProfileParams
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

// GetMessageID returns the MessageID of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetError() string {
	return o.Error
}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(
	params *SetCurrentProfileParams,
) (*SetCurrentProfileResponse, error) {
	params.RequestType = "SetCurrentProfile"
	data := &SetCurrentProfileResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
