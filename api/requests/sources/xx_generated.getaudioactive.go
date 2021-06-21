// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAudioActiveParams represents the params body for the "GetAudioActive" request.
Get the audio's active status of a specified source.
Since 4.9.0.
*/
type GetAudioActiveParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetAudioActive".
func (o *GetAudioActiveParams) GetSelfName() string {
	return "GetAudioActive"
}

/*
GetAudioActiveResponse represents the response body for the "GetAudioActive" request.
Get the audio's active status of a specified source.
Since v4.9.0.
*/
type GetAudioActiveResponse struct {
	requests.ResponseBasic

	// Audio active status of the source.
	AudioActive bool `json:"audioActive"`
}

// GetAudioActive sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetAudioActive(params *GetAudioActiveParams) (*GetAudioActiveResponse, error) {
	data := &GetAudioActiveResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
