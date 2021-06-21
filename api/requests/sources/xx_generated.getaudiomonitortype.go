// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAudioMonitorTypeParams represents the params body for the "GetAudioMonitorType" request.
Get the audio monitoring type of the specified source.
Since 4.8.0.
*/
type GetAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetAudioMonitorType".
func (o *GetAudioMonitorTypeParams) GetSelfName() string {
	return "GetAudioMonitorType"
}

/*
GetAudioMonitorTypeResponse represents the response body for the "GetAudioMonitorType" request.
Get the audio monitoring type of the specified source.
Since v4.8.0.
*/
type GetAudioMonitorTypeResponse struct {
	requests.ResponseBasic

	// The monitor type in use. Options: `none`, `monitorOnly`, `monitorAndOutput`.
	MonitorType string `json:"monitorType,omitempty"`
}

// GetAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetAudioMonitorType(params *GetAudioMonitorTypeParams) (*GetAudioMonitorTypeResponse, error) {
	data := &GetAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
