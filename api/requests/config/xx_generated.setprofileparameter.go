// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetProfileParameter request.
type SetProfileParameterParams struct {
	// Category of the parameter to set
	ParameterCategory *string `json:"parameterCategory,omitempty"`

	// Name of the parameter to set
	ParameterName *string `json:"parameterName,omitempty"`

	// Value of the parameter to set. Use `null` to delete
	ParameterValue *string `json:"parameterValue,omitempty"`
}

func NewSetProfileParameterParams() *SetProfileParameterParams {
	return &SetProfileParameterParams{}
}
func (o *SetProfileParameterParams) WithParameterCategory(x string) *SetProfileParameterParams {
	o.ParameterCategory = &x
	return o
}
func (o *SetProfileParameterParams) WithParameterName(x string) *SetProfileParameterParams {
	o.ParameterName = &x
	return o
}
func (o *SetProfileParameterParams) WithParameterValue(x string) *SetProfileParameterParams {
	o.ParameterValue = &x
	return o
}

// Returns the associated request.
func (o *SetProfileParameterParams) GetRequestName() string {
	return "SetProfileParameter"
}

// Represents the response body for the SetProfileParameter request.
type SetProfileParameterResponse struct {
	_response
}

// Sets the value of a parameter in the current profile's configuration.
func (c *Client) SetProfileParameter(params *SetProfileParameterParams) (*SetProfileParameterResponse, error) {
	data := &SetProfileParameterResponse{}
	return data, c.client.SendRequest(params, data)
}
