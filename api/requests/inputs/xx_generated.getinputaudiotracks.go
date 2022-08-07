// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
GetInputAudioTracksParams represents the params body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksParams struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputAudioTracksResponse represents the response body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksResponse struct {
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`
}

// GetInputAudioTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioTracks(params *GetInputAudioTracksParams) (*GetInputAudioTracksResponse, error) {
	resp, err := c.SendRequest("GetInputAudioTracks", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputAudioTracksResponse), nil
}
