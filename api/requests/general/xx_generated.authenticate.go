// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AuthenticateParams represents the params body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
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

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of AuthenticateResponse
func (o *AuthenticateResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of AuthenticateResponse
func (o *AuthenticateResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of AuthenticateResponse
func (o *AuthenticateResponse) GetError() string {
	return o.Error
}

// Authenticate sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) Authenticate(params *AuthenticateParams) (*AuthenticateResponse, error) {
	params.RequestType = "Authenticate"
	data := &AuthenticateResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
