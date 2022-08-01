// This file has been automatically generated. Don't edit it.

package ui

/*
OpenInputInteractDialogParams represents the params body for the "OpenInputInteractDialog" request.
Opens the interact dialog of an input.
*/
type OpenInputInteractDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

/*
OpenInputInteractDialogResponse represents the response body for the "OpenInputInteractDialog" request.
Opens the interact dialog of an input.
*/
type OpenInputInteractDialogResponse struct{}

// OpenInputInteractDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputInteractDialog(
	params *OpenInputInteractDialogParams,
) (*OpenInputInteractDialogResponse, error) {
	resp, err := c.SendRequest("OpenInputInteractDialog", params)
	if err != nil {
		return nil, err
	}
	return resp.(*OpenInputInteractDialogResponse), nil
}
