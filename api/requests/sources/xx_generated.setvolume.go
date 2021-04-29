// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetVolumeParams represents the params body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`

	// Desired volume. Must be between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

// GetRequestType returns the RequestType of SetVolumeParams
func (o *SetVolumeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetVolumeParams
func (o *SetVolumeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetVolumeParams
func (o *SetVolumeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetVolumeResponse represents the response body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetVolumeResponse
func (o *SetVolumeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetVolumeResponse
func (o *SetVolumeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetVolumeResponse
func (o *SetVolumeResponse) GetError() string {
	return o.Error
}

// SetVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetVolume(params *SetVolumeParams) (*SetVolumeResponse, error) {
	params.RequestType = "SetVolume"
	data := &SetVolumeResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
