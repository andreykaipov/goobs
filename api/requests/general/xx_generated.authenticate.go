// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AuthenticateParams represents the params body for the "Authenticate" request.
Attempt to authenticate the client to the server.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateParams struct {
	requests.ParamsBasic

	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}

// Name just returns "Authenticate".
func (o *AuthenticateParams) Name() string {
	return "Authenticate"
}

/*
AuthenticateResponse represents the response body for the "Authenticate" request.
Attempt to authenticate the client to the server.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateResponse struct {
	requests.ResponseBasic
}

// Authenticate sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) Authenticate(params *AuthenticateParams) (*AuthenticateResponse, error) {
	data := &AuthenticateResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
