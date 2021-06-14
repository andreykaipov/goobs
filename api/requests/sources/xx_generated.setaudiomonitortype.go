// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetAudioMonitorTypeParams represents the params body for the "SetAudioMonitorType" request.
Set the audio monitoring type of the specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetAudioMonitorType.
*/
type SetAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// The monitor type to use. Options: `none`, `monitorOnly`, `monitorAndOutput`.
	MonitorType string `json:"monitorType"`

	// Source name.
	SourceName string `json:"sourceName"`
}

// GetSelfName just returns "SetAudioMonitorType".
func (o *SetAudioMonitorTypeParams) GetSelfName() string {
	return "SetAudioMonitorType"
}

/*
SetAudioMonitorTypeResponse represents the response body for the "SetAudioMonitorType" request.
Set the audio monitoring type of the specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetAudioMonitorType.
*/
type SetAudioMonitorTypeResponse struct {
	requests.ResponseBasic
}

// SetAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetAudioMonitorType(
	params *SetAudioMonitorTypeParams,
) (*SetAudioMonitorTypeResponse, error) {
	data := &SetAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
