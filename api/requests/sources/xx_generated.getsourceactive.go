// This file has been automatically generated. Don't edit it.

package sources

// Represents the request body for the GetSourceActive request.
type GetSourceActiveParams struct {
	// Name of the source to get the active state of
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source to get the active state of
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewGetSourceActiveParams() *GetSourceActiveParams {
	return &GetSourceActiveParams{}
}
func (o *GetSourceActiveParams) WithSourceName(x string) *GetSourceActiveParams {
	o.SourceName = &x
	return o
}
func (o *GetSourceActiveParams) WithSourceUuid(x string) *GetSourceActiveParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSourceActiveParams) GetRequestName() string {
	return "GetSourceActive"
}

// Represents the response body for the GetSourceActive request.
type GetSourceActiveResponse struct {
	_response

	// Whether the source is showing in Program
	VideoActive bool `json:"videoActive,omitempty"`

	// Whether the source is showing in the UI (Preview, Projector, Properties)
	VideoShowing bool `json:"videoShowing,omitempty"`
}

/*
Gets the active and show state of a source.

**Compatible with inputs and scenes.**
*/
func (c *Client) GetSourceActive(paramss ...*GetSourceActiveParams) (*GetSourceActiveResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourceActiveParams{{}}
	}
	params := paramss[0]
	data := &GetSourceActiveResponse{}
	return data, c.client.SendRequest(params, data)
}
