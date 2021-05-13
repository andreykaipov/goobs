// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetRecordingFolderParams represents the params body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderParams struct {
	requests.ParamsBasic
}

// Name just returns "GetRecordingFolder".
func (o *GetRecordingFolderParams) Name() string {
	return "GetRecordingFolder"
}

/*
GetRecordingFolderResponse represents the response body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderResponse struct {
	requests.ResponseBasic

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
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
	data := &GetRecordingFolderResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
