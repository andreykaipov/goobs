// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentProfileParams represents the params body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentProfileParams
func (o *GetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentProfileParams
func (o *GetCurrentProfileParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentProfileParams
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

// GetMessageID returns the MessageID of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetError() string {
	return o.Error
}

// GetCurrentProfile sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentProfile(
	paramss ...*GetCurrentProfileParams,
) (*GetCurrentProfileResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProfileParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentProfile"
	data := &GetCurrentProfileResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
