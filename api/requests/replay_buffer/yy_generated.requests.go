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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SaveReplayBufferParams represents the params body for the "SaveReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveReplayBuffer.
*/
type SaveReplayBufferParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of SaveReplayBufferParams
func (o *SaveReplayBufferParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SaveReplayBufferParams
func (o *SaveReplayBufferParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SaveReplayBufferParams
func (o *SaveReplayBufferParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SaveReplayBufferResponse represents the response body for the "SaveReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveReplayBuffer.
*/
type SaveReplayBufferResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SaveReplayBufferResponse
func (o *SaveReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SaveReplayBufferResponse
func (o *SaveReplayBufferResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SaveReplayBufferResponse
func (o *SaveReplayBufferResponse) GetError() string {
	return o.Error
}

// SaveReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) SaveReplayBuffer(
	paramss ...*SaveReplayBufferParams,
) (*SaveReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SaveReplayBufferParams{{}}
	}
	params := paramss[0]
	params.RequestType = "SaveReplayBuffer"
	data := &SaveReplayBufferResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
