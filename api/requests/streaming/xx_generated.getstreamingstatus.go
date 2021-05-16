// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamingStatusParams represents the params body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusParams struct {
	requests.ParamsBasic
}

// Name just returns "GetStreamingStatus".
func (o *GetStreamingStatusParams) Name() string {
	return "GetStreamingStatus"
}

/*
GetStreamingStatusResponse represents the response body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusResponse struct {
	requests.ResponseBasic

	// Always false. Retrocompatibility with OBSRemote.
	PreviewOnly bool `json:"preview-only"`

	// Time elapsed since recording started (only present if currently recording).
	RecTimecode string `json:"rec-timecode"`

	// Current recording status.
	Recording bool `json:"recording"`

	// Time elapsed since streaming started (only present if currently streaming).
	StreamTimecode string `json:"stream-timecode"`

	// Current streaming status.
	Streaming bool `json:"streaming"`
}

// GetStreamingStatus sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStreamingStatus(
	paramss ...*GetStreamingStatusParams,
) (*GetStreamingStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamingStatusParams{{}}
	}
	params := paramss[0]
	data := &GetStreamingStatusResponse{}
	if err := c.WriteMessage(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
