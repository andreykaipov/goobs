// This file has been automatically generated. Don't edit it.

package sources

/*
GetSourceActiveParams represents the params body for the "GetSourceActive" request.
Gets the active and show state of a source.

**Compatible with inputs and scenes.**
*/
type GetSourceActiveParams struct {
	// Name of the source to get the active state of
	SourceName string `json:"sourceName,omitempty"`
}

/*
GetSourceActiveResponse represents the response body for the "GetSourceActive" request.
Gets the active and show state of a source.

**Compatible with inputs and scenes.**
*/
type GetSourceActiveResponse struct {
	// Whether the source is showing in Program
	VideoActive bool `json:"videoActive,omitempty"`

	// Whether the source is showing in the UI (Preview, Projector, Properties)
	VideoShowing bool `json:"videoShowing,omitempty"`
}

// GetSourceActive sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceActive(params *GetSourceActiveParams) (*GetSourceActiveResponse, error) {
	resp, err := c.SendRequest("GetSourceActive", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSourceActiveResponse), nil
}
