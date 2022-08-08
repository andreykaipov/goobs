// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputName request.
type SetInputNameParams struct {
	// Current input name
	InputName string `json:"inputName,omitempty"`

	// New name for the input
	NewInputName string `json:"newInputName,omitempty"`
}

// Returns the associated request.
func (o *SetInputNameParams) GetRequestName() string {
	return "SetInputName"
}

// Represents the response body for the SetInputName request.
type SetInputNameResponse struct{}

// Sets the name of an input (rename).
func (c *Client) SetInputName(params *SetInputNameParams) (*SetInputNameResponse, error) {
	data := &SetInputNameResponse{}
	return data, c.SendRequest(params, data)
}
