// This file has been automatically generated. Don't edit it.

package stream

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamStatusParams represents the params body for the "GetStreamStatus" request.
Gets the status of the stream output.
*/
type GetStreamStatusParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetStreamStatus".
func (o *GetStreamStatusParams) GetSelfName() string {
	return "GetStreamStatus"
}

/*
GetStreamStatusResponse represents the response body for the "GetStreamStatus" request.
Gets the status of the stream output.
*/
type GetStreamStatusResponse struct {
	requests.ResponseBasic

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// Number of bytes sent by the output
	OutputBytes float64 `json:"outputBytes,omitempty"`

	// Current duration in milliseconds for the output
	OutputDuration float64 `json:"outputDuration,omitempty"`

	// Whether the output is currently reconnecting
	OutputReconnecting bool `json:"outputReconnecting,omitempty"`

	// Number of frames skipped by the output's process
	OutputSkippedFrames float64 `json:"outputSkippedFrames,omitempty"`

	// Current formatted timecode string for the output
	OutputTimecode string `json:"outputTimecode,omitempty"`

	// Total number of frames delivered by the output's process
	OutputTotalFrames float64 `json:"outputTotalFrames,omitempty"`
}

// GetStreamStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetStreamStatus(paramss ...*GetStreamStatusParams) (*GetStreamStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamStatusParams{{}}
	}
	params := paramss[0]
	data := &GetStreamStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
