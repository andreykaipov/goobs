// This file has been automatically generated. Don't edit it.

package record

/*
StopRecordParams represents the params body for the "StopRecord" request.
Stops the record output.
*/
type StopRecordParams struct{}

/*
StopRecordResponse represents the response body for the "StopRecord" request.
Stops the record output.
*/
type StopRecordResponse struct{}

// StopRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopRecord(paramss ...*StopRecordParams) (*StopRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StopRecord", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StopRecordResponse), nil
}
