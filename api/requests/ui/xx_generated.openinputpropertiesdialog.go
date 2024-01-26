// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenInputPropertiesDialog request.
type OpenInputPropertiesDialogParams struct {
	// Name of the input to open the dialog of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to open the dialog of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewOpenInputPropertiesDialogParams() *OpenInputPropertiesDialogParams {
	return &OpenInputPropertiesDialogParams{}
}
func (o *OpenInputPropertiesDialogParams) WithInputName(x string) *OpenInputPropertiesDialogParams {
	o.InputName = &x
	return o
}
func (o *OpenInputPropertiesDialogParams) WithInputUuid(x string) *OpenInputPropertiesDialogParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *OpenInputPropertiesDialogParams) GetRequestName() string {
	return "OpenInputPropertiesDialog"
}

// Represents the response body for the OpenInputPropertiesDialog request.
type OpenInputPropertiesDialogResponse struct {
	_response
}

// Opens the properties dialog of an input.
func (c *Client) OpenInputPropertiesDialog(
	paramss ...*OpenInputPropertiesDialogParams,
) (*OpenInputPropertiesDialogResponse, error) {
	if len(paramss) == 0 {
		paramss = []*OpenInputPropertiesDialogParams{{}}
	}
	params := paramss[0]
	data := &OpenInputPropertiesDialogResponse{}
	return data, c.client.SendRequest(params, data)
}
