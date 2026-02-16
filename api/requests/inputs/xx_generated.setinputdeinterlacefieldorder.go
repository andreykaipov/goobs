// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputDeinterlaceFieldOrder request.
type SetInputDeinterlaceFieldOrderParams struct {
	// Deinterlace field order for the input
	InputDeinterlaceFieldOrder *string `json:"inputDeinterlaceFieldOrder,omitempty"`

	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputDeinterlaceFieldOrderParams() *SetInputDeinterlaceFieldOrderParams {
	return &SetInputDeinterlaceFieldOrderParams{}
}

func (o *SetInputDeinterlaceFieldOrderParams) WithInputDeinterlaceFieldOrder(
	x string,
) *SetInputDeinterlaceFieldOrderParams {
	o.InputDeinterlaceFieldOrder = &x
	return o
}
func (o *SetInputDeinterlaceFieldOrderParams) WithInputName(x string) *SetInputDeinterlaceFieldOrderParams {
	o.InputName = &x
	return o
}
func (o *SetInputDeinterlaceFieldOrderParams) WithInputUuid(x string) *SetInputDeinterlaceFieldOrderParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputDeinterlaceFieldOrderParams) GetRequestName() string {
	return "SetInputDeinterlaceFieldOrder"
}

// Represents the response body for the SetInputDeinterlaceFieldOrder request.
type SetInputDeinterlaceFieldOrderResponse struct {
	_response
}

/*
Sets the deinterlace field order of an input.

Note: Deinterlacing functionality is restricted to async inputs only.
*/
func (c *Client) SetInputDeinterlaceFieldOrder(
	params *SetInputDeinterlaceFieldOrderParams,
) (*SetInputDeinterlaceFieldOrderResponse, error) {
	data := &SetInputDeinterlaceFieldOrderResponse{}
	return data, c.client.SendRequest(params, data)
}
