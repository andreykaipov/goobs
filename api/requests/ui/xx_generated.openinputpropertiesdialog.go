// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenInputPropertiesDialog request.
type OpenInputPropertiesDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *OpenInputPropertiesDialogParams) GetRequestName() string {
	return "OpenInputPropertiesDialog"
}

// Represents the response body for the OpenInputPropertiesDialog request.
type OpenInputPropertiesDialogResponse struct{}

// Opens the properties dialog of an input.
func (c *Client) OpenInputPropertiesDialog(
	params *OpenInputPropertiesDialogParams,
) (*OpenInputPropertiesDialogResponse, error) {
	data := &OpenInputPropertiesDialogResponse{}
	return data, c.SendRequest(params, data)
}
