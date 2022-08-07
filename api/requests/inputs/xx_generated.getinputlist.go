// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
GetInputListParams represents the params body for the "GetInputList" request.
Gets an array of all inputs in OBS.
*/
type GetInputListParams struct {
	// Restrict the array to only inputs of the specified kind
	InputKind string `json:"inputKind,omitempty"`
}

/*
GetInputListResponse represents the response body for the "GetInputList" request.
Gets an array of all inputs in OBS.
*/
type GetInputListResponse struct {
	Inputs []*typedefs.Input `json:"inputs,omitempty"`
}

// GetInputList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetInputList(paramss ...*GetInputListParams) (*GetInputListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputListParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetInputList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputListResponse), nil
}
