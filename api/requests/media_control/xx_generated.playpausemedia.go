// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
PlayPauseMediaParams represents the params body for the "PlayPauseMedia" request.
Pause or play a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note :Leaving out `playPause` toggles the current pause state
Since 4.9.0.
*/
type PlayPauseMediaParams struct {
	requests.ParamsBasic

	// (optional) Whether to pause or play the source. `false` for play, `true` for pause.
	PlayPause bool `json:"playPause"`

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "PlayPauseMedia".
func (o *PlayPauseMediaParams) GetSelfName() string {
	return "PlayPauseMedia"
}

/*
PlayPauseMediaResponse represents the response body for the "PlayPauseMedia" request.
Pause or play a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Note :Leaving out `playPause` toggles the current pause state
Since v4.9.0.
*/
type PlayPauseMediaResponse struct {
	requests.ResponseBasic
}

// PlayPauseMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) PlayPauseMedia(params *PlayPauseMediaParams) (*PlayPauseMediaResponse, error) {
	data := &PlayPauseMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
