// This file has been automatically generated. Don't edit it.

package general

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetStatsParams represents the params body for the "GetStats" request.
Get OBS stats (almost the same info as provided in OBS' stats window)

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#GetStats.
*/
type GetStatsParams struct {
	requests.ParamsBasic
}

// Name just returns "GetStats".
func (o *GetStatsParams) Name() string {
	return "GetStats"
}

/*
GetStatsResponse represents the response body for the "GetStats" request.
Get OBS stats (almost the same info as provided in OBS' stats window)

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#GetStats.
*/
type GetStatsResponse struct {
	requests.ResponseBasic

	// OBS stats
	Stats []typedefs.OBSStats `json:"stats"`
}

// GetStats sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStats(paramss ...*GetStatsParams) (*GetStatsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStatsParams{{}}
	}
	params := paramss[0]
	data := &GetStatsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
