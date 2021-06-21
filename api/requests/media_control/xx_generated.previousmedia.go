// This file has been automatically generated. Don't edit it.

package mediacontrol

import requests "github.com/andreykaipov/goobs/api/requests"

/*
PreviousMediaParams represents the params body for the "PreviousMedia" request.
Go to the previous media item in the playlist. Supports only vlc media source (as of OBS v25.0.8)
Since 4.9.0.
*/
type PreviousMediaParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "PreviousMedia".
func (o *PreviousMediaParams) GetSelfName() string {
	return "PreviousMedia"
}

/*
PreviousMediaResponse represents the response body for the "PreviousMedia" request.
Go to the previous media item in the playlist. Supports only vlc media source (as of OBS v25.0.8)
Since v4.9.0.
*/
type PreviousMediaResponse struct {
	requests.ResponseBasic
}

// PreviousMedia sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) PreviousMedia(params *PreviousMediaParams) (*PreviousMediaResponse, error) {
	data := &PreviousMediaResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
