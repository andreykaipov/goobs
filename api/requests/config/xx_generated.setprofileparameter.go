// This file has been automatically generated. Don't edit it.

package config

/*
SetProfileParameterParams represents the params body for the "SetProfileParameter" request.
Sets the value of a parameter in the current profile's configuration.
*/
type SetProfileParameterParams struct {
	// Category of the parameter to set
	ParameterCategory string `json:"parameterCategory,omitempty"`

	// Name of the parameter to set
	ParameterName string `json:"parameterName,omitempty"`

	// Value of the parameter to set. Use `null` to delete
	ParameterValue string `json:"parameterValue,omitempty"`
}

/*
SetProfileParameterResponse represents the response body for the "SetProfileParameter" request.
Sets the value of a parameter in the current profile's configuration.
*/
type SetProfileParameterResponse struct{}

// SetProfileParameter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetProfileParameter(params *SetProfileParameterParams) (*SetProfileParameterResponse, error) {
	resp, err := c.SendRequest("SetProfileParameter", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetProfileParameterResponse), nil
}
