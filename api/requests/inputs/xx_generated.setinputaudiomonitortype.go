// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputAudioMonitorTypeParams represents the params body for the "SetInputAudioMonitorType" request.
Sets the audio monitor type of an input.
*/
type SetInputAudioMonitorTypeParams struct {
	requests.ParamsBasic

	// Name of the input to set the audio monitor type of
	InputName string `json:"inputName,omitempty"`

	// Audio monitor type
	MonitorType string `json:"monitorType,omitempty"`
}

// GetSelfName just returns "SetInputAudioMonitorType".
func (o *SetInputAudioMonitorTypeParams) GetSelfName() string {
	return "SetInputAudioMonitorType"
}

/*
SetInputAudioMonitorTypeResponse represents the response body for the "SetInputAudioMonitorType" request.
Sets the audio monitor type of an input.
*/
type SetInputAudioMonitorTypeResponse struct {
	requests.ResponseBasic
}

// SetInputAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioMonitorType(
	params *SetInputAudioMonitorTypeParams,
) (*SetInputAudioMonitorTypeResponse, error) {
	data := &SetInputAudioMonitorTypeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
