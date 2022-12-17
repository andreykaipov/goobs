// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the StopRecord request.
type StopRecordParams struct{}

// Returns the associated request.
func (o *StopRecordParams) GetRequestName() string {
	return "StopRecord"
}

// Represents the response body for the StopRecord request.
type StopRecordResponse struct {
	// File name for the saved recording
	OutputPath string `json:"outputPath,omitempty"`
}

// Stops the record output.
func (c *Client) StopRecord(paramss ...*StopRecordParams) (*StopRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordParams{{}}
	}
	params := paramss[0]
	data := &StopRecordResponse{}
	return data, c.SendRequest(params, data)
}
