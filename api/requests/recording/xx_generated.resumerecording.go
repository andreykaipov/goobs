// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResumeRecordingParams represents the params body for the "ResumeRecording" request.
Resume/unpause the current recording (if paused).
Returns an error if recording is not active or not paused.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#ResumeRecording.
*/
type ResumeRecordingParams struct {
	requests.ParamsBasic
}

// Name just returns "ResumeRecording".
func (o *ResumeRecordingParams) Name() string {
	return "ResumeRecording"
}

/*
ResumeRecordingResponse represents the response body for the "ResumeRecording" request.
Resume/unpause the current recording (if paused).
Returns an error if recording is not active or not paused.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#ResumeRecording.
*/
type ResumeRecordingResponse struct {
	requests.ResponseBasic
}

// ResumeRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ResumeRecording(
	paramss ...*ResumeRecordingParams,
) (*ResumeRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ResumeRecordingParams{{}}
	}
	params := paramss[0]
	data := &ResumeRecordingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
