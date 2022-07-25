// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveInputParams represents the params body for the "RemoveInput" request.
Removes an existing input.

Note: Will immediately remove all associated scene items.
*/
type RemoveInputParams struct {
	requests.ParamsBasic

	// Name of the input to remove
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "RemoveInput".
func (o *RemoveInputParams) GetSelfName() string {
	return "RemoveInput"
}

/*
RemoveInputResponse represents the response body for the "RemoveInput" request.
Removes an existing input.

Note: Will immediately remove all associated scene items.
*/
type RemoveInputResponse struct {
	requests.ResponseBasic
}

// RemoveInput sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveInput(params *RemoveInputParams) (*RemoveInputResponse, error) {
	data := &RemoveInputResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
