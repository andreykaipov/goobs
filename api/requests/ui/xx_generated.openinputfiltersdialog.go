// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenInputFiltersDialog request.
type OpenInputFiltersDialogParams struct {
	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *OpenInputFiltersDialogParams) GetRequestName() string {
	return "OpenInputFiltersDialog"
}

// Represents the response body for the OpenInputFiltersDialog request.
type OpenInputFiltersDialogResponse struct{}

// Opens the filters dialog of an input.
func (c *Client) OpenInputFiltersDialog(params *OpenInputFiltersDialogParams) (*OpenInputFiltersDialogResponse, error) {
	data := &OpenInputFiltersDialogResponse{}
	return data, c.SendRequest(params, data)
}
