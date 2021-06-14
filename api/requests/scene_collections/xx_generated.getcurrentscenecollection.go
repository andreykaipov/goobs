// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentSceneCollectionParams represents the params body for the "GetCurrentSceneCollection" request.
Get the name of the current scene collection.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionParams struct {
	requests.ParamsBasic
}

// Name just returns "GetCurrentSceneCollection".
func (o *GetCurrentSceneCollectionParams) Name() string {
	return "GetCurrentSceneCollection"
}

/*
GetCurrentSceneCollectionResponse represents the response body for the "GetCurrentSceneCollection" request.
Get the name of the current scene collection.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetCurrentSceneCollection.
*/
type GetCurrentSceneCollectionResponse struct {
	requests.ResponseBasic

	// Name of the currently active scene collection.
	ScName string `json:"sc-name"`
}

// GetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentSceneCollection(
	paramss ...*GetCurrentSceneCollectionParams,
) (*GetCurrentSceneCollectionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneCollectionParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentSceneCollectionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
