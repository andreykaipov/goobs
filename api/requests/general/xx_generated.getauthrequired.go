// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.
Tells the client if authentication is required. If so, returns authentication parameters `challenge`
and `salt` (see "Authentication" for more information).

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredParams struct {
	requests.ParamsBasic
}

// Name just returns "GetAuthRequired".
func (o *GetAuthRequiredParams) Name() string {
	return "GetAuthRequired"
}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.
Tells the client if authentication is required. If so, returns authentication parameters `challenge`
and `salt` (see "Authentication" for more information).

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredResponse struct {
	requests.ResponseBasic

	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

// GetAuthRequired sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetAuthRequired(
	paramss ...*GetAuthRequiredParams,
) (*GetAuthRequiredResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetAuthRequiredParams{{}}
	}
	params := paramss[0]
	data := &GetAuthRequiredResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
