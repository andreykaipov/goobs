// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetGroupListParams represents the params body for the "GetGroupList" request.
Gets an array of all groups in OBS.

Groups in OBS are actually scenes, but renamed and modified. In obs-websocket, we treat them as scenes where we can.
*/
type GetGroupListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetGroupList".
func (o *GetGroupListParams) GetSelfName() string {
	return "GetGroupList"
}

/*
GetGroupListResponse represents the response body for the "GetGroupList" request.
Gets an array of all groups in OBS.

Groups in OBS are actually scenes, but renamed and modified. In obs-websocket, we treat them as scenes where we can.
*/
type GetGroupListResponse struct {
	requests.ResponseBasic

	// Array of group names
	Groups []string `json:"groups,omitempty"`
}

// GetGroupList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetGroupList(paramss ...*GetGroupListParams) (*GetGroupListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetGroupListParams{{}}
	}
	params := paramss[0]
	data := &GetGroupListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
