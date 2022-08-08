// This file has been automatically generated. Don't edit it.

package sources

// Represents the request body for the GetSourceActive request.
type GetSourceActiveParams struct {
	// Name of the source to get the active state of
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *GetSourceActiveParams) GetRequestName() string {
	return "GetSourceActive"
}

// Represents the response body for the GetSourceActive request.
type GetSourceActiveResponse struct {
	// Whether the source is showing in Program
	VideoActive bool `json:"videoActive,omitempty"`

	// Whether the source is showing in the UI (Preview, Projector, Properties)
	VideoShowing bool `json:"videoShowing,omitempty"`
}

/*
Gets the active and show state of a source.

**Compatible with inputs and scenes.**
*/
func (c *Client) GetSourceActive(params *GetSourceActiveParams) (*GetSourceActiveResponse, error) {
	data := &GetSourceActiveResponse{}
	return data, c.SendRequest(params, data)
}
