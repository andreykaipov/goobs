// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetRecordingFolderParams represents the params body for the "SetRecordingFolder" request.
In the current profile, sets the recording folder of the Simple and Advanced output modes to the specified value.

Note: If `SetRecordingFolder` is called while a recording is
in progress, the change won't be applied immediately and will be
effective on the next recording.
Since 4.1.0.
*/
type SetRecordingFolderParams struct {
	requests.ParamsBasic

	// Path of the recording folder.
	RecFolder string `json:"rec-folder,omitempty"`
}

// GetSelfName just returns "SetRecordingFolder".
func (o *SetRecordingFolderParams) GetSelfName() string {
	return "SetRecordingFolder"
}

/*
SetRecordingFolderResponse represents the response body for the "SetRecordingFolder" request.
In the current profile, sets the recording folder of the Simple and Advanced output modes to the specified value.

Note: If `SetRecordingFolder` is called while a recording is
in progress, the change won't be applied immediately and will be
effective on the next recording.
Since v4.1.0.
*/
type SetRecordingFolderResponse struct {
	requests.ResponseBasic
}

// SetRecordingFolder sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetRecordingFolder(params *SetRecordingFolderParams) (*SetRecordingFolderResponse, error) {
	data := &SetRecordingFolderResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
