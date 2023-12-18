// This file has been automatically generated. Don't edit it.

package requests

func GetStatusForCode(code int) string {
	switch code {
	case 0:
		// Unknown status, should never be used.
		return "Unknown"
	case 10:
		// For internal use to signify a successful field check.
		return "NoError"
	case 100:
		// The request has succeeded.
		return "Success"
	case 203:
		// The `requestType` field is missing from the request data.
		return "MissingRequestType"
	case 204:
		// The request type is invalid or does not exist.
		return "UnknownRequestType"
	case 205:
		/*
		   Generic error code.

		   Note: A comment is required to be provided by obs-websocket.
		*/
		return "GenericError"
	case 206:
		// The request batch execution type is not supported.
		return "UnsupportedRequestBatchExecutionType"
	case 300:
		// A required request field is missing.
		return "MissingRequestField"
	case 301:
		// The request does not have a valid requestData object.
		return "MissingRequestData"
	case 400:
		/*
		   Generic invalid request field message.

		   Note: A comment is required to be provided by obs-websocket.
		*/
		return "InvalidRequestField"
	case 401:
		// A request field has the wrong data type.
		return "InvalidRequestFieldType"
	case 402:
		// A request field (number) is outside of the allowed range.
		return "RequestFieldOutOfRange"
	case 403:
		// A request field (string or array) is empty and cannot be.
		return "RequestFieldEmpty"
	case 404:
		// There are too many request fields (eg. a request takes two optionals, where only one is allowed at a time).
		return "TooManyRequestFields"
	case 500:
		// An output is running and cannot be in order to perform the request.
		return "OutputRunning"
	case 501:
		// An output is not running and should be.
		return "OutputNotRunning"
	case 502:
		// An output is paused and should not be.
		return "OutputPaused"
	case 503:
		// An output is not paused and should be.
		return "OutputNotPaused"
	case 504:
		// An output is disabled and should not be.
		return "OutputDisabled"
	case 505:
		// Studio mode is active and cannot be.
		return "StudioModeActive"
	case 506:
		// Studio mode is not active and should be.
		return "StudioModeNotActive"
	case 600:
		/*
		   The resource was not found.

		   Note: Resources are any kind of object in obs-websocket, like inputs, profiles, outputs, etc.
		*/
		return "ResourceNotFound"
	case 601:
		// The resource already exists.
		return "ResourceAlreadyExists"
	case 602:
		// The type of resource found is invalid.
		return "InvalidResourceType"
	case 603:
		// There are not enough instances of the resource in order to perform the request.
		return "NotEnoughResources"
	case 604:
		// The state of the resource is invalid. For example, if the resource is blocked from being accessed.
		return "InvalidResourceState"
	case 605:
		// The specified input (obs_source_t-OBS_SOURCE_TYPE_INPUT) had the wrong kind.
		return "InvalidInputKind"
	case 606:
		/*
		   The resource does not support being configured.

		   This is particularly relevant to transitions, where they do not always have changeable settings.
		*/
		return "ResourceNotConfigurable"
	case 607:
		// The specified filter (obs_source_t-OBS_SOURCE_TYPE_FILTER) had the wrong kind.
		return "InvalidFilterKind"
	case 700:
		// Creating the resource failed.
		return "ResourceCreationFailed"
	case 701:
		// Performing an action on the resource failed.
		return "ResourceActionFailed"
	case 702:
		/*
		   Processing the request failed unexpectedly.

		   Note: A comment is required to be provided by obs-websocket.
		*/
		return "RequestProcessingFailed"
	case 703:
		// The combination of request fields cannot be used to perform an action.
		return "CannotAct"
	default:
		return "EvenMoreUnknown"
	}
}
