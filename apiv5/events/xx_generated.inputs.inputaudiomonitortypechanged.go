// This file has been automatically generated. Don't edit it.

package events

/*
InputAudioMonitorTypeChanged represents the event body for the "InputAudioMonitorTypeChanged" event.
Since v5.0.0.
*/
type InputAudioMonitorTypeChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// New monitor type of the input
	MonitorType string `json:"monitorType,omitempty"`
}
