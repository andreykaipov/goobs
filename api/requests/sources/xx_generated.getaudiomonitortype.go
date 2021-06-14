// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAudioMonitorTypeParams represents the params body for the "GetAudioMonitorType" request.
Get the audio monitoring type of the specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetAudioMonitorType.
*/
type GetAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName"`
}

// Name just returns "GetAudioMonitorType".
func (o *GetAudioMonitorTypeParams) Name() string {
	return "GetAudioMonitorType"
}

/*
GetAudioMonitorTypeResponse represents the response body for the "GetAudioMonitorType" request.
Get the audio monitoring type of the specified source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetAudioMonitorType.
*/
type GetAudioMonitorTypeResponse struct {
	requests.ResponseBasic

	// The monitor type in use. Options: `none`, `monitorOnly`, `monitorAndOutput`.
	MonitorType string `json:"monitorType"`
}

// GetAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetAudioMonitorType(
	params *GetAudioMonitorTypeParams,
) (*GetAudioMonitorTypeResponse, error) {
	data := &GetAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
