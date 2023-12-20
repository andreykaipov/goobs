// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the StartRecord request.
type StartRecordParams struct{}

// Returns the associated request.
func (o *StartRecordParams) GetRequestName() string {
	return "StartRecord"
}

// Represents the response body for the StartRecord request.
type StartRecordResponse struct {
	_response
}

// Starts the record output.
func (c *Client) StartRecord(paramss ...*StartRecordParams) (*StartRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordParams{{}}
	}
	params := paramss[0]
	data := &StartRecordResponse{}
	return data, c.client.SendRequest(params, data)
}
