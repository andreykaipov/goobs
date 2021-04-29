// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

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
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
