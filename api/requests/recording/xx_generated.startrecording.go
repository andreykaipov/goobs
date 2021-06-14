// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartRecordingParams represents the params body for the "StartRecording" request.
Start recording.
Will return an `error` if recording is already active.
Since 4.1.0.
*/
type StartRecordingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartRecording".
func (o *StartRecordingParams) GetSelfName() string {
	return "StartRecording"
}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.
Start recording.
Will return an `error` if recording is already active.
Since v4.1.0.
*/
type StartRecordingResponse struct {
	requests.ResponseBasic
}

// StartRecording sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StartRecording(paramss ...*StartRecordingParams) (*StartRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordingParams{{}}
	}
	params := paramss[0]
	data := &StartRecordingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
