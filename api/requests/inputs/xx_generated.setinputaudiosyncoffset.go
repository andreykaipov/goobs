// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputAudioSyncOffset request.
type SetInputAudioSyncOffsetParams struct {
	// New audio sync offset in milliseconds
	InputAudioSyncOffset *float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input to set the audio sync offset of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to set the audio sync offset of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputAudioSyncOffsetParams() *SetInputAudioSyncOffsetParams {
	return &SetInputAudioSyncOffsetParams{}
}
func (o *SetInputAudioSyncOffsetParams) WithInputAudioSyncOffset(x float64) *SetInputAudioSyncOffsetParams {
	o.InputAudioSyncOffset = &x
	return o
}
func (o *SetInputAudioSyncOffsetParams) WithInputName(x string) *SetInputAudioSyncOffsetParams {
	o.InputName = &x
	return o
}
func (o *SetInputAudioSyncOffsetParams) WithInputUuid(x string) *SetInputAudioSyncOffsetParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputAudioSyncOffsetParams) GetRequestName() string {
	return "SetInputAudioSyncOffset"
}

// Represents the response body for the SetInputAudioSyncOffset request.
type SetInputAudioSyncOffsetResponse struct {
	_response
}

// Sets the audio sync offset of an input.
func (c *Client) SetInputAudioSyncOffset(
	params *SetInputAudioSyncOffsetParams,
) (*SetInputAudioSyncOffsetResponse, error) {
	data := &SetInputAudioSyncOffsetResponse{}
	return data, c.client.SendRequest(params, data)
}
