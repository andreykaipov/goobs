// This file has been automatically generated. Don't edit it.

package record

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the ToggleRecordPause request.
type ToggleRecordPauseParams struct{}

// Returns the associated request.
func (o *ToggleRecordPauseParams) GetRequestName() string {
	return "ToggleRecordPause"
}

// Represents the response body for the ToggleRecordPause request.
type ToggleRecordPauseResponse struct {
	api.ResponseCommon
}

// Toggles pause on the record output.
func (c *Client) ToggleRecordPause(paramss ...*ToggleRecordPauseParams) (*ToggleRecordPauseResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordPauseParams{{}}
	}
	params := paramss[0]
	data := &ToggleRecordPauseResponse{}
	return data, c.SendRequest(params, data)
}
