// This file has been automatically generated. Don't edit it.

package ui

import requests "github.com/andreykaipov/goobs/api/requests"

/*
OpenInputInteractDialogParams represents the params body for the "OpenInputInteractDialog" request.
Opens the interact dialog of an input.
*/
type OpenInputInteractDialogParams struct {
	requests.ParamsBasic

	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "OpenInputInteractDialog".
func (o *OpenInputInteractDialogParams) GetSelfName() string {
	return "OpenInputInteractDialog"
}

/*
OpenInputInteractDialogResponse represents the response body for the "OpenInputInteractDialog" request.
Opens the interact dialog of an input.
*/
type OpenInputInteractDialogResponse struct {
	requests.ResponseBasic
}

// OpenInputInteractDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputInteractDialog(
	params *OpenInputInteractDialogParams,
) (*OpenInputInteractDialogResponse, error) {
	data := &OpenInputInteractDialogResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
