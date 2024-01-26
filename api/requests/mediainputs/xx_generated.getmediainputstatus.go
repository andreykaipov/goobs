// This file has been automatically generated. Don't edit it.

package mediainputs

// Represents the request body for the GetMediaInputStatus request.
type GetMediaInputStatusParams struct {
	// Name of the media input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the media input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewGetMediaInputStatusParams() *GetMediaInputStatusParams {
	return &GetMediaInputStatusParams{}
}
func (o *GetMediaInputStatusParams) WithInputName(x string) *GetMediaInputStatusParams {
	o.InputName = &x
	return o
}
func (o *GetMediaInputStatusParams) WithInputUuid(x string) *GetMediaInputStatusParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *GetMediaInputStatusParams) GetRequestName() string {
	return "GetMediaInputStatus"
}

// Represents the response body for the GetMediaInputStatus request.
type GetMediaInputStatusResponse struct {
	_response

	// Position of the cursor in milliseconds. `null` if not playing
	MediaCursor float64 `json:"mediaCursor,omitempty"`

	// Total duration of the playing media in milliseconds. `null` if not playing
	MediaDuration float64 `json:"mediaDuration,omitempty"`

	// State of the media input
	MediaState string `json:"mediaState,omitempty"`
}

/*
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
func (c *Client) GetMediaInputStatus(paramss ...*GetMediaInputStatusParams) (*GetMediaInputStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetMediaInputStatusParams{{}}
	}
	params := paramss[0]
	data := &GetMediaInputStatusResponse{}
	return data, c.client.SendRequest(params, data)
}
