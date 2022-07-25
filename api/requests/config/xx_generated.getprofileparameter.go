// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetProfileParameterParams represents the params body for the "GetProfileParameter" request.
Gets a parameter from the current profile's configuration.
*/
type GetProfileParameterParams struct {
	requests.ParamsBasic

	// Category of the parameter to get
	ParameterCategory string `json:"parameterCategory,omitempty"`

	// Name of the parameter to get
	ParameterName string `json:"parameterName,omitempty"`
}

// GetSelfName just returns "GetProfileParameter".
func (o *GetProfileParameterParams) GetSelfName() string {
	return "GetProfileParameter"
}

/*
GetProfileParameterResponse represents the response body for the "GetProfileParameter" request.
Gets a parameter from the current profile's configuration.
*/
type GetProfileParameterResponse struct {
	requests.ResponseBasic

	// Default value associated with the parameter. `null` if no default
	DefaultParameterValue string `json:"defaultParameterValue,omitempty"`

	// Value associated with the parameter. `null` if not set and no default
	ParameterValue string `json:"parameterValue,omitempty"`
}

// GetProfileParameter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetProfileParameter(params *GetProfileParameterParams) (*GetProfileParameterResponse, error) {
	data := &GetProfileParameterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
