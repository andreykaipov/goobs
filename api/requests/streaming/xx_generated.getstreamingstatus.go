// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamingStatusParams represents the params body for the "GetStreamingStatus" request.
Get current streaming and recording status.
Since 0.3.
*/
type GetStreamingStatusParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetStreamingStatus".
func (o *GetStreamingStatusParams) GetSelfName() string {
	return "GetStreamingStatus"
}

/*
GetStreamingStatusResponse represents the response body for the "GetStreamingStatus" request.
Get current streaming and recording status.
Since v0.3.
*/
type GetStreamingStatusResponse struct {
	requests.ResponseBasic

	// Always false. Retrocompatibility with OBSRemote.
	PreviewOnly bool `json:"preview-only"`

	// Time elapsed since recording started (only present if currently recording).
	RecTimecode string `json:"rec-timecode,omitempty"`

	// Current recording status.
	Recording bool `json:"recording"`

	// If recording is paused.
	RecordingPaused bool `json:"recording-paused"`

	// Time elapsed since streaming started (only present if currently streaming).
	StreamTimecode string `json:"stream-timecode,omitempty"`

	// Current streaming status.
	Streaming bool `json:"streaming"`

	// Current virtual cam status.
	Virtualcam bool `json:"virtualcam"`

	// Time elapsed since virtual cam started (only present if virtual cam currently active).
	VirtualcamTimecode string `json:"virtualcam-timecode,omitempty"`
}

// GetStreamingStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetStreamingStatus(paramss ...*GetStreamingStatusParams) (*GetStreamingStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamingStatusParams{{}}
	}
	params := paramss[0]
	data := &GetStreamingStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
