// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopReplayBufferParams represents the params body for the "StartStopReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopReplayBuffer.
*/
type StartStopReplayBufferParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartStopReplayBufferParams
func (o *StartStopReplayBufferParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStopReplayBufferParams
func (o *StartStopReplayBufferParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStopReplayBufferParams
func (o *StartStopReplayBufferParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopReplayBufferResponse represents the response body for the "StartStopReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopReplayBuffer.
*/
type StartStopReplayBufferResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStopReplayBufferResponse
func (o *StartStopReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStopReplayBufferResponse
func (o *StartStopReplayBufferResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStopReplayBufferResponse
func (o *StartStopReplayBufferResponse) GetError() string {
	return o.Error
}

// StartStopReplayBuffer sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopReplayBuffer(
	paramss ...*StartStopReplayBufferParams,
) (*StartStopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopReplayBufferParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartStopReplayBuffer"
	data := &StartStopReplayBufferResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
