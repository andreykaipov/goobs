// This file has been automatically generated. Don't edit it.

package outputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetOutputList request.
type GetOutputListParams struct{}

// Returns the associated request.
func (o *GetOutputListParams) GetRequestName() string {
	return "GetOutputList"
}

// Represents the response body for the GetOutputList request.
type GetOutputListResponse struct {
	_response

	// Array of outputs
	Outputs []*typedefs.Output `json:"outputs,omitempty"`
}

// Gets the list of available outputs.
func (c *Client) GetOutputList(paramss ...*GetOutputListParams) (*GetOutputListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetOutputListParams{{}}
	}
	params := paramss[0]
	data := &GetOutputListResponse{}
	return data, c.client.SendRequest(params, data)
}
