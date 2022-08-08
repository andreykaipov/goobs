// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the SetInputAudioTracks request.
type SetInputAudioTracksParams struct {
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *SetInputAudioTracksParams) GetRequestName() string {
	return "SetInputAudioTracks"
}

// Represents the response body for the SetInputAudioTracks request.
type SetInputAudioTracksResponse struct{}

// Sets the enable state of audio tracks of an input.
func (c *Client) SetInputAudioTracks(params *SetInputAudioTracksParams) (*SetInputAudioTracksResponse, error) {
	data := &SetInputAudioTracksResponse{}
	return data, c.SendRequest(params, data)
}
