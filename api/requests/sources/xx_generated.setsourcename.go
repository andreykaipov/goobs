// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceNameParams represents the params body for the "SetSourceName" request.
Sets (aka rename) the name of a source. Also works with scenes since scenes are technically sources in OBS.

Note: If the new name already exists as a source, OBS will automatically modify the name to not interfere.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSourceName.
*/
type SetSourceNameParams struct {
	requests.ParamsBasic

	// New source name.
	NewName string `json:"newName"`

	// Source name.
	SourceName string `json:"sourceName"`
}

// GetSelfName just returns "SetSourceName".
func (o *SetSourceNameParams) GetSelfName() string {
	return "SetSourceName"
}

/*
SetSourceNameResponse represents the response body for the "SetSourceName" request.
Sets (aka rename) the name of a source. Also works with scenes since scenes are technically sources in OBS.

Note: If the new name already exists as a source, OBS will automatically modify the name to not interfere.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSourceName.
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
