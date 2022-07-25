// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputAudioSyncOffsetParams represents the params body for the "SetInputAudioSyncOffset" request.
Sets the audio sync offset of an input.
*/
type SetInputAudioSyncOffsetParams struct {
	requests.ParamsBasic

	// New audio sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input to set the audio sync offset of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "SetInputAudioSyncOffset".
func (o *SetInputAudioSyncOffsetParams) GetSelfName() string {
	return "SetInputAudioSyncOffset"
}

/*
SetInputAudioSyncOffsetResponse represents the response body for the "SetInputAudioSyncOffset" request.
Sets the audio sync offset of an input.
*/
type SetInputAudioSyncOffsetResponse struct {
	requests.ResponseBasic
}

// SetInputAudioSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioSyncOffset(
	params *SetInputAudioSyncOffsetParams,
) (*SetInputAudioSyncOffsetResponse, error) {
	data := &SetInputAudioSyncOffsetResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
