// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVolumeParams represents the params body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of GetVolumeParams
func (o *GetVolumeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetVolumeParams
func (o *GetVolumeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetVolumeParams
func (o *GetVolumeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetVolumeResponse represents the response body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeResponse struct {
	requests.Response

	// Indicates whether the source is muted.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`

	// Volume of the source. Between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

// GetMessageID returns the MessageID of GetVolumeResponse
func (o *GetVolumeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetVolumeResponse
func (o *GetVolumeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetVolumeResponse
func (o *GetVolumeResponse) GetError() string {
	return o.Error
}

// GetVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetVolume(params *GetVolumeParams) (*GetVolumeResponse, error) {
	params.RequestType = "GetVolume"
	data := &GetVolumeResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
