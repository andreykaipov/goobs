// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputNameParams represents the params body for the "SetInputName" request.
Sets the name of an input (rename).
*/
type SetInputNameParams struct {
	requests.ParamsBasic

	// Current input name
	InputName string `json:"inputName,omitempty"`

	// New name for the input
	NewInputName string `json:"newInputName,omitempty"`
}

// GetSelfName just returns "SetInputName".
func (o *SetInputNameParams) GetSelfName() string {
	return "SetInputName"
}

/*
SetInputNameResponse represents the response body for the "SetInputName" request.
Sets the name of an input (rename).
*/
type SetInputNameResponse struct {
	requests.ResponseBasic
}

// SetInputName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputName(params *SetInputNameParams) (*SetInputNameResponse, error) {
	data := &SetInputNameResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
