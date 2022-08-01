package events

// Error is used to wrap any errors we might encounter in our eventing loop.
type Error struct {
	Err error
}

// WrapError takes an error and wraps it in an Error event.
func WrapError(err error) *Error {
	return &Error{err}
}
