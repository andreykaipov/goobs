// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetSpecialInputs request.
type GetSpecialInputsParams struct{}

// Returns the associated request.
func (o *GetSpecialInputsParams) GetRequestName() string {
	return "GetSpecialInputs"
}

// Represents the response body for the GetSpecialInputs request.
type GetSpecialInputsResponse struct {
	// Name of the Desktop Audio input
	Desktop1 string `json:"desktop1,omitempty"`

	// Name of the Desktop Audio 2 input
	Desktop2 string `json:"desktop2,omitempty"`

	// Name of the Mic/Auxiliary Audio input
	Mic1 string `json:"mic1,omitempty"`

	// Name of the Mic/Auxiliary Audio 2 input
	Mic2 string `json:"mic2,omitempty"`

	// Name of the Mic/Auxiliary Audio 3 input
	Mic3 string `json:"mic3,omitempty"`

	// Name of the Mic/Auxiliary Audio 4 input
	Mic4 string `json:"mic4,omitempty"`
}

// Gets the names of all special inputs.
func (c *Client) GetSpecialInputs(paramss ...*GetSpecialInputsParams) (*GetSpecialInputsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSpecialInputsParams{{}}
	}
	params := paramss[0]
	data := &GetSpecialInputsResponse{}
	return data, c.SendRequest(params, data)
}
