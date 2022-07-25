// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputAudioMonitorTypeParams represents the params body for the "GetInputAudioMonitorType" request.
Gets the audio monitor type of an input.

The available audio monitor types are:

- `OBS_MONITORING_TYPE_NONE`
- `OBS_MONITORING_TYPE_MONITOR_ONLY`
- `OBS_MONITORING_TYPE_MONITOR_AND_OUTPUT`
*/
type GetInputAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// Name of the input to get the audio monitor type of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputAudioMonitorType".
func (o *GetInputAudioMonitorTypeParams) GetSelfName() string {
	return "GetInputAudioMonitorType"
}

/*
GetInputAudioMonitorTypeResponse represents the response body for the "GetInputAudioMonitorType" request.
Gets the audio monitor type of an input.

The available audio monitor types are:

- `OBS_MONITORING_TYPE_NONE`
- `OBS_MONITORING_TYPE_MONITOR_ONLY`
- `OBS_MONITORING_TYPE_MONITOR_AND_OUTPUT`
*/
type GetInputAudioMonitorTypeResponse struct {
	requests.ResponseBasic

	// Audio monitor type
	MonitorType string `json:"monitorType,omitempty"`
}

// GetInputAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioMonitorType(
	params *GetInputAudioMonitorTypeParams,
) (*GetInputAudioMonitorTypeResponse, error) {
	data := &GetInputAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
