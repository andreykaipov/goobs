// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputMute request.
type SetInputMuteParams struct {
	// Whether to mute the input or not
	InputMuted *bool `json:"inputMuted,omitempty"`

	// Name of the input to set the mute state of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to set the mute state of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputMuteParams() *SetInputMuteParams {
	return &SetInputMuteParams{}
}
func (o *SetInputMuteParams) WithInputMuted(x bool) *SetInputMuteParams {
	o.InputMuted = &x
	return o
}
func (o *SetInputMuteParams) WithInputName(x string) *SetInputMuteParams {
	o.InputName = &x
	return o
}
func (o *SetInputMuteParams) WithInputUuid(x string) *SetInputMuteParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputMuteParams) GetRequestName() string {
	return "SetInputMute"
}

// Represents the response body for the SetInputMute request.
type SetInputMuteResponse struct {
	_response
}

// Sets the audio mute state of an input.
func (c *Client) SetInputMute(params *SetInputMuteParams) (*SetInputMuteResponse, error) {
	data := &SetInputMuteResponse{}
	return data, c.client.SendRequest(params, data)
}
