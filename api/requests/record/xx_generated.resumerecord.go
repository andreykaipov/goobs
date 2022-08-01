// This file has been automatically generated. Don't edit it.

package record

/*
ResumeRecordParams represents the params body for the "ResumeRecord" request.
Resumes the record output.
*/
type ResumeRecordParams struct{}

/*
ResumeRecordResponse represents the response body for the "ResumeRecord" request.
Resumes the record output.
*/
type ResumeRecordResponse struct{}

// ResumeRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ResumeRecord(paramss ...*ResumeRecordParams) (*ResumeRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ResumeRecordParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("ResumeRecord", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ResumeRecordResponse), nil
}
