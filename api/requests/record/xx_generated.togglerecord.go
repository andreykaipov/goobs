// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleRecordParams represents the params body for the "ToggleRecord" request.
Toggles the status of the record output.
*/
type ToggleRecordParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleRecord".
func (o *ToggleRecordParams) GetSelfName() string {
	return "ToggleRecord"
}

/*
ToggleRecordResponse represents the response body for the "ToggleRecord" request.
Toggles the status of the record output.
*/
type ToggleRecordResponse struct {
	requests.ResponseBasic
}

// ToggleRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ToggleRecord(paramss ...*ToggleRecordParams) (*ToggleRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleRecordParams{{}}
	}
	params := paramss[0]
	data := &ToggleRecordResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
