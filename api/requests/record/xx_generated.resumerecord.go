// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResumeRecordParams represents the params body for the "ResumeRecord" request.
Resumes the record output.
*/
type ResumeRecordParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ResumeRecord".
func (o *ResumeRecordParams) GetSelfName() string {
	return "ResumeRecord"
}

/*
ResumeRecordResponse represents the response body for the "ResumeRecord" request.
Resumes the record output.
*/
type ResumeRecordResponse struct {
	requests.ResponseBasic
}

// ResumeRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ResumeRecord(paramss ...*ResumeRecordParams) (*ResumeRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ResumeRecordParams{{}}
	}
	params := paramss[0]
	data := &ResumeRecordResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
