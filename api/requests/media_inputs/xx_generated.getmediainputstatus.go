// This file has been automatically generated. Don't edit it.

package mediainputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMediaInputStatusParams represents the params body for the "GetMediaInputStatus" request.
Gets the status of a media input.

Media States:

- `OBS_MEDIA_STATE_NONE`
- `OBS_MEDIA_STATE_PLAYING`
- `OBS_MEDIA_STATE_OPENING`
- `OBS_MEDIA_STATE_BUFFERING`
- `OBS_MEDIA_STATE_PAUSED`
- `OBS_MEDIA_STATE_STOPPED`
- `OBS_MEDIA_STATE_ENDED`
- `OBS_MEDIA_STATE_ERROR`
*/
type GetMediaInputStatusParams struct {
	requests.ParamsBasic

	// Name of the media input
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetMediaInputStatus".
func (o *GetMediaInputStatusParams) GetSelfName() string {
	return "GetMediaInputStatus"
}

/*
GetMediaInputStatusResponse represents the response body for the "GetMediaInputStatus" request.
Gets the status of a media input.

Media States:

- `OBS_MEDIA_STATE_NONE`
- `OBS_MEDIA_STATE_PLAYING`
- `OBS_MEDIA_STATE_OPENING`
- `OBS_MEDIA_STATE_BUFFERING`
- `OBS_MEDIA_STATE_PAUSED`
- `OBS_MEDIA_STATE_STOPPED`
- `OBS_MEDIA_STATE_ENDED`
- `OBS_MEDIA_STATE_ERROR`
*/
type GetMediaInputStatusResponse struct {
	requests.ResponseBasic

	// Position of the cursor in milliseconds. `null` if not playing
	MediaCursor float64 `json:"mediaCursor,omitempty"`

	// Total duration of the playing media in milliseconds. `null` if not playing
	MediaDuration float64 `json:"mediaDuration,omitempty"`

	// State of the media input
	MediaState string `json:"mediaState,omitempty"`
}

// GetMediaInputStatus sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMediaInputStatus(params *GetMediaInputStatusParams) (*GetMediaInputStatusResponse, error) {
	data := &GetMediaInputStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
