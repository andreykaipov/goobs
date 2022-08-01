// This file has been automatically generated. Don't edit it.

package record

/*
StartRecordParams represents the params body for the "StartRecord" request.
Starts the record output.
*/
type StartRecordParams struct{}

/*
StartRecordResponse represents the response body for the "StartRecord" request.
Starts the record output.
*/
type StartRecordResponse struct{}

// StartRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StartRecord(paramss ...*StartRecordParams) (*StartRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StartRecord", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StartRecordResponse), nil
}
