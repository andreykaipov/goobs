// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.
Change the active scene collection.
Since 4.0.0.
*/
type SetCurrentSceneCollectionParams struct {
	requests.ParamsBasic

	// Name of the desired scene collection.
	ScName string `json:"sc-name,omitempty"`
}

// GetSelfName just returns "SetCurrentSceneCollection".
func (o *SetCurrentSceneCollectionParams) GetSelfName() string {
	return "SetCurrentSceneCollection"
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.
Change the active scene collection.
Since v4.0.0.
*/
type SetCurrentSceneCollectionResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	data := &SetCurrentSceneCollectionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
