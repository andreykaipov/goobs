// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetRecordingFolderParams represents the params body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderParams struct {
	requests.ParamsBasic

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

// Name just returns "SetRecordingFolder".
func (o *SetRecordingFolderParams) Name() string {
	return "SetRecordingFolder"
}

/*
SetRecordingFolderResponse represents the response body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderResponse struct {
	requests.ResponseBasic
}

// SetRecordingFolder sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetRecordingFolder(
	params *SetRecordingFolderParams,
) (*SetRecordingFolderResponse, error) {
	data := &SetRecordingFolderResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
