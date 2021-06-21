// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetAudioMonitorTypeParams represents the params body for the "SetAudioMonitorType" request.
Set the audio monitoring type of the specified source.
Since 4.8.0.
*/
type SetAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// The monitor type to use. Options: `none`, `monitorOnly`, `monitorAndOutput`.
	MonitorType string `json:"monitorType,omitempty"`

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetAudioMonitorType".
func (o *SetAudioMonitorTypeParams) GetSelfName() string {
	return "SetAudioMonitorType"
}

/*
SetAudioMonitorTypeResponse represents the response body for the "SetAudioMonitorType" request.
Set the audio monitoring type of the specified source.
Since v4.8.0.
*/
type SetAudioMonitorTypeResponse struct {
	requests.ResponseBasic
}

// SetAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetAudioMonitorType(params *SetAudioMonitorTypeParams) (*SetAudioMonitorTypeResponse, error) {
	data := &SetAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
