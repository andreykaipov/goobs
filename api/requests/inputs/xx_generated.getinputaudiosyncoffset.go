// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputAudioSyncOffsetParams represents the params body for the "GetInputAudioSyncOffset" request.
Gets the audio sync offset of an input.

Note: The audio sync offset can be negative too!
*/
type GetInputAudioSyncOffsetParams struct {
	requests.ParamsBasic

	// Name of the input to get the audio sync offset of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputAudioSyncOffset".
func (o *GetInputAudioSyncOffsetParams) GetSelfName() string {
	return "GetInputAudioSyncOffset"
}

/*
GetInputAudioSyncOffsetResponse represents the response body for the "GetInputAudioSyncOffset" request.
Gets the audio sync offset of an input.

Note: The audio sync offset can be negative too!
*/
type GetInputAudioSyncOffsetResponse struct {
	requests.ResponseBasic

	// Audio sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`
}

// GetInputAudioSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioSyncOffset(
	params *GetInputAudioSyncOffsetParams,
) (*GetInputAudioSyncOffsetResponse, error) {
	data := &GetInputAudioSyncOffsetResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
