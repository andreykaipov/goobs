// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveFilterFromSourceParams represents the params body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceParams struct {
	requests.Params

	// Name of the filter to remove
	FilterName string `json:"filterName"`

	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of RemoveFilterFromSourceParams
func (o *RemoveFilterFromSourceParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of RemoveFilterFromSourceParams
func (o *RemoveFilterFromSourceParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on RemoveFilterFromSourceParams
func (o *RemoveFilterFromSourceParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
RemoveFilterFromSourceResponse represents the response body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of RemoveFilterFromSourceResponse
func (o *RemoveFilterFromSourceResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of RemoveFilterFromSourceResponse
func (o *RemoveFilterFromSourceResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of RemoveFilterFromSourceResponse
func (o *RemoveFilterFromSourceResponse) GetError() string {
	return o.Error
}

// RemoveFilterFromSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveFilterFromSource(
	params *RemoveFilterFromSourceParams,
) (*RemoveFilterFromSourceResponse, error) {
	params.RequestType = "RemoveFilterFromSource"
	data := &RemoveFilterFromSourceResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
