// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleRecordPauseParams represents the params body for the "ToggleRecordPause" request.
Toggles pause on the record output.
*/
type ToggleRecordPauseParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleRecordPause".
func (o *ToggleRecordPauseParams) GetSelfName() string {
	return "ToggleRecordPause"
}

/*
ToggleRecordPauseResponse represents the response body for the "ToggleRecordPause" request.
Toggles pause on the record output.
*/
type ToggleRecordPauseResponse struct {
	requests.ResponseBasic
}

// ToggleRecordPause sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) ToggleRecordPause(paramss ...*ToggleRecordPauseParams) (*ToggleRecordPauseResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordPauseParams{{}}
	}
	params := paramss[0]
	data := &ToggleRecordPauseResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
