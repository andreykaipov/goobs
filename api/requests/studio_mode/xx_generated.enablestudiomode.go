// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
EnableStudioModeParams represents the params body for the "EnableStudioMode" request.
Enables Studio Mode.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeParams struct {
	requests.ParamsBasic
}

// Name just returns "EnableStudioMode".
func (o *EnableStudioModeParams) Name() string {
	return "EnableStudioMode"
}

/*
EnableStudioModeResponse represents the response body for the "EnableStudioMode" request.
Enables Studio Mode.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeResponse struct {
	requests.ResponseBasic
}

// EnableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) EnableStudioMode(
	paramss ...*EnableStudioModeParams,
) (*EnableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*EnableStudioModeParams{{}}
	}
	params := paramss[0]
	data := &EnableStudioModeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
