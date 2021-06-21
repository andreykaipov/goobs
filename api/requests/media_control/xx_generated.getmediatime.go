// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMediaTimeParams represents the params body for the "GetMediaTime" request.
Get the current timestamp of media in milliseconds. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since 4.9.0.
*/
type GetMediaTimeParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetMediaTime".
func (o *GetMediaTimeParams) GetSelfName() string {
	return "GetMediaTime"
}

/*
GetMediaTimeResponse represents the response body for the "GetMediaTime" request.
Get the current timestamp of media in milliseconds. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since v4.9.0.
*/
type GetMediaTimeResponse struct {
	requests.ResponseBasic

	// The time in milliseconds since the start of the media.
	Timestamp int `json:"timestamp,omitempty"`
}

// GetMediaTime sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMediaTime(params *GetMediaTimeParams) (*GetMediaTimeResponse, error) {
	data := &GetMediaTimeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
