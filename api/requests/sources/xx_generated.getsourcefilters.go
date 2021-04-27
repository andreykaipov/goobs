// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFiltersParams represents the params body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersParams struct {
	requests.Params

	// Source name
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of GetSourceFiltersParams
func (o *GetSourceFiltersParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSourceFiltersParams
func (o *GetSourceFiltersParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSourceFiltersParams
func (o *GetSourceFiltersParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourceFiltersResponse represents the response body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersResponse struct {
	requests.Response

	Filters []struct {
		// Filter name
		Name string `json:"name"`

		// Filter settings
		Settings map[string]interface{} `json:"settings"`

		// Filter type
		Type string `json:"type"`
	} `json:"filters"`
}

// GetMessageID returns the MessageID of GetSourceFiltersResponse
func (o *GetSourceFiltersResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSourceFiltersResponse
func (o *GetSourceFiltersResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSourceFiltersResponse
func (o *GetSourceFiltersResponse) GetError() string {
	return o.Error
}

// GetSourceFilters sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilters(
	params *GetSourceFiltersParams,
) (*GetSourceFiltersResponse, error) {
	params.RequestType = "GetSourceFilters"
	data := &GetSourceFiltersResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
