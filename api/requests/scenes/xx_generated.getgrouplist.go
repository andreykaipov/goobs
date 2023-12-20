// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetGroupList request.
type GetGroupListParams struct{}

// Returns the associated request.
func (o *GetGroupListParams) GetRequestName() string {
	return "GetGroupList"
}

// Represents the response body for the GetGroupList request.
type GetGroupListResponse struct {
	_response

	// Array of group names
	Groups []string `json:"groups,omitempty"`
}

/*
Gets an array of all groups in OBS.

Groups in OBS are actually scenes, but renamed and modified. In obs-websocket, we treat them as scenes where we can.
*/
func (c *Client) GetGroupList(paramss ...*GetGroupListParams) (*GetGroupListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetGroupListParams{{}}
	}
	params := paramss[0]
	data := &GetGroupListResponse{}
	return data, c.client.SendRequest(params, data)
}
