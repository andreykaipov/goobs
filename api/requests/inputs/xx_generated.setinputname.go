// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputNameParams represents the params body for the "SetInputName" request.
Sets the name of an input (rename).
*/
type SetInputNameParams struct {
	// Current input name
	InputName string `json:"inputName,omitempty"`

	// New name for the input
	NewInputName string `json:"newInputName,omitempty"`
}

/*
SetInputNameResponse represents the response body for the "SetInputName" request.
Sets the name of an input (rename).
*/
type SetInputNameResponse struct{}

// SetInputName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputName(params *SetInputNameParams) (*SetInputNameResponse, error) {
	resp, err := c.SendRequest("SetInputName", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputNameResponse), nil
}
