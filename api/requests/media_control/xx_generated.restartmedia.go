// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RestartMediaParams represents the params body for the "RestartMedia" request.
Restart a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since 4.9.0.
*/
type RestartMediaParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "RestartMedia".
func (o *RestartMediaParams) GetSelfName() string {
	return "RestartMedia"
}

/*
RestartMediaResponse represents the response body for the "RestartMedia" request.
Restart a media source. Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
Since v4.9.0.
*/
type RestartMediaResponse struct {
	requests.ResponseBasic
}

// RestartMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RestartMedia(params *RestartMediaParams) (*RestartMediaResponse, error) {
	data := &RestartMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
