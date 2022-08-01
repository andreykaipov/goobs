// This file has been automatically generated. Don't edit it.

package filters

/*
SetSourceFilterNameParams represents the params body for the "SetSourceFilterName" request.
Sets the name of a source filter (rename).
*/
type SetSourceFilterNameParams struct {
	// Current name of the filter
	FilterName string `json:"filterName,omitempty"`

	// New name for the filter
	NewFilterName string `json:"newFilterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

/*
SetSourceFilterNameResponse represents the response body for the "SetSourceFilterName" request.
Sets the name of a source filter (rename).
*/
type SetSourceFilterNameResponse struct{}

// SetSourceFilterName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterName(params *SetSourceFilterNameParams) (*SetSourceFilterNameResponse, error) {
	resp, err := c.SendRequest("SetSourceFilterName", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSourceFilterNameResponse), nil
}
