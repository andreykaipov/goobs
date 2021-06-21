// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.
Tells the client if authentication is required. If so, returns authentication parameters `challenge`
and `salt` (see "Authentication" for more information).
Since 0.3.
*/
type GetAuthRequiredParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetAuthRequired".
func (o *GetAuthRequiredParams) GetSelfName() string {
	return "GetAuthRequired"
}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.
Tells the client if authentication is required. If so, returns authentication parameters `challenge`
and `salt` (see "Authentication" for more information).
Since v0.3.
*/
type GetAuthRequiredResponse struct {
	requests.ResponseBasic

	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge,omitempty"`

	Salt string `json:"salt,omitempty"`
}

// GetAuthRequired sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetAuthRequired(paramss ...*GetAuthRequiredParams) (*GetAuthRequiredResponse, error) {
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
