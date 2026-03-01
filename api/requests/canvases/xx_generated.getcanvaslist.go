// This file has been automatically generated. Don't edit it.

package canvases

// Represents the request body for the GetCanvasList request.
type GetCanvasListParams struct{}

// Returns the associated request.
func (o *GetCanvasListParams) GetRequestName() string {
	return "GetCanvasList"
}

// Represents the response body for the GetCanvasList request.
type GetCanvasListResponse struct {
	_response

	// Array of canvases
	Canvases []map[string]any `json:"canvases,omitempty"`
}

// Gets an array of canvases in OBS.
func (c *Client) GetCanvasList(paramss ...*GetCanvasListParams) (*GetCanvasListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCanvasListParams{{}}
	}
	params := paramss[0]
	data := &GetCanvasListResponse{}
	return data, c.client.SendRequest(params, data)
}
