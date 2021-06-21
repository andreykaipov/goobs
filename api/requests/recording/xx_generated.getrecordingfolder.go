// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetRecordingFolderParams represents the params body for the "GetRecordingFolder" request.
Get the path of  the current recording folder.
Since 4.1.0.
*/
type GetRecordingFolderParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetRecordingFolder".
func (o *GetRecordingFolderParams) GetSelfName() string {
	return "GetRecordingFolder"
}

/*
GetRecordingFolderResponse represents the response body for the "GetRecordingFolder" request.
Get the path of  the current recording folder.
Since v4.1.0.
*/
type GetRecordingFolderResponse struct {
	requests.ResponseBasic

	// Path of the recording folder.
	RecFolder string `json:"rec-folder,omitempty"`
}

// GetRecordingFolder sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetRecordingFolder(paramss ...*GetRecordingFolderParams) (*GetRecordingFolderResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetRecordingFolderParams{{}}
	}
	params := paramss[0]
	data := &GetRecordingFolderResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
