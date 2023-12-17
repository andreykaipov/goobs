// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetInputAudioTracks request.
type GetInputAudioTracksParams struct {
	// Name of the input
	InputName *string `json:"inputName,omitempty"`
}

func NewGetInputAudioTracksParams() *GetInputAudioTracksParams {
	return &GetInputAudioTracksParams{}
}
func (o *GetInputAudioTracksParams) WithInputName(x string) *GetInputAudioTracksParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *GetInputAudioTracksParams) GetRequestName() string {
	return "GetInputAudioTracks"
}

// Represents the response body for the GetInputAudioTracks request.
type GetInputAudioTracksResponse struct {
	// Object of audio tracks and associated enable states
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`
}

// Gets the enable state of all audio tracks of an input.
func (c *Client) GetInputAudioTracks(params *GetInputAudioTracksParams) (*GetInputAudioTracksResponse, error) {
	data := &GetInputAudioTracksResponse{}
	return data, c.SendRequest(params, data)
}
