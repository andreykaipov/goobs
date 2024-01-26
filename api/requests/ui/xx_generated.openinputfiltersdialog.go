// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenInputFiltersDialog request.
type OpenInputFiltersDialogParams struct {
	// Name of the input to open the dialog of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to open the dialog of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewOpenInputFiltersDialogParams() *OpenInputFiltersDialogParams {
	return &OpenInputFiltersDialogParams{}
}
func (o *OpenInputFiltersDialogParams) WithInputName(x string) *OpenInputFiltersDialogParams {
	o.InputName = &x
	return o
}
func (o *OpenInputFiltersDialogParams) WithInputUuid(x string) *OpenInputFiltersDialogParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *OpenInputFiltersDialogParams) GetRequestName() string {
	return "OpenInputFiltersDialog"
}

// Represents the response body for the OpenInputFiltersDialog request.
type OpenInputFiltersDialogResponse struct {
	_response
}

// Opens the filters dialog of an input.
func (c *Client) OpenInputFiltersDialog(
	paramss ...*OpenInputFiltersDialogParams,
) (*OpenInputFiltersDialogResponse, error) {
	if len(paramss) == 0 {
		paramss = []*OpenInputFiltersDialogParams{{}}
	}
	params := paramss[0]
	data := &OpenInputFiltersDialogResponse{}
	return data, c.client.SendRequest(params, data)
}
