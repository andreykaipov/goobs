// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.
Toggle recording on or off.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingParams struct {
	requests.ParamsBasic
}

// Name just returns "StartStopRecording".
func (o *StartStopRecordingParams) Name() string {
	return "StartStopRecording"
}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.
Toggle recording on or off.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingResponse struct {
	requests.ResponseBasic
}

// StartStopRecording sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopRecording(
	paramss ...*StartStopRecordingParams,
) (*StartStopRecordingResponse, error) {
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
