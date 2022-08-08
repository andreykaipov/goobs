// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the ToggleRecord request.
type ToggleRecordParams struct{}

// Returns the associated request.
func (o *ToggleRecordParams) GetRequestName() string {
	return "ToggleRecord"
}

// Represents the response body for the ToggleRecord request.
type ToggleRecordResponse struct{}

// Toggles the status of the record output.
func (c *Client) ToggleRecord(paramss ...*ToggleRecordParams) (*ToggleRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordParams{{}}
	}
	params := paramss[0]
	data := &ToggleRecordResponse{}
	return data, c.SendRequest(params, data)
}
