// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputListParams represents the params body for the "GetInputList" request.
Gets an array of all inputs in OBS.
*/
type GetInputListParams struct {
	requests.ParamsBasic

	// Restrict the array to only inputs of the specified kind
	InputKind string `json:"inputKind,omitempty"`
}

// GetSelfName just returns "GetInputList".
func (o *GetInputListParams) GetSelfName() string {
	return "GetInputList"
}

/*
GetInputListResponse represents the response body for the "GetInputList" request.
Gets an array of all inputs in OBS.
*/
type GetInputListResponse struct {
	requests.ResponseBasic

	// Array of inputs
	Inputs []interface{} `json:"inputs,omitempty"`
}

// GetInputList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputList(params *GetInputListParams) (*GetInputListResponse, error) {
	data := &GetInputListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
