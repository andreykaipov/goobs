// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTracksParams represents the params body for the "SetTracks" request.
Changes whether an audio track is active for a source.
Since 4.9.1.
*/
type SetTracksParams struct {
	requests.ParamsBasic

	// Whether audio track is active or not.
	Active bool `json:"active"`

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Audio tracks 1-6.
	Track int `json:"track,omitempty"`
}

// GetSelfName just returns "SetTracks".
func (o *SetTracksParams) GetSelfName() string {
	return "SetTracks"
}

/*
SetTracksResponse represents the response body for the "SetTracks" request.
Changes whether an audio track is active for a source.
Since v4.9.1.
*/
type SetTracksResponse struct {
	requests.ResponseBasic
}

// SetTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTracks(params *SetTracksParams) (*SetTracksResponse, error) {
	data := &SetTracksResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
