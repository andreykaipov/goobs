// This file has been automatically generated. Don't edit it.

package filters

/*
GetSourceFilterListParams represents the params body for the "GetSourceFilterList" request.
Gets an array of all of a source's filters.
*/
type GetSourceFilterListParams struct {
	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

/*
GetSourceFilterListResponse represents the response body for the "GetSourceFilterList" request.
Gets an array of all of a source's filters.
*/
type GetSourceFilterListResponse struct {
	// Array of filters
	Filters []interface{} `json:"filters,omitempty"`
}

// GetSourceFilterList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterList(params *GetSourceFilterListParams) (*GetSourceFilterListResponse, error) {
	resp, err := c.SendRequest("GetSourceFilterList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSourceFilterListResponse), nil
}
