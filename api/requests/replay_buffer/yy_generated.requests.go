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

func (o *StartStopReplayBufferParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StartStopReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartStopReplayBufferResponse) GetStatus() string {
	return o.Status
}
func (o *StartStopReplayBufferResponse) GetError() string {
	return o.Error
}

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartReplayBuffer.
*/
type StartReplayBufferParams struct {
	requests.Params
}

func (o *StartReplayBufferParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StartReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartReplayBufferResponse) GetStatus() string {
	return o.Status
}
func (o *StartReplayBufferResponse) GetError() string {
	return o.Error
}

/*
StopReplayBufferParams represents the params body for the "StopReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopReplayBuffer.
*/
type StopReplayBufferParams struct {
	requests.Params
}

func (o *StopReplayBufferParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StopReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StopReplayBufferResponse) GetStatus() string {
	return o.Status
}
func (o *StopReplayBufferResponse) GetError() string {
	return o.Error
}

/*
SaveReplayBufferParams represents the params body for the "SaveReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveReplayBuffer.
*/
type SaveReplayBufferParams struct {
	requests.Params
}

func (o *SaveReplayBufferParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *SaveReplayBufferResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SaveReplayBufferResponse) GetStatus() string {
	return o.Status
}
func (o *SaveReplayBufferResponse) GetError() string {
	return o.Error
}
