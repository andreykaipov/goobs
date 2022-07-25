// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
PauseRecordParams represents the params body for the "PauseRecord" request.
Pauses the record output.
*/
type PauseRecordParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "PauseRecord".
func (o *PauseRecordParams) GetSelfName() string {
	return "PauseRecord"
}

/*
PauseRecordResponse represents the response body for the "PauseRecord" request.
Pauses the record output.
*/
type PauseRecordResponse struct {
	requests.ResponseBasic
}

// PauseRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) PauseRecord(paramss ...*PauseRecordParams) (*PauseRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*PauseRecordParams{{}}
	}
	params := paramss[0]
	data := &PauseRecordResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
