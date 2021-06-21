// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
NextMediaParams represents the params body for the "NextMedia" request.
Skip to the next media item in the playlist. Supports only vlc media source (as of OBS v25.0.8)
Since 4.9.0.
*/
type NextMediaParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "NextMedia".
func (o *NextMediaParams) GetSelfName() string {
	return "NextMedia"
}

/*
NextMediaResponse represents the response body for the "NextMedia" request.
Skip to the next media item in the playlist. Supports only vlc media source (as of OBS v25.0.8)
Since v4.9.0.
*/
type NextMediaResponse struct {
	requests.ResponseBasic
}

// NextMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) NextMedia(params *NextMediaParams) (*NextMediaResponse, error) {
	data := &NextMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
