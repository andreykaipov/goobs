// This file has been automatically generated. Don't edit it.

package ui

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetMonitorList request.
type GetMonitorListParams struct{}

// Returns the associated request.
func (o *GetMonitorListParams) GetRequestName() string {
	return "GetMonitorList"
}

// Represents the response body for the GetMonitorList request.
type GetMonitorListResponse struct {
	// List of detected monitors
	Monitors []*typedefs.Monitor `json:"monitors,omitempty"`
}

// Gets a list of connected monitors and information about them.
func (c *Client) GetMonitorList(paramss ...*GetMonitorListParams) (*GetMonitorListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetMonitorListParams{{}}
	}
	params := paramss[0]
	data := &GetMonitorListResponse{}
	return data, c.SendRequest(params, data)
}
