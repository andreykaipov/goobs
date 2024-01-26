// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the RemoveInput request.
type RemoveInputParams struct {
	// Name of the input to remove
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to remove
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewRemoveInputParams() *RemoveInputParams {
	return &RemoveInputParams{}
}
func (o *RemoveInputParams) WithInputName(x string) *RemoveInputParams {
	o.InputName = &x
	return o
}
func (o *RemoveInputParams) WithInputUuid(x string) *RemoveInputParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *RemoveInputParams) GetRequestName() string {
	return "RemoveInput"
}

// Represents the response body for the RemoveInput request.
type RemoveInputResponse struct {
	_response
}

/*
Removes an existing input.

Note: Will immediately remove all associated scene items.
*/
func (c *Client) RemoveInput(paramss ...*RemoveInputParams) (*RemoveInputResponse, error) {
	if len(paramss) == 0 {
		paramss = []*RemoveInputParams{{}}
	}
	params := paramss[0]
	data := &RemoveInputResponse{}
	return data, c.client.SendRequest(params, data)
}
