// This file has been automatically generated. Don't edit it.

package inputs

/*
RemoveInputParams represents the params body for the "RemoveInput" request.
Removes an existing input.

Note: Will immediately remove all associated scene items.
*/
type RemoveInputParams struct {
	// Name of the input to remove
	InputName string `json:"inputName,omitempty"`
}

/*
RemoveInputResponse represents the response body for the "RemoveInput" request.
Removes an existing input.

Note: Will immediately remove all associated scene items.
*/
type RemoveInputResponse struct{}

// RemoveInput sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveInput(params *RemoveInputParams) (*RemoveInputResponse, error) {
	resp, err := c.SendRequest("RemoveInput", params)
	if err != nil {
		return nil, err
	}
	return resp.(*RemoveInputResponse), nil
}
