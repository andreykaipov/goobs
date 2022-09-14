// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenSourceProjector request.
type OpenSourceProjectorParams struct {
	// Monitor index, use `GetMonitorList` to obtain index
	MonitorIndex float64 `json:"monitorIndex,omitempty"`

	// Size/Position data for a windowed projector, in Qt Base64 encoded format. Mutually exclusive with `monitorIndex`
	ProjectorGeometry string `json:"projectorGeometry,omitempty"`

	// Name of the source to open a projector for
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *OpenSourceProjectorParams) GetRequestName() string {
	return "OpenSourceProjector"
}

// Represents the response body for the OpenSourceProjector request.
type OpenSourceProjectorResponse struct{}

/*
Opens a projector for a source.

Note: This request serves to provide feature parity with 4.x. It is very likely to be changed/deprecated in a future release.
*/
func (c *Client) OpenSourceProjector(params *OpenSourceProjectorParams) (*OpenSourceProjectorResponse, error) {
	data := &OpenSourceProjectorResponse{}
	return data, c.SendRequest(params, data)
}
