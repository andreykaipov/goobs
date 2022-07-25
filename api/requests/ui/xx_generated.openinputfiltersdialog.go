// This file has been automatically generated. Don't edit it.

package ui

import requests "github.com/andreykaipov/goobs/api/requests"

/*
OpenInputFiltersDialogParams represents the params body for the "OpenInputFiltersDialog" request.
Opens the filters dialog of an input.
*/
type OpenInputFiltersDialogParams struct {
	requests.ParamsBasic

	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "OpenInputFiltersDialog".
func (o *OpenInputFiltersDialogParams) GetSelfName() string {
	return "OpenInputFiltersDialog"
}

/*
OpenInputFiltersDialogResponse represents the response body for the "OpenInputFiltersDialog" request.
Opens the filters dialog of an input.
*/
type OpenInputFiltersDialogResponse struct {
	requests.ResponseBasic
}

// OpenInputFiltersDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputFiltersDialog(params *OpenInputFiltersDialogParams) (*OpenInputFiltersDialogResponse, error) {
	data := &OpenInputFiltersDialogResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
