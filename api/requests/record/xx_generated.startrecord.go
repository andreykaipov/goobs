// This file has been automatically generated. Don't edit it.

package record

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartRecordParams represents the params body for the "StartRecord" request.
Starts the record output.
*/
type StartRecordParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartRecord".
func (o *StartRecordParams) GetSelfName() string {
	return "StartRecord"
}

/*
StartRecordResponse represents the response body for the "StartRecord" request.
Starts the record output.
*/
type StartRecordResponse struct {
	requests.ResponseBasic
}

// StartRecord sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StartRecord(paramss ...*StartRecordParams) (*StartRecordResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordParams{{}}
	}
	params := paramss[0]
	data := &StartRecordResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
