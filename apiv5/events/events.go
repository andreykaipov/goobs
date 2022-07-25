package events

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
)

type Event interface{}

// Error is used to wrap any errors we might encounter in our eventing loop.
type Error struct {
	Err error
}

// WrapError takes an error and wraps it Error event.
func WrapError(err error) *Error {
	return &Error{
		err,
	}
}

// Parse takes a raw JSON message, first looks at the included event type to
// figure out what Go type we should return, and then we unmarshal that raw data
// into the correct Go type.
func Parse(raw json.RawMessage) (Event, error) {
	etype, err := jsonparser.GetString(raw, "eventType")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find 'eventType' key: %w", err)
	}

	switch knownEvent := GetEventForType(etype); knownEvent {
	case nil:
		return nil, nil
	default:
		data, _, _, err := jsonparser.Get(raw, "eventData")
		if err != nil {
			return nil, fmt.Errorf("Couldn't find 'eventData' key: %w", err)
		}

		if err := json.Unmarshal(data, knownEvent); err != nil {
			return nil, fmt.Errorf(
				"Couldn't unmarshal %s into an event type of %q: %s",
				raw,
				etype,
				err,
			)
		}

		return knownEvent, nil
	}
}
