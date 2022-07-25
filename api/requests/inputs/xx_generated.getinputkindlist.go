// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputKindListParams represents the params body for the "GetInputKindList" request.
Gets an array of all available input kinds in OBS.
*/
type GetInputKindListParams struct {
	requests.ParamsBasic

	// True == Return all kinds as unversioned, False == Return with version suffixes (if available)
	Unversioned bool `json:"unversioned,omitempty"`
}

// GetSelfName just returns "GetInputKindList".
func (o *GetInputKindListParams) GetSelfName() string {
	return "GetInputKindList"
}

/*
GetInputKindListResponse represents the response body for the "GetInputKindList" request.
Gets an array of all available input kinds in OBS.
*/
type GetInputKindListResponse struct {
	requests.ResponseBasic

	// Array of input kinds
	InputKinds []string `json:"inputKinds,omitempty"`
}

// GetInputKindList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputKindList(params *GetInputKindListParams) (*GetInputKindListResponse, error) {
	data := &GetInputKindListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
