// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputAudioTracksParams represents the params body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksParams struct {
	requests.ParamsBasic

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputAudioTracks".
func (o *GetInputAudioTracksParams) GetSelfName() string {
	return "GetInputAudioTracks"
}

/*
GetInputAudioTracksResponse represents the response body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksResponse struct {
	requests.ResponseBasic

	// Object of audio tracks and associated enable states
	InputAudioTracks interface{} `json:"inputAudioTracks,omitempty"`
}

// GetInputAudioTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioTracks(params *GetInputAudioTracksParams) (*GetInputAudioTracksResponse, error) {
	data := &GetInputAudioTracksResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
