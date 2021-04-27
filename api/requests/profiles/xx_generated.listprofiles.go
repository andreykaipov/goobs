// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ListProfilesParams represents the params body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ListProfilesParams
func (o *ListProfilesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ListProfilesParams
func (o *ListProfilesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ListProfilesParams
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

// GetMessageID returns the MessageID of ListProfilesResponse
func (o *ListProfilesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ListProfilesResponse
func (o *ListProfilesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ListProfilesResponse
func (o *ListProfilesResponse) GetError() string {
	return o.Error
}

// ListProfiles sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ListProfiles(paramss ...*ListProfilesParams) (*ListProfilesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListProfilesParams{{}}
	}
	params := paramss[0]
	params.RequestType = "ListProfiles"
	data := &ListProfilesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
