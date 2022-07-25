// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSpecialInputsParams represents the params body for the "GetSpecialInputs" request.
Gets the names of all special inputs.
*/
type GetSpecialInputsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSpecialInputs".
func (o *GetSpecialInputsParams) GetSelfName() string {
	return "GetSpecialInputs"
}

/*
GetSpecialInputsResponse represents the response body for the "GetSpecialInputs" request.
Gets the names of all special inputs.
*/
type GetSpecialInputsResponse struct {
	requests.ResponseBasic

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

// GetSpecialInputs sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetSpecialInputs(paramss ...*GetSpecialInputsParams) (*GetSpecialInputsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSpecialInputsParams{{}}
	}
	params := paramss[0]
	data := &GetSpecialInputsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
