// This file has been automatically generated. Don't edit it.

package ui

import requests "github.com/andreykaipov/goobs/api/requests"

/*
OpenInputPropertiesDialogParams represents the params body for the "OpenInputPropertiesDialog" request.
Opens the properties dialog of an input.
*/
type OpenInputPropertiesDialogParams struct {
	requests.ParamsBasic

	// Name of the input to open the dialog of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "OpenInputPropertiesDialog".
func (o *OpenInputPropertiesDialogParams) GetSelfName() string {
	return "OpenInputPropertiesDialog"
}

/*
OpenInputPropertiesDialogResponse represents the response body for the "OpenInputPropertiesDialog" request.
Opens the properties dialog of an input.
*/
type OpenInputPropertiesDialogResponse struct {
	requests.ResponseBasic
}

// OpenInputPropertiesDialog sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenInputPropertiesDialog(
	params *OpenInputPropertiesDialogParams,
) (*OpenInputPropertiesDialogResponse, error) {
	data := &OpenInputPropertiesDialogResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
