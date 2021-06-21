// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopMediaParams represents the params body for the "StopMedia" request.
Stop a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since 4.9.0.
*/
type StopMediaParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "StopMedia".
func (o *StopMediaParams) GetSelfName() string {
	return "StopMedia"
}

/*
StopMediaResponse represents the response body for the "StopMedia" request.
Stop a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since v4.9.0.
*/
type StopMediaResponse struct {
	requests.ResponseBasic
}

// StopMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StopMedia(params *StopMediaParams) (*StopMediaResponse, error) {
	data := &StopMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
