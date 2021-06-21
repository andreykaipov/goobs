// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMediaDurationParams represents the params body for the "GetMediaDuration" request.
Get the length of media in milliseconds. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note: For some reason, for the first 5 or so seconds that the media is playing, the total duration can be off by upwards of 50ms.
Since 4.9.0.
*/
type GetMediaDurationParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetMediaDuration".
func (o *GetMediaDurationParams) GetSelfName() string {
	return "GetMediaDuration"
}

/*
GetMediaDurationResponse represents the response body for the "GetMediaDuration" request.
Get the length of media in milliseconds. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note: For some reason, for the first 5 or so seconds that the media is playing, the total duration can be off by upwards of 50ms.
Since v4.9.0.
*/
type GetMediaDurationResponse struct {
	requests.ResponseBasic

	// The total length of media in milliseconds..
	MediaDuration int `json:"mediaDuration,omitempty"`
}

// GetMediaDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMediaDuration(params *GetMediaDurationParams) (*GetMediaDurationResponse, error) {
	data := &GetMediaDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
