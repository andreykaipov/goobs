// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSyncOffsetParams represents the params body for the "GetSyncOffset" request.
Get the audio sync offset of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// Name just returns "GetSyncOffset".
func (o *GetSyncOffsetParams) Name() string {
	return "GetSyncOffset"
}

/*
GetSyncOffsetResponse represents the response body for the "GetSyncOffset" request.
Get the audio sync offset of a specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetResponse struct {
	requests.ResponseBasic

	// Source name.
	Name string `json:"name"`

	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}

// GetSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSyncOffset(params *GetSyncOffsetParams) (*GetSyncOffsetResponse, error) {
	data := &GetSyncOffsetResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
