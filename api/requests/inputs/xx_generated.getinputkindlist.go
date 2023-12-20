// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputKindList request.
type GetInputKindListParams struct {
	// True == Return all kinds as unversioned, False == Return with version suffixes (if available)
	Unversioned *bool `json:"unversioned,omitempty"`
}

func NewGetInputKindListParams() *GetInputKindListParams {
	return &GetInputKindListParams{}
}
func (o *GetInputKindListParams) WithUnversioned(x bool) *GetInputKindListParams {
	o.Unversioned = &x
	return o
}

// Returns the associated request.
func (o *GetInputKindListParams) GetRequestName() string {
	return "GetInputKindList"
}

// Represents the response body for the GetInputKindList request.
type GetInputKindListResponse struct {
	_response

	// Array of input kinds
	InputKinds []string `json:"inputKinds,omitempty"`
}

// Gets an array of all available input kinds in OBS.
func (c *Client) GetInputKindList(paramss ...*GetInputKindListParams) (*GetInputKindListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputKindListParams{{}}
	}
	params := paramss[0]
	data := &GetInputKindListResponse{}
	return data, c.client.SendRequest(params, data)
}
