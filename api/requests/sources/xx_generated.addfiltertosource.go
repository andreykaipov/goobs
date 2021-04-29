// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AddFilterToSourceParams represents the params body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceParams struct {
	requests.Params

	// Name of the new filter
	FilterName string `json:"filterName"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Filter type
	FilterType string `json:"filterType"`

	// Name of the source on which the filter is added
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of AddFilterToSourceParams
func (o *AddFilterToSourceParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of AddFilterToSourceParams
func (o *AddFilterToSourceParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on AddFilterToSourceParams
func (o *AddFilterToSourceParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
AddFilterToSourceResponse represents the response body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of AddFilterToSourceResponse
func (o *AddFilterToSourceResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of AddFilterToSourceResponse
func (o *AddFilterToSourceResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of AddFilterToSourceResponse
func (o *AddFilterToSourceResponse) GetError() string {
	return o.Error
}

// AddFilterToSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) AddFilterToSource(
	params *AddFilterToSourceParams,
) (*AddFilterToSourceResponse, error) {
	params.RequestType = "AddFilterToSource"
	data := &AddFilterToSourceResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
