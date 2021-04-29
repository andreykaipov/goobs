package events

// EventBasic represents the common fields of any event.
type EventBasic struct {
	// The name of the event.
	UpdateType string `json:"update-type"`

	// Time elapsed between now and stream start (only present if OBS Studio
	// is streaming). Timecodes in format `HH:MM:SS.mmm`.
	StreamTimecode string `json:"stream-timecode,omitempty"`

	// Time elapsed between now and recording start (only present if OBS
	// Studio is recording). Timecodes in format `HH:MM:SS.mmm`.
	RecordingTimecode string `json:"rec-timecode,omitempty"`
}

// Event describes the behavior of any event. Used to abstract the functionality
// of any event that embeds EventBasic within their fields.
type Event interface {
	GetUpdateType() string
	GetStreamTimecode() string
	GetRecordingTimecode() string
}

// GetUpdateType does what it says.
func (e *EventBasic) GetUpdateType() string {
	return e.UpdateType
}

// GetStreamTimecode does what it says.
func (e *EventBasic) GetStreamTimecode() string {
	return e.StreamTimecode
}

// GetRecordingTimecode does what says.
func (e *EventBasic) GetRecordingTimecode() string {
	return e.RecordingTimecode
}

// Error is used to wrap any errors we might encounter in our eventing loop.
type Error struct {
	EventBasic
	Err error
}

// WrapError takes an error and wraps it in our custom Error event.
func WrapError(err error) *Error {
	return &Error{
		EventBasic{UpdateType: "Error"},
		err,
	}
}
