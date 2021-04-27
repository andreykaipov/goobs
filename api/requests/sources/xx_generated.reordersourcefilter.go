// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ReorderSourceFilterParams represents the params body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterParams struct {
	requests.Params

	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of ReorderSourceFilterParams
func (o *ReorderSourceFilterParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ReorderSourceFilterParams
func (o *ReorderSourceFilterParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ReorderSourceFilterParams
func (o *ReorderSourceFilterParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ReorderSourceFilterResponse represents the response body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ReorderSourceFilterResponse
func (o *ReorderSourceFilterResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ReorderSourceFilterResponse
func (o *ReorderSourceFilterResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ReorderSourceFilterResponse
func (o *ReorderSourceFilterResponse) GetError() string {
	return o.Error
}

// ReorderSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSourceFilter(
	params *ReorderSourceFilterParams,
) (*ReorderSourceFilterResponse, error) {
	params.RequestType = "ReorderSourceFilter"
	data := &ReorderSourceFilterResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
