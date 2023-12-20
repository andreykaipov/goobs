// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the PauseRecord request.
type PauseRecordParams struct{}

// Returns the associated request.
func (o *PauseRecordParams) GetRequestName() string {
	return "PauseRecord"
}

// Represents the response body for the PauseRecord request.
type PauseRecordResponse struct {
	_response
}

// Pauses the record output.
func (c *Client) PauseRecord(paramss ...*PauseRecordParams) (*PauseRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*PauseRecordParams{{}}
	}
	params := paramss[0]
	data := &PauseRecordResponse{}
	return data, c.client.SendRequest(params, data)
}
