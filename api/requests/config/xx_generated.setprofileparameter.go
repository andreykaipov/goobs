// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetProfileParameterParams represents the params body for the "SetProfileParameter" request.
Sets the value of a parameter in the current profile's configuration.
*/
type SetProfileParameterParams struct {
	requests.ParamsBasic

	// Category of the parameter to set
	ParameterCategory string `json:"parameterCategory,omitempty"`

	// Name of the parameter to set
	ParameterName string `json:"parameterName,omitempty"`

	// Value of the parameter to set. Use `null` to delete
	ParameterValue string `json:"parameterValue,omitempty"`
}

// GetSelfName just returns "SetProfileParameter".
func (o *SetProfileParameterParams) GetSelfName() string {
	return "SetProfileParameter"
}

/*
SetProfileParameterResponse represents the response body for the "SetProfileParameter" request.
Sets the value of a parameter in the current profile's configuration.
*/
type SetProfileParameterResponse struct {
	requests.ResponseBasic
}

// SetProfileParameter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetProfileParameter(params *SetProfileParameterParams) (*SetProfileParameterResponse, error) {
	data := &SetProfileParameterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
