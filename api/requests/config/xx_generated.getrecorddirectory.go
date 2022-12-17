// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the GetRecordDirectory request.
type GetRecordDirectoryParams struct{}

// Returns the associated request.
func (o *GetRecordDirectoryParams) GetRequestName() string {
	return "GetRecordDirectory"
}

// Represents the response body for the GetRecordDirectory request.
type GetRecordDirectoryResponse struct {
	// Output directory
	RecordDirectory string `json:"recordDirectory,omitempty"`
}

// Gets the current directory that the record output is set to.
func (c *Client) GetRecordDirectory(paramss ...*GetRecordDirectoryParams) (*GetRecordDirectoryResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordDirectoryParams{{}}
	}
	params := paramss[0]
	data := &GetRecordDirectoryResponse{}
	return data, c.SendRequest(params, data)
}
