// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the GetRecordStatus request.
type GetRecordStatusParams struct{}

// Returns the associated request.
func (o *GetRecordStatusParams) GetRequestName() string {
	return "GetRecordStatus"
}

// Represents the response body for the GetRecordStatus request.
type GetRecordStatusResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// Number of bytes sent by the output
	OutputBytes float64 `json:"outputBytes,omitempty"`

	// Current duration in milliseconds for the output
	OutputDuration float64 `json:"outputDuration,omitempty"`

	// Whether the output is paused
	OutputPaused bool `json:"outputPaused,omitempty"`

	// Current formatted timecode string for the output
	OutputTimecode string `json:"outputTimecode,omitempty"`
}

// Gets the status of the record output.
func (c *Client) GetRecordStatus(paramss ...*GetRecordStatusParams) (*GetRecordStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordStatusParams{{}}
	}
	params := paramss[0]
	data := &GetRecordStatusResponse{}
	return data, c.client.SendRequest(params, data)
}
