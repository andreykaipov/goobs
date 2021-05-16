// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartRecordingParams represents the params body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingParams struct {
	requests.ParamsBasic
}

// Name just returns "StartRecording".
func (o *StartRecordingParams) Name() string {
	return "StartRecording"
}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingResponse struct {
	requests.ResponseBasic
}

// StartRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
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
