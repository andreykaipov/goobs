// This file has been automatically generated. Don't edit it.

package rconfig

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetRecordDirectoryParams represents the params body for the "GetRecordDirectory" request.
Gets the current directory that the record output is set to.
*/
type GetRecordDirectoryParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetRecordDirectory".
func (o *GetRecordDirectoryParams) GetSelfName() string {
	return "GetRecordDirectory"
}

/*
GetRecordDirectoryResponse represents the response body for the "GetRecordDirectory" request.
Gets the current directory that the record output is set to.
*/
type GetRecordDirectoryResponse struct {
	requests.ResponseBasic

	// Output directory
	RecordDirectory string `json:"recordDirectory,omitempty"`
}

// GetRecordDirectory sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetRecordDirectory(paramss ...*GetRecordDirectoryParams) (*GetRecordDirectoryResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordDirectoryParams{{}}
	}
	params := paramss[0]
	data := &GetRecordDirectoryResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
