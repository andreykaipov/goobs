// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSyncOffsetParams represents the params body for the "SetSyncOffset" request.
Set the audio sync offset of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetParams struct {
	requests.ParamsBasic

	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`

	// Source name.
	Source string `json:"source"`
}

// Name just returns "SetSyncOffset".
func (o *SetSyncOffsetParams) Name() string {
	return "SetSyncOffset"
}

/*
SetSyncOffsetResponse represents the response body for the "SetSyncOffset" request.
Set the audio sync offset of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetResponse struct {
	requests.ResponseBasic
}

// SetSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSyncOffset(params *SetSyncOffsetParams) (*SetSyncOffsetResponse, error) {
	data := &SetSyncOffsetResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
