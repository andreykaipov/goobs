// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DisableStudioModeParams represents the params body for the "DisableStudioMode" request.
Disables Studio Mode.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeParams struct {
	requests.ParamsBasic
}

// Name just returns "DisableStudioMode".
func (o *DisableStudioModeParams) Name() string {
	return "DisableStudioMode"
}

/*
DisableStudioModeResponse represents the response body for the "DisableStudioMode" request.
Disables Studio Mode.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeResponse struct {
	requests.ResponseBasic
}

// DisableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) DisableStudioMode(
	paramss ...*DisableStudioModeParams,
) (*DisableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*DisableStudioModeParams{{}}
	}
	params := paramss[0]
	data := &DisableStudioModeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
