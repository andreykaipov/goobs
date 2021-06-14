// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
PauseRecordingParams represents the params body for the "PauseRecording" request.
Pause the current recording.
Returns an error if recording is not active or already paused.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#PauseRecording.
*/
type PauseRecordingParams struct {
	requests.ParamsBasic
}

// Name just returns "PauseRecording".
func (o *PauseRecordingParams) Name() string {
	return "PauseRecording"
}

/*
PauseRecordingResponse represents the response body for the "PauseRecording" request.
Pause the current recording.
Returns an error if recording is not active or already paused.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#PauseRecording.
*/
type PauseRecordingResponse struct {
	requests.ResponseBasic
}

// PauseRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) PauseRecording(paramss ...*PauseRecordingParams) (*PauseRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*PauseRecordingParams{{}}
	}
	params := paramss[0]
	data := &PauseRecordingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
