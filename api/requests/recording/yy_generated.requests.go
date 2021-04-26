// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartStopRecordingParams
func (o *StartStopRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStopRecordingParams
func (o *StartStopRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStopRecordingParams
func (o *StartStopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetError() string {
	return o.Error
}

// StartStopRecording sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopRecording(
	paramss ...*StartStopRecordingParams,
) (*StartStopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartStopRecording"
	data := &StartStopRecordingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
StartRecordingParams represents the params body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartRecordingParams
func (o *StartRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartRecordingParams
func (o *StartRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartRecordingParams
func (o *StartRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartRecordingResponse
func (o *StartRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartRecordingResponse
func (o *StartRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartRecordingResponse
func (o *StartRecordingResponse) GetError() string {
	return o.Error
}

// StartRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StartRecording(paramss ...*StartRecordingParams) (*StartRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartRecording"
	data := &StartRecordingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
StopRecordingParams represents the params body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StopRecordingParams
func (o *StopRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StopRecordingParams
func (o *StopRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StopRecordingParams
func (o *StopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopRecordingResponse represents the response body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StopRecordingResponse
func (o *StopRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StopRecordingResponse
func (o *StopRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StopRecordingResponse
func (o *StopRecordingResponse) GetError() string {
	return o.Error
}

// StopRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopRecording(paramss ...*StopRecordingParams) (*StopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StopRecording"
	data := &StopRecordingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetRecordingFolderParams represents the params body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderParams struct {
	requests.Params

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

// GetRequestType returns the RequestType of SetRecordingFolderParams
func (o *SetRecordingFolderParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetRecordingFolderParams
func (o *SetRecordingFolderParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetRecordingFolderParams
func (o *SetRecordingFolderParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetRecordingFolderResponse represents the response body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetRecordingFolderResponse
func (o *SetRecordingFolderResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetRecordingFolderResponse
func (o *SetRecordingFolderResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetRecordingFolderResponse
func (o *SetRecordingFolderResponse) GetError() string {
	return o.Error
}

// SetRecordingFolder sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetRecordingFolder(
	params *SetRecordingFolderParams,
) (*SetRecordingFolderResponse, error) {
	params.RequestType = "SetRecordingFolder"
	data := &SetRecordingFolderResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetRecordingFolderParams represents the params body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetRecordingFolderParams
func (o *GetRecordingFolderParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetRecordingFolderParams
func (o *GetRecordingFolderParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetRecordingFolderParams
func (o *GetRecordingFolderParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetRecordingFolderResponse represents the response body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderResponse struct {
	requests.Response

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

// GetMessageID returns the MessageID of GetRecordingFolderResponse
func (o *GetRecordingFolderResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetRecordingFolderResponse
func (o *GetRecordingFolderResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetRecordingFolderResponse
func (o *GetRecordingFolderResponse) GetError() string {
	return o.Error
}

// GetRecordingFolder sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetRecordingFolder(
	paramss ...*GetRecordingFolderParams,
) (*GetRecordingFolderResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordingFolderParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetRecordingFolder"
	data := &GetRecordingFolderResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
