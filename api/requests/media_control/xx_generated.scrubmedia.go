// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ScrubMediaParams represents the params body for the "ScrubMedia" request.
Scrub media using a supplied offset. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note: Due to processing/network delays, this request is not perfect. The processing rate of this request has also not been tested.
Since 4.9.0.
*/
type ScrubMediaParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Millisecond offset (positive or negative) to offset the current media position.
	TimeOffset int `json:"timeOffset,omitempty"`
}

// GetSelfName just returns "ScrubMedia".
func (o *ScrubMediaParams) GetSelfName() string {
	return "ScrubMedia"
}

/*
ScrubMediaResponse represents the response body for the "ScrubMedia" request.
Scrub media using a supplied offset. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note: Due to processing/network delays, this request is not perfect. The processing rate of this request has also not been tested.
Since v4.9.0.
*/
type ScrubMediaResponse struct {
	requests.ResponseBasic
}

// ScrubMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ScrubMedia(params *ScrubMediaParams) (*ScrubMediaResponse, error) {
	data := &ScrubMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
