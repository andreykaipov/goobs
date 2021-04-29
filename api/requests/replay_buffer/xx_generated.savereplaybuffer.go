// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

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
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
