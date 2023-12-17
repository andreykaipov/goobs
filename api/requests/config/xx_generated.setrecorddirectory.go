// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetRecordDirectory request.
type SetRecordDirectoryParams struct {
	// Output directory
	RecordDirectory string `json:"recordDirectory,omitempty"`
}

// Returns the associated request.
func (o *SetRecordDirectoryParams) GetRequestName() string {
	return "SetRecordDirectory"
}

// Represents the response body for the SetRecordDirectory request.
type SetRecordDirectoryResponse struct{}

// Sets the current directory that the record output writes files to.
func (c *Client) SetRecordDirectory(params *SetRecordDirectoryParams) (*SetRecordDirectoryResponse, error) {
	data := &SetRecordDirectoryResponse{}
	return data, c.SendRequest(params, data)
}
