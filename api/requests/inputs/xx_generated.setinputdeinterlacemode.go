// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputDeinterlaceMode request.
type SetInputDeinterlaceModeParams struct {
	// Deinterlace mode for the input
	InputDeinterlaceMode *string `json:"inputDeinterlaceMode,omitempty"`

	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputDeinterlaceModeParams() *SetInputDeinterlaceModeParams {
	return &SetInputDeinterlaceModeParams{}
}
func (o *SetInputDeinterlaceModeParams) WithInputDeinterlaceMode(x string) *SetInputDeinterlaceModeParams {
	o.InputDeinterlaceMode = &x
	return o
}
func (o *SetInputDeinterlaceModeParams) WithInputName(x string) *SetInputDeinterlaceModeParams {
	o.InputName = &x
	return o
}
func (o *SetInputDeinterlaceModeParams) WithInputUuid(x string) *SetInputDeinterlaceModeParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputDeinterlaceModeParams) GetRequestName() string {
	return "SetInputDeinterlaceMode"
}

// Represents the response body for the SetInputDeinterlaceMode request.
type SetInputDeinterlaceModeResponse struct {
	_response
}

/*
Sets the deinterlace mode of an input.

Note: Deinterlacing functionality is restricted to async inputs only.
*/
func (c *Client) SetInputDeinterlaceMode(
	params *SetInputDeinterlaceModeParams,
) (*SetInputDeinterlaceModeResponse, error) {
	data := &SetInputDeinterlaceModeResponse{}
	return data, c.client.SendRequest(params, data)
}
