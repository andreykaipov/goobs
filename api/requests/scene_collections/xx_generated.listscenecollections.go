// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.
List available scene collections

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsParams struct {
	requests.ParamsBasic
}

// Name just returns "ListSceneCollections".
func (o *ListSceneCollectionsParams) Name() string {
	return "ListSceneCollections"
}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.
List available scene collections

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsResponse struct {
	requests.ResponseBasic

	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}

// ListSceneCollections sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) ListSceneCollections(
	paramss ...*ListSceneCollectionsParams,
) (*ListSceneCollectionsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListSceneCollectionsParams{{}}
	}
	params := paramss[0]
	data := &ListSceneCollectionsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
