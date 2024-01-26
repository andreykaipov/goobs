// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the OpenSourceProjector request.
type OpenSourceProjectorParams struct {
	// Monitor index, use `GetMonitorList` to obtain index
	MonitorIndex *int `json:"monitorIndex,omitempty"`

	// Size/Position data for a windowed projector, in Qt Base64 encoded format. Mutually exclusive with `monitorIndex`
	ProjectorGeometry *string `json:"projectorGeometry,omitempty"`

	// Name of the source to open a projector for
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source to open a projector for
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewOpenSourceProjectorParams() *OpenSourceProjectorParams {
	return &OpenSourceProjectorParams{}
}
func (o *OpenSourceProjectorParams) WithMonitorIndex(x int) *OpenSourceProjectorParams {
	o.MonitorIndex = &x
	return o
}
func (o *OpenSourceProjectorParams) WithProjectorGeometry(x string) *OpenSourceProjectorParams {
	o.ProjectorGeometry = &x
	return o
}
func (o *OpenSourceProjectorParams) WithSourceName(x string) *OpenSourceProjectorParams {
	o.SourceName = &x
	return o
}
func (o *OpenSourceProjectorParams) WithSourceUuid(x string) *OpenSourceProjectorParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *OpenSourceProjectorParams) GetRequestName() string {
	return "OpenSourceProjector"
}

// Represents the response body for the OpenSourceProjector request.
type OpenSourceProjectorResponse struct {
	_response
}

/*
Opens a projector for a source.

Note: This request serves to provide feature parity with 4.x. It is very likely to be changed/deprecated in a future release.
*/
func (c *Client) OpenSourceProjector(paramss ...*OpenSourceProjectorParams) (*OpenSourceProjectorResponse, error) {
	if len(paramss) == 0 {
		paramss = []*OpenSourceProjectorParams{{}}
	}
	params := paramss[0]
	data := &OpenSourceProjectorResponse{}
	return data, c.client.SendRequest(params, data)
}
