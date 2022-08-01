// This file has been automatically generated. Don't edit it.

package ui

/*
OpenInputFiltersDialogParams represents the params body for the "OpenInputFiltersDialog" request.
Opens the filters dialog of an input.
*/
type OpenInputFiltersDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

/*
OpenInputFiltersDialogResponse represents the response body for the "OpenInputFiltersDialog" request.
Opens the filters dialog of an input.
*/
type OpenInputFiltersDialogResponse struct{}

// OpenInputFiltersDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputFiltersDialog(params *OpenInputFiltersDialogParams) (*OpenInputFiltersDialogResponse, error) {
	resp, err := c.SendRequest("OpenInputFiltersDialog", params)
	if err != nil {
		return nil, err
	}
	return resp.(*OpenInputFiltersDialogResponse), nil
}
