// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the SplitRecordFile request.
type SplitRecordFileParams struct{}

// Returns the associated request.
func (o *SplitRecordFileParams) GetRequestName() string {
	return "SplitRecordFile"
}

// Represents the response body for the SplitRecordFile request.
type SplitRecordFileResponse struct {
	_response
}

// Splits the current file being recorded into a new file.
func (c *Client) SplitRecordFile(paramss ...*SplitRecordFileParams) (*SplitRecordFileResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SplitRecordFileParams{{}}
	}
	params := paramss[0]
	data := &SplitRecordFileResponse{}
	return data, c.client.SendRequest(params, data)
}
