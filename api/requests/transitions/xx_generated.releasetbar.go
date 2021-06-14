// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ReleaseTBarParams represents the params body for the "ReleaseTBar" request.
Release the T-Bar (like a user releasing their mouse button after moving it).
*YOU MUST CALL THIS if you called `SetTBarPosition` with the `release` parameter set to `false`.*
Since 4.9.0.
*/
type ReleaseTBarParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ReleaseTBar".
func (o *ReleaseTBarParams) GetSelfName() string {
	return "ReleaseTBar"
}

/*
ReleaseTBarResponse represents the response body for the "ReleaseTBar" request.
Release the T-Bar (like a user releasing their mouse button after moving it).
*YOU MUST CALL THIS if you called `SetTBarPosition` with the `release` parameter set to `false`.*
Since v4.9.0.
*/
type ReleaseTBarResponse struct {
	requests.ResponseBasic
}

// ReleaseTBar sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ReleaseTBar(paramss ...*ReleaseTBarParams) (*ReleaseTBarResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ReleaseTBarParams{{}}
	}
	params := paramss[0]
	data := &ReleaseTBarResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
