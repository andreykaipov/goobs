// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetMediaTimeParams represents the params body for the "SetMediaTime" request.
Set the timestamp of a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since 4.9.0.
*/
type SetMediaTimeParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Milliseconds to set the timestamp to.
	Timestamp int `json:"timestamp,omitempty"`
}

// GetSelfName just returns "SetMediaTime".
func (o *SetMediaTimeParams) GetSelfName() string {
	return "SetMediaTime"
}

/*
SetMediaTimeResponse represents the response body for the "SetMediaTime" request.
Set the timestamp of a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since v4.9.0.
*/
type SetMediaTimeResponse struct {
	requests.ResponseBasic
}

// SetMediaTime sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetMediaTime(params *SetMediaTimeParams) (*SetMediaTimeResponse, error) {
	data := &SetMediaTimeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
