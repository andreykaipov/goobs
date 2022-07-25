// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputAudioTracksParams represents the params body for the "SetInputAudioTracks" request.
Sets the enable state of audio tracks of an input.
*/
type SetInputAudioTracksParams struct {
	requests.ParamsBasic

	// Track settings to apply
	InputAudioTracks interface{} `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "SetInputAudioTracks".
func (o *SetInputAudioTracksParams) GetSelfName() string {
	return "SetInputAudioTracks"
}

/*
SetInputAudioTracksResponse represents the response body for the "SetInputAudioTracks" request.
Sets the enable state of audio tracks of an input.
*/
type SetInputAudioTracksResponse struct {
	requests.ResponseBasic
}

// SetInputAudioTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioTracks(params *SetInputAudioTracksParams) (*SetInputAudioTracksResponse, error) {
	data := &SetInputAudioTracksResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
