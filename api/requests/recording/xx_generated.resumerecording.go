// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResumeRecordingParams represents the params body for the "ResumeRecording" request.
Resume/unpause the current recording (if paused).
Returns an error if recording is not active or not paused.
Since 4.7.0.
*/
type ResumeRecordingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ResumeRecording".
func (o *ResumeRecordingParams) GetSelfName() string {
	return "ResumeRecording"
}

/*
ResumeRecordingResponse represents the response body for the "ResumeRecording" request.
Resume/unpause the current recording (if paused).
Returns an error if recording is not active or not paused.
Since v4.7.0.
*/
type ResumeRecordingResponse struct {
	requests.ResponseBasic
}

// ResumeRecording sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) ResumeRecording(paramss ...*ResumeRecordingParams) (*ResumeRecordingResponse, error) {
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
