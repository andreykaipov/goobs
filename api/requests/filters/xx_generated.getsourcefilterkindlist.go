// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the GetSourceFilterKindList request.
type GetSourceFilterKindListParams struct{}

// Returns the associated request.
func (o *GetSourceFilterKindListParams) GetRequestName() string {
	return "GetSourceFilterKindList"
}

// Represents the response body for the GetSourceFilterKindList request.
type GetSourceFilterKindListResponse struct {
	_response

	// Array of source filter kinds
	SourceFilterKinds []string `json:"sourceFilterKinds,omitempty"`
}

/*
Gets an array of all available source filter kinds.

Similar to `GetInputKindList`
*/
func (c *Client) GetSourceFilterKindList(
	paramss ...*GetSourceFilterKindListParams,
) (*GetSourceFilterKindListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourceFilterKindListParams{{}}
	}
	params := paramss[0]
	data := &GetSourceFilterKindListResponse{}
	return data, c.client.SendRequest(params, data)
}
