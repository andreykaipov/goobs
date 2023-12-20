// This file has been automatically generated. Don't edit it.

package closecodes

const (
	// For internal use only to tell the request handler not to perform any close action.
	DontClose = 0

	// Unknown reason, should never be used.
	UnknownReason = 4000

	// The server was unable to decode the incoming websocket message.
	MessageDecodeError = 4002

	// A data field is required but missing from the payload.
	MissingDataField = 4003

	// A data field's value type is invalid.
	InvalidDataFieldType = 4004

	// A data field's value is invalid.
	InvalidDataFieldValue = 4005

	// The specified `op` was invalid or missing.
	UnknownOpCode = 4006

	// The client sent a websocket message without first sending `Identify` message.
	NotIdentified = 4007

	/*
	   The client sent an `Identify` message while already identified.

	   Note: Once a client has identified, only `Reidentify` may be used to change session parameters.
	*/
	AlreadyIdentified = 4008

	// The authentication attempt (via `Identify`) failed.
	AuthenticationFailed = 4009

	// The server detected the usage of an old version of the obs-websocket RPC protocol.
	UnsupportedRpcVersion = 4010

	/*
	   The websocket session has been invalidated by the obs-websocket server.

	   Note: This is the code used by the `Kick` button in the UI Session List. If you receive this code, you must not automatically reconnect.
	*/
	SessionInvalidated = 4011

	// A requested feature is not supported due to hardware/software limitations.
	UnsupportedFeature = 4012
)
