// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTracksParams represents the params body for the "GetTracks" request.
Gets whether an audio track is active for a source.
Since 4.9.1.
*/
type GetTracksParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetTracks".
func (o *GetTracksParams) GetSelfName() string {
	return "GetTracks"
}

/*
GetTracksResponse represents the response body for the "GetTracks" request.
Gets whether an audio track is active for a source.
Since v4.9.1.
*/
type GetTracksResponse struct {
	requests.ResponseBasic

	Track1 bool `json:"track1"`

	Track2 bool `json:"track2"`

	Track3 bool `json:"track3"`

	Track4 bool `json:"track4"`

	Track5 bool `json:"track5"`

	Track6 bool `json:"track6"`
}

// GetTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetTracks(params *GetTracksParams) (*GetTracksResponse, error) {
	data := &GetTracksResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
