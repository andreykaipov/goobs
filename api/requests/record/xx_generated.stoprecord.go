// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopRecordParams represents the params body for the "StopRecord" request.
Stops the record output.
*/
type StopRecordParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopRecord".
func (o *StopRecordParams) GetSelfName() string {
	return "StopRecord"
}

/*
StopRecordResponse represents the response body for the "StopRecord" request.
Stops the record output.
*/
type StopRecordResponse struct {
	requests.ResponseBasic
}

// StopRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopRecord(paramss ...*StopRecordParams) (*StopRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordParams{{}}
	}
	params := paramss[0]
	data := &StopRecordResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
