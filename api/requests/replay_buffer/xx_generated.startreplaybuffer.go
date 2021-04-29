// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartReplayBuffer.
*/
type StartReplayBufferParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartReplayBufferParams
func (o *StartReplayBufferParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartReplayBufferParams
func (o *StartReplayBufferParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartReplayBufferParams
func (o *StartReplayBufferParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartReplayBufferResponse represents the response body for the "StartReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartReplayBuffer.
*/
type StartReplayBufferResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartReplayBufferResponse
func (o *StartReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartReplayBufferResponse
func (o *StartReplayBufferResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartReplayBufferResponse
func (o *StartReplayBufferResponse) GetError() string {
	return o.Error
}

// StartReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartReplayBuffer(
	paramss ...*StartReplayBufferParams,
) (*StartReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartReplayBufferParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartReplayBuffer"
	data := &StartReplayBufferResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
