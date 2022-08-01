// This file has been automatically generated. Don't edit it.

package rconfig

/*
GetRecordDirectoryParams represents the params body for the "GetRecordDirectory" request.
Gets the current directory that the record output is set to.
*/
type GetRecordDirectoryParams struct{}

/*
GetRecordDirectoryResponse represents the response body for the "GetRecordDirectory" request.
Gets the current directory that the record output is set to.
*/
type GetRecordDirectoryResponse struct {
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
	resp, err := c.SendRequest("GetRecordDirectory", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetRecordDirectoryResponse), nil
}
