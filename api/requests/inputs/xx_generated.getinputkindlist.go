// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputKindListParams represents the params body for the "GetInputKindList" request.
Gets an array of all available input kinds in OBS.
*/
type GetInputKindListParams struct {
	// True == Return all kinds as unversioned, False == Return with version suffixes (if available)
	Unversioned *bool `json:"unversioned,omitempty"`
}

/*
GetInputKindListResponse represents the response body for the "GetInputKindList" request.
Gets an array of all available input kinds in OBS.
*/
type GetInputKindListResponse struct {
	// Array of input kinds
	InputKinds []string `json:"inputKinds,omitempty"`
}

// GetInputKindList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetInputKindList(paramss ...*GetInputKindListParams) (*GetInputKindListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputKindListParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetInputKindList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputKindListResponse), nil
}
