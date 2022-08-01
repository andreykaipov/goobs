// This file has been automatically generated. Don't edit it.

package record

/*
ToggleRecordPauseParams represents the params body for the "ToggleRecordPause" request.
Toggles pause on the record output.
*/
type ToggleRecordPauseParams struct{}

/*
ToggleRecordPauseResponse represents the response body for the "ToggleRecordPause" request.
Toggles pause on the record output.
*/
type ToggleRecordPauseResponse struct{}

// ToggleRecordPause sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) ToggleRecordPause(paramss ...*ToggleRecordPauseParams) (*ToggleRecordPauseResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordPauseParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("ToggleRecordPause", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ToggleRecordPauseResponse), nil
}
