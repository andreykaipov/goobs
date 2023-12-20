// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the ResumeRecord request.
type ResumeRecordParams struct{}

// Returns the associated request.
func (o *ResumeRecordParams) GetRequestName() string {
	return "ResumeRecord"
}

// Represents the response body for the ResumeRecord request.
type ResumeRecordResponse struct {
	_response
}

// Resumes the record output.
func (c *Client) ResumeRecord(paramss ...*ResumeRecordParams) (*ResumeRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ResumeRecordParams{{}}
	}
	params := paramss[0]
	data := &ResumeRecordResponse{}
	return data, c.client.SendRequest(params, data)
}
