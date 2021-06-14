// This file has been automatically generated. Don't edit it.

package events

/*
SourceAudioMixersChanged represents the event body for the "SourceAudioMixersChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SourceAudioMixersChanged.
*/
type SourceAudioMixersChanged struct {
	EventBasic

	// Raw mixer flags (little-endian, one bit per mixer) as an hexadecimal value
	HexMixersValue string `json:"hexMixersValue"`

	RoutingStatus []struct {
		// Routing status
		Enabled bool `json:"enabled"`

		// Mixer number
		Id int `json:"id"`
	} `json:"routingStatus"`

	// Source name
	SourceName string `json:"sourceName"`
}
