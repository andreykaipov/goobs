// This file has been automatically generated. Don't edit it.

package config

/*
GetProfileParameterParams represents the params body for the "GetProfileParameter" request.
Gets a parameter from the current profile's configuration.
*/
type GetProfileParameterParams struct {
	// Category of the parameter to get
	ParameterCategory string `json:"parameterCategory,omitempty"`

	// Name of the parameter to get
	ParameterName string `json:"parameterName,omitempty"`
}

/*
GetProfileParameterResponse represents the response body for the "GetProfileParameter" request.
Gets a parameter from the current profile's configuration.
*/
type GetProfileParameterResponse struct {
	// Default value associated with the parameter. `null` if no default
	DefaultParameterValue string `json:"defaultParameterValue,omitempty"`

	// Value associated with the parameter. `null` if not set and no default
	ParameterValue string `json:"parameterValue,omitempty"`
}

// GetProfileParameter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetProfileParameter(params *GetProfileParameterParams) (*GetProfileParameterResponse, error) {
	resp, err := c.SendRequest("GetProfileParameter", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetProfileParameterResponse), nil
}
