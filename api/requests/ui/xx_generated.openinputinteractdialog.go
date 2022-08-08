// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenInputInteractDialog request.
type OpenInputInteractDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *OpenInputInteractDialogParams) GetRequestName() string {
	return "OpenInputInteractDialog"
}

// Represents the response body for the OpenInputInteractDialog request.
type OpenInputInteractDialogResponse struct{}

// Opens the interact dialog of an input.
func (c *Client) OpenInputInteractDialog(
	params *OpenInputInteractDialogParams,
) (*OpenInputInteractDialogResponse, error) {
	data := &OpenInputInteractDialogResponse{}
	return data, c.SendRequest(params, data)
}
