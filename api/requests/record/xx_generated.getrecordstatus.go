// This file has been automatically generated. Don't edit it.

package record

/*
GetRecordStatusParams represents the params body for the "GetRecordStatus" request.
Gets the status of the record output.
*/
type GetRecordStatusParams struct{}

/*
GetRecordStatusResponse represents the response body for the "GetRecordStatus" request.
Gets the status of the record output.
*/
type GetRecordStatusResponse struct {
	// Whether the output is paused
	OuputPaused bool `json:"ouputPaused,omitempty"`

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// Number of bytes sent by the output
	OutputBytes float64 `json:"outputBytes,omitempty"`

	// Current duration in milliseconds for the output
	OutputDuration float64 `json:"outputDuration,omitempty"`

	// Current formatted timecode string for the output
	OutputTimecode string `json:"outputTimecode,omitempty"`
}

// GetRecordStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetRecordStatus(paramss ...*GetRecordStatusParams) (*GetRecordStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordStatusParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetRecordStatus", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetRecordStatusResponse), nil
}
