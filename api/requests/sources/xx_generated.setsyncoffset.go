// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSyncOffsetParams represents the params body for the "SetSyncOffset" request.
Set the audio sync offset of a specified source.
Since 4.2.0.
*/
type SetSyncOffsetParams struct {
	requests.ParamsBasic

	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset,omitempty"`

	// Source name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "SetSyncOffset".
func (o *SetSyncOffsetParams) GetSelfName() string {
	return "SetSyncOffset"
}

/*
SetSyncOffsetResponse represents the response body for the "SetSyncOffset" request.
Set the audio sync offset of a specified source.
Since v4.2.0.
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
