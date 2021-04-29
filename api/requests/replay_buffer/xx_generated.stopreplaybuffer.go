// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopReplayBufferParams represents the params body for the "StopReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopReplayBuffer.
*/
type StopReplayBufferParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StopReplayBufferParams
func (o *StopReplayBufferParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StopReplayBufferParams
func (o *StopReplayBufferParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StopReplayBufferParams
func (o *StopReplayBufferParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopReplayBufferResponse represents the response body for the "StopReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopReplayBuffer.
*/
type StopReplayBufferResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StopReplayBufferResponse
func (o *StopReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StopReplayBufferResponse
func (o *StopReplayBufferResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StopReplayBufferResponse
func (o *StopReplayBufferResponse) GetError() string {
	return o.Error
}

// StopReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopReplayBuffer(
	paramss ...*StopReplayBufferParams,
) (*StopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopReplayBufferParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StopReplayBuffer"
	data := &StopReplayBufferResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
