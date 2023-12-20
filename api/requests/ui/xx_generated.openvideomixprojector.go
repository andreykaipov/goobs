// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenVideoMixProjector request.
type OpenVideoMixProjectorParams struct {
	// Monitor index, use `GetMonitorList` to obtain index
	MonitorIndex *int `json:"monitorIndex,omitempty"`

	// Size/Position data for a windowed projector, in Qt Base64 encoded format. Mutually exclusive with `monitorIndex`
	ProjectorGeometry *string `json:"projectorGeometry,omitempty"`

	// Type of mix to open
	VideoMixType *string `json:"videoMixType,omitempty"`
}

func NewOpenVideoMixProjectorParams() *OpenVideoMixProjectorParams {
	return &OpenVideoMixProjectorParams{}
}
func (o *OpenVideoMixProjectorParams) WithMonitorIndex(x int) *OpenVideoMixProjectorParams {
	o.MonitorIndex = &x
	return o
}
func (o *OpenVideoMixProjectorParams) WithProjectorGeometry(x string) *OpenVideoMixProjectorParams {
	o.ProjectorGeometry = &x
	return o
}
func (o *OpenVideoMixProjectorParams) WithVideoMixType(x string) *OpenVideoMixProjectorParams {
	o.VideoMixType = &x
	return o
}

// Returns the associated request.
func (o *OpenVideoMixProjectorParams) GetRequestName() string {
	return "OpenVideoMixProjector"
}

// Represents the response body for the OpenVideoMixProjector request.
type OpenVideoMixProjectorResponse struct {
	_response
}

/*
Opens a projector for a specific output video mix.

Mix types:

- `OBS_WEBSOCKET_VIDEO_MIX_TYPE_PREVIEW`
- `OBS_WEBSOCKET_VIDEO_MIX_TYPE_PROGRAM`
- `OBS_WEBSOCKET_VIDEO_MIX_TYPE_MULTIVIEW`

Note: This request serves to provide feature parity with 4.x. It is very likely to be changed/deprecated in a future release.
*/
func (c *Client) OpenVideoMixProjector(params *OpenVideoMixProjectorParams) (*OpenVideoMixProjectorResponse, error) {
	data := &OpenVideoMixProjectorResponse{}
	return data, c.client.SendRequest(params, data)
}
