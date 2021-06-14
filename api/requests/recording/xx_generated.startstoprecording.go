// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.
Toggle recording on or off (depending on the current recording state).
Since 0.3.
*/
type StartStopRecordingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartStopRecording".
func (o *StartStopRecordingParams) GetSelfName() string {
	return "StartStopRecording"
}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.
Toggle recording on or off (depending on the current recording state).
Since v0.3.
*/
type StartStopRecordingResponse struct {
	requests.ResponseBasic
}

// StartStopRecording sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) StartStopRecording(paramss ...*StartStopRecordingParams) (*StartStopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopRecordingParams{{}}
	}
	params := paramss[0]
	data := &StartStopRecordingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
