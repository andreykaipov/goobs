// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSyncOffsetParams represents the params body for the "GetSyncOffset" request.
Get the audio sync offset of a specified source.
Since 4.2.0.
*/
type GetSyncOffsetParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "GetSyncOffset".
func (o *GetSyncOffsetParams) GetSelfName() string {
	return "GetSyncOffset"
}

/*
GetSyncOffsetResponse represents the response body for the "GetSyncOffset" request.
Get the audio sync offset of a specified source.
Since v4.2.0.
*/
type GetSyncOffsetResponse struct {
	requests.ResponseBasic

	// Source name.
	Name string `json:"name,omitempty"`

	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset,omitempty"`
}

// GetSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSyncOffset(params *GetSyncOffsetParams) (*GetSyncOffsetResponse, error) {
	data := &GetSyncOffsetResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
