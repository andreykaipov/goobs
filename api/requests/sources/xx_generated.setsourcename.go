// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceNameParams represents the params body for the "SetSourceName" request.
Sets (aka rename) the name of a source. Also works with scenes since scenes are technically sources in OBS.

Note: If the new name already exists as a source, obs-websocket will return an error.
Since 4.8.0.
*/
type SetSourceNameParams struct {
	requests.ParamsBasic

	// New source name.
	NewName string `json:"newName,omitempty"`

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceName".
func (o *SetSourceNameParams) GetSelfName() string {
	return "SetSourceName"
}

/*
SetSourceNameResponse represents the response body for the "SetSourceName" request.
Sets (aka rename) the name of a source. Also works with scenes since scenes are technically sources in OBS.

Note: If the new name already exists as a source, obs-websocket will return an error.
Since v4.8.0.
*/
type SetSourceNameResponse struct {
	requests.ResponseBasic
}

// SetSourceName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceName(params *SetSourceNameParams) (*SetSourceNameResponse, error) {
	data := &SetSourceNameResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
