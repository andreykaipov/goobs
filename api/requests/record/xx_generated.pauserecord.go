// This file has been automatically generated. Don't edit it.

package record

/*
PauseRecordParams represents the params body for the "PauseRecord" request.
Pauses the record output.
*/
type PauseRecordParams struct{}

/*
PauseRecordResponse represents the response body for the "PauseRecord" request.
Pauses the record output.
*/
type PauseRecordResponse struct{}

// PauseRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) PauseRecord(paramss ...*PauseRecordParams) (*PauseRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*PauseRecordParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("PauseRecord", params)
	if err != nil {
		return nil, err
	}
	return resp.(*PauseRecordResponse), nil
}
