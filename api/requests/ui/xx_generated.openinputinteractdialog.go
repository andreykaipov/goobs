// This file has been automatically generated. Don't edit it.

package ui

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the OpenInputInteractDialog request.
type OpenInputInteractDialogParams struct {
	// Name of the input to open the dialog of
	InputName *string `json:"inputName,omitempty"`
}

func NewOpenInputInteractDialogParams() *OpenInputInteractDialogParams {
	return &OpenInputInteractDialogParams{}
}
func (o *OpenInputInteractDialogParams) WithInputName(x string) *OpenInputInteractDialogParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *OpenInputInteractDialogParams) GetRequestName() string {
	return "OpenInputInteractDialog"
}

// Represents the response body for the OpenInputInteractDialog request.
type OpenInputInteractDialogResponse struct {
	api.ResponseCommon
}

// Opens the interact dialog of an input.
func (c *Client) OpenInputInteractDialog(
	params *OpenInputInteractDialogParams,
) (*OpenInputInteractDialogResponse, error) {
	data := &OpenInputInteractDialogResponse{}
	return data, c.SendRequest(params, data)
}
