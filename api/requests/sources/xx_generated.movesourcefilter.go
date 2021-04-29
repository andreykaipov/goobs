// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
MoveSourceFilterParams represents the params body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterParams struct {
	requests.Params

	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or
	// "bottom".
	MovementType string `json:"movementType"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of MoveSourceFilterParams
func (o *MoveSourceFilterParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of MoveSourceFilterParams
func (o *MoveSourceFilterParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on MoveSourceFilterParams
func (o *MoveSourceFilterParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
MoveSourceFilterResponse represents the response body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of MoveSourceFilterResponse
func (o *MoveSourceFilterResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of MoveSourceFilterResponse
func (o *MoveSourceFilterResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of MoveSourceFilterResponse
func (o *MoveSourceFilterResponse) GetError() string {
	return o.Error
}

// MoveSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) MoveSourceFilter(
	params *MoveSourceFilterParams,
) (*MoveSourceFilterResponse, error) {
	params.RequestType = "MoveSourceFilter"
	data := &MoveSourceFilterResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
