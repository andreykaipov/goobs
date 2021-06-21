// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMediaStateParams represents the params body for the "GetMediaState" request.
Get the current playing state of a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since 4.9.0.
*/
type GetMediaStateParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetMediaState".
func (o *GetMediaStateParams) GetSelfName() string {
	return "GetMediaState"
}

/*
GetMediaStateResponse represents the response body for the "GetMediaState" request.
Get the current playing state of a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since v4.9.0.
*/
type GetMediaStateResponse struct {
	requests.ResponseBasic

	// The media state of the provided source. States: `none`, `playing`, `opening`, `buffering`, `paused`, `stopped`,
	// `ended`, `error`, `unknown`
	MediaState string `json:"mediaState,omitempty"`
}

// GetMediaState sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMediaState(params *GetMediaStateParams) (*GetMediaStateResponse, error) {
	data := &GetMediaStateResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
