// This file has been automatically generated. Don't edit it.

package scenecollections

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.
List available scene collections
Since 4.0.0.
*/
type ListSceneCollectionsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ListSceneCollections".
func (o *ListSceneCollectionsParams) GetSelfName() string {
	return "ListSceneCollections"
}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.
List available scene collections
Since v4.0.0.
*/
type ListSceneCollectionsResponse struct {
	requests.ResponseBasic

	// Scene collections list
	SceneCollections []typedefs.ScenesCollection `json:"scene-collections,omitempty"`
}

// ListSceneCollections sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) ListSceneCollections(paramss ...*ListSceneCollectionsParams) (*ListSceneCollectionsResponse, error) {
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
