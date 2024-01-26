// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputAudioMonitorTypeChanged event.

The monitor type of an input has changed.

Available types are:

- `OBS_MONITORING_TYPE_NONE`
- `OBS_MONITORING_TYPE_MONITOR_ONLY`
- `OBS_MONITORING_TYPE_MONITOR_AND_OUTPUT`
*/
type InputAudioMonitorTypeChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`

	// New monitor type of the input
	MonitorType string `json:"monitorType,omitempty"`
}
