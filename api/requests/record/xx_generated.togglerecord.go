// This file has been automatically generated. Don't edit it.

package record

/*
ToggleRecordParams represents the params body for the "ToggleRecord" request.
Toggles the status of the record output.
*/
type ToggleRecordParams struct{}

/*
ToggleRecordResponse represents the response body for the "ToggleRecord" request.
Toggles the status of the record output.
*/
type ToggleRecordResponse struct{}

// ToggleRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ToggleRecord(paramss ...*ToggleRecordParams) (*ToggleRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("ToggleRecord", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ToggleRecordResponse), nil
}
