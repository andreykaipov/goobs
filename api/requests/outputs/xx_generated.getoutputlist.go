// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetOutputList request.
type GetOutputListParams struct{}

// Returns the associated request.
func (o *GetOutputListParams) GetRequestName() string {
	return "GetOutputList"
}

// Represents the response body for the GetOutputList request.
type GetOutputListResponse struct {
	// Array of outputs
	Outputs []interface{} `json:"outputs,omitempty"`
}

// Gets the list of available outputs.
func (c *Client) GetOutputList(paramss ...*GetOutputListParams) (*GetOutputListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetOutputListParams{{}}
	}
	params := paramss[0]
	data := &GetOutputListResponse{}
	return data, c.SendRequest(params, data)
}
