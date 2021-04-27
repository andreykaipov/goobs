// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesListParams represents the params body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetSourcesListParams
func (o *GetSourcesListParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSourcesListParams
func (o *GetSourcesListParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSourcesListParams
func (o *GetSourcesListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourcesListResponse represents the response body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListResponse struct {
	requests.Response

	Sources []struct {
		// Unique source name
		Name string `json:"name"`

		// Source type. Value is one of the following: "input", "filter", "transition", "scene" or
		// "unknown"
		Type string `json:"type"`

		// Non-unique source internal type (a.k.a type id)
		TypeId string `json:"typeId"`
	} `json:"sources"`
}

// GetMessageID returns the MessageID of GetSourcesListResponse
func (o *GetSourcesListResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSourcesListResponse
func (o *GetSourcesListResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSourcesListResponse
func (o *GetSourcesListResponse) GetError() string {
	return o.Error
}

// GetSourcesList sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSourcesList(paramss ...*GetSourcesListParams) (*GetSourcesListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourcesListParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetSourcesList"
	data := &GetSourcesListResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
