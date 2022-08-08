// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the GetProfileParameter request.
type GetProfileParameterParams struct {
	// Category of the parameter to get
	ParameterCategory string `json:"parameterCategory,omitempty"`

	// Name of the parameter to get
	ParameterName string `json:"parameterName,omitempty"`
}

// Returns the associated request.
func (o *GetProfileParameterParams) GetRequestName() string {
	return "GetProfileParameter"
}

// Represents the response body for the GetProfileParameter request.
type GetProfileParameterResponse struct {
	// Default value associated with the parameter. `null` if no default
	DefaultParameterValue string `json:"defaultParameterValue,omitempty"`

	// Value associated with the parameter. `null` if not set and no default
	ParameterValue string `json:"parameterValue,omitempty"`
}

// Gets a parameter from the current profile's configuration.
func (c *Client) GetProfileParameter(params *GetProfileParameterParams) (*GetProfileParameterResponse, error) {
	data := &GetProfileParameterResponse{}
	return data, c.SendRequest(params, data)
}
