// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the SetInputAudioTracks request.
type SetInputAudioTracksParams struct {
	// Track settings to apply
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewSetInputAudioTracksParams() *SetInputAudioTracksParams {
	return &SetInputAudioTracksParams{}
}
func (o *SetInputAudioTracksParams) WithInputAudioTracks(x *typedefs.InputAudioTracks) *SetInputAudioTracksParams {
	o.InputAudioTracks = x
	return o
}
func (o *SetInputAudioTracksParams) WithInputName(x string) *SetInputAudioTracksParams {
	o.InputName = &x
	return o
}
func (o *SetInputAudioTracksParams) WithInputUuid(x string) *SetInputAudioTracksParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *SetInputAudioTracksParams) GetRequestName() string {
	return "SetInputAudioTracks"
}

// Represents the response body for the SetInputAudioTracks request.
type SetInputAudioTracksResponse struct {
	_response
}

// Sets the enable state of audio tracks of an input.
func (c *Client) SetInputAudioTracks(params *SetInputAudioTracksParams) (*SetInputAudioTracksResponse, error) {
	data := &SetInputAudioTracksResponse{}
	return data, c.client.SendRequest(params, data)
}
