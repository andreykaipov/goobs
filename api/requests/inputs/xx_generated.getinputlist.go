// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetInputList request.
type GetInputListParams struct {
	// Restrict the array to only inputs of the specified kind
	InputKind string `json:"inputKind,omitempty"`
}

// Returns the associated request.
func (o *GetInputListParams) GetRequestName() string {
	return "GetInputList"
}

// Represents the response body for the GetInputList request.
type GetInputListResponse struct {
	Inputs []*typedefs.Input `json:"inputs,omitempty"`
}

// Gets an array of all inputs in OBS.
func (c *Client) GetInputList(paramss ...*GetInputListParams) (*GetInputListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputListParams{{}}
	}
	params := paramss[0]
	data := &GetInputListResponse{}
	return data, c.SendRequest(params, data)
}
