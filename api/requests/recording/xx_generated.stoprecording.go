// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopRecordingParams represents the params body for the "StopRecording" request.
Stop recording.
Will return an `error` if recording is not active.
Since 4.1.0.
*/
type StopRecordingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopRecording".
func (o *StopRecordingParams) GetSelfName() string {
	return "StopRecording"
}

/*
StopRecordingResponse represents the response body for the "StopRecording" request.
Stop recording.
Will return an `error` if recording is not active.
Since v4.1.0.
*/
type StopRecordingResponse struct {
	requests.ResponseBasic
}

// StopRecording sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopRecording(paramss ...*StopRecordingParams) (*StopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordingParams{{}}
	}
	params := paramss[0]
	data := &StopRecordingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
