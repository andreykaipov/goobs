// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMediaSourcesListParams represents the params body for the "GetMediaSourcesList" request.
List the media state of all media sources (vlc and media source)
Since 4.9.0.
*/
type GetMediaSourcesListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetMediaSourcesList".
func (o *GetMediaSourcesListParams) GetSelfName() string {
	return "GetMediaSourcesList"
}

/*
GetMediaSourcesListResponse represents the response body for the "GetMediaSourcesList" request.
List the media state of all media sources (vlc and media source)
Since v4.9.0.
*/
type GetMediaSourcesListResponse struct {
	requests.ResponseBasic

	MediaSources []*MediaSource `json:"mediaSources,omitempty"`
}

type MediaSource struct {
	// The current state of media for that source. States: `none`, `playing`, `opening`, `buffering`, `paused`,
	// `stopped`, `ended`, `error`, `unknown`
	MediaState string `json:"mediaState,omitempty"`

	// Unique source internal type (a.k.a `ffmpeg_source` or `vlc_source`)
	SourceKind string `json:"sourceKind,omitempty"`

	// Unique source name
	SourceName string `json:"sourceName,omitempty"`
}

// GetMediaSourcesList sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetMediaSourcesList(paramss ...*GetMediaSourcesListParams) (*GetMediaSourcesListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetMediaSourcesListParams{{}}
	}
	params := paramss[0]
	data := &GetMediaSourcesListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
