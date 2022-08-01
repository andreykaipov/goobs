package api

// Logger is a interface compatible with both the stdlib's logger and some
// third-party loggers.
type Logger interface{ Printf(string, ...interface{}) }

// LoggerWithWrite helps us anonymously satisfy a Writer interface
type LoggerWithWrite func([]byte) (int, error)

func (f LoggerWithWrite) Write(p []byte) (int, error) { return f(p) }
