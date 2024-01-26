// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputName request.
type SetInputNameParams struct {
	// Current input name
	InputName *string `json:"inputName,omitempty"`

	// Current input UUID
	InputUuid *string `json:"inputUuid,omitempty"`

	// New name for the input
	NewInputName *string `json:"newInputName,omitempty"`
}

func NewSetInputNameParams() *SetInputNameParams {
	return &SetInputNameParams{}
}
func (o *SetInputNameParams) WithInputName(x string) *SetInputNameParams {
	o.InputName = &x
	return o
}
func (o *SetInputNameParams) WithInputUuid(x string) *SetInputNameParams {
	o.InputUuid = &x
	return o
}
func (o *SetInputNameParams) WithNewInputName(x string) *SetInputNameParams {
	o.NewInputName = &x
	return o
}

// Returns the associated request.
func (o *SetInputNameParams) GetRequestName() string {
	return "SetInputName"
}

// Represents the response body for the SetInputName request.
type SetInputNameResponse struct {
	_response
}

// Sets the name of an input (rename).
func (c *Client) SetInputName(params *SetInputNameParams) (*SetInputNameResponse, error) {
	data := &SetInputNameResponse{}
	return data, c.client.SendRequest(params, data)
}
