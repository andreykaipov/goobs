// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the PressInputPropertiesButton request.
type PressInputPropertiesButtonParams struct {
	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`

	// Name of the button property to press
	PropertyName *string `json:"propertyName,omitempty"`
}

func NewPressInputPropertiesButtonParams() *PressInputPropertiesButtonParams {
	return &PressInputPropertiesButtonParams{}
}
func (o *PressInputPropertiesButtonParams) WithInputName(x string) *PressInputPropertiesButtonParams {
	o.InputName = &x
	return o
}
func (o *PressInputPropertiesButtonParams) WithInputUuid(x string) *PressInputPropertiesButtonParams {
	o.InputUuid = &x
	return o
}
func (o *PressInputPropertiesButtonParams) WithPropertyName(x string) *PressInputPropertiesButtonParams {
	o.PropertyName = &x
	return o
}

// Returns the associated request.
func (o *PressInputPropertiesButtonParams) GetRequestName() string {
	return "PressInputPropertiesButton"
}

// Represents the response body for the PressInputPropertiesButton request.
type PressInputPropertiesButtonResponse struct {
	_response
}

/*
Presses a button in the properties of an input.

Some known `propertyName` values are:

- `refreshnocache` - Browser source reload button

Note: Use this in cases where there is a button in the properties of an input that cannot be accessed in any other way. For example, browser sources, where there is a refresh button.
*/
func (c *Client) PressInputPropertiesButton(
	params *PressInputPropertiesButtonParams,
) (*PressInputPropertiesButtonResponse, error) {
	data := &PressInputPropertiesButtonResponse{}
	return data, c.client.SendRequest(params, data)
}
