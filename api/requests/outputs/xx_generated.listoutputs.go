// This file has been automatically generated. Don't edit it.

package outputs

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
ListOutputsParams represents the params body for the "ListOutputs" request.
List existing outputs
Since 4.7.0.
*/
type ListOutputsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ListOutputs".
func (o *ListOutputsParams) GetSelfName() string {
	return "ListOutputs"
}

/*
ListOutputsResponse represents the response body for the "ListOutputs" request.
List existing outputs
Since v4.7.0.
*/
type ListOutputsResponse struct {
	requests.ResponseBasic

	// Outputs list
	Outputs []typedefs.Output `json:"outputs,omitempty"`
}

// ListOutputs sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ListOutputs(paramss ...*ListOutputsParams) (*ListOutputsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListOutputsParams{{}}
	}
	params := paramss[0]
	data := &ListOutputsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
