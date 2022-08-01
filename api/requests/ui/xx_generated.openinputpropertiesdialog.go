// This file has been automatically generated. Don't edit it.

package ui

/*
OpenInputPropertiesDialogParams represents the params body for the "OpenInputPropertiesDialog" request.
Opens the properties dialog of an input.
*/
type OpenInputPropertiesDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

/*
OpenInputPropertiesDialogResponse represents the response body for the "OpenInputPropertiesDialog" request.
Opens the properties dialog of an input.
*/
type OpenInputPropertiesDialogResponse struct{}

// OpenInputPropertiesDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputPropertiesDialog(
	params *OpenInputPropertiesDialogParams,
) (*OpenInputPropertiesDialogResponse, error) {
	resp, err := c.SendRequest("OpenInputPropertiesDialog", params)
	if err != nil {
		return nil, err
	}
	return resp.(*OpenInputPropertiesDialogResponse), nil
}
