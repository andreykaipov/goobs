// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionPositionParams represents the params body for the "GetTransitionPosition" request.
Get the position of the current transition.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetTransitionPosition.
*/
type GetTransitionPositionParams struct {
	requests.ParamsBasic
}

// Name just returns "GetTransitionPosition".
func (o *GetTransitionPositionParams) Name() string {
	return "GetTransitionPosition"
}

/*
GetTransitionPositionResponse represents the response body for the "GetTransitionPosition" request.
Get the position of the current transition.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetTransitionPosition.
*/
type GetTransitionPositionResponse struct {
	requests.ResponseBasic

	// current transition position. This value will be between 0.0 and 1.0. Note: Transition returns
	// 1.0 when not active.
	Position float64 `json:"position"`
}

// GetTransitionPosition sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionPosition(
	paramss ...*GetTransitionPositionParams,
) (*GetTransitionPositionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionPositionParams{{}}
	}
	params := paramss[0]
	data := &GetTransitionPositionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
