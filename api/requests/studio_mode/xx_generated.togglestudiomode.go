// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleStudioModeParams represents the params body for the "ToggleStudioMode" request.
Toggles Studio Mode (depending on the current state of studio mode).
Since 4.1.0.
*/
type ToggleStudioModeParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleStudioMode".
func (o *ToggleStudioModeParams) GetSelfName() string {
	return "ToggleStudioMode"
}

/*
ToggleStudioModeResponse represents the response body for the "ToggleStudioMode" request.
Toggles Studio Mode (depending on the current state of studio mode).
Since v4.1.0.
*/
type ToggleStudioModeResponse struct {
	requests.ResponseBasic
}

// ToggleStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) ToggleStudioMode(paramss ...*ToggleStudioModeParams) (*ToggleStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStudioModeParams{{}}
	}
	params := paramss[0]
	data := &ToggleStudioModeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
