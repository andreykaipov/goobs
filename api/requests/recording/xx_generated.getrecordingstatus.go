// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetRecordingStatusParams represents the params body for the "GetRecordingStatus" request.
Get current recording status.
Since 4.9.0.
*/
type GetRecordingStatusParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetRecordingStatus".
func (o *GetRecordingStatusParams) GetSelfName() string {
	return "GetRecordingStatus"
}

/*
GetRecordingStatusResponse represents the response body for the "GetRecordingStatus" request.
Get current recording status.
Since v4.9.0.
*/
type GetRecordingStatusResponse struct {
	requests.ResponseBasic

	// Current recording status.
	IsRecording bool `json:"isRecording"`

	// Whether the recording is paused or not.
	IsRecordingPaused bool `json:"isRecordingPaused"`

	// Time elapsed since recording started (only present if currently recording).
	RecordTimecode string `json:"recordTimecode,omitempty"`

	// Absolute path to the recording file (only present if currently recording).
	RecordingFilename string `json:"recordingFilename,omitempty"`
}

// GetRecordingStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetRecordingStatus(paramss ...*GetRecordingStatusParams) (*GetRecordingStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordingStatusParams{{}}
	}
	params := paramss[0]
	data := &GetRecordingStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
