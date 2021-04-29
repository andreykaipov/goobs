// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

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
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
