// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
OpenProjectorParams represents the params body for the "OpenProjector" request.
Open a projector window or create a projector on a monitor. Requires OBS v24.0.4 or newer.
Since 4.8.0.
*/
type OpenProjectorParams struct {
	requests.ParamsBasic

	// Size and position of the projector window (only if monitor is -1). Encoded in Base64 using [Qt's geometry
	// encoding](https://doc.qt.io/qt-5/qwidget.html#saveGeometry). Corresponds to OBS's saved projectors.
	Geometry string `json:"geometry,omitempty"`

	// Monitor to open the projector on. If -1 or omitted, opens a window.
	Monitor int `json:"monitor,omitempty"`

	// Name of the source or scene to be displayed (ignored for other projector types).
	Name string `json:"name,omitempty"`

	// Type of projector: `Preview` (default), `Source`, `Scene`, `StudioProgram`, or `Multiview` (case insensitive).
	Type string `json:"type,omitempty"`
}

// GetSelfName just returns "OpenProjector".
func (o *OpenProjectorParams) GetSelfName() string {
	return "OpenProjector"
}

/*
OpenProjectorResponse represents the response body for the "OpenProjector" request.
Open a projector window or create a projector on a monitor. Requires OBS v24.0.4 or newer.
Since v4.8.0.
*/
type OpenProjectorResponse struct {
	requests.ResponseBasic
}

// OpenProjector sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OpenProjector(params *OpenProjectorParams) (*OpenProjectorResponse, error) {
	data := &OpenProjectorResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
