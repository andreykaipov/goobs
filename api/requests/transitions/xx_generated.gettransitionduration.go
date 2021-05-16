// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationParams struct {
	requests.ParamsBasic
}

// Name just returns "GetTransitionDuration".
func (o *GetTransitionDurationParams) Name() string {
	return "GetTransitionDuration"
}

/*
GetTransitionDurationResponse represents the response body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationResponse struct {
	requests.ResponseBasic

	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}

// GetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionDuration(
	paramss ...*GetTransitionDurationParams,
) (*GetTransitionDurationResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionDurationParams{{}}
	}
	params := paramss[0]
	data := &GetTransitionDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
