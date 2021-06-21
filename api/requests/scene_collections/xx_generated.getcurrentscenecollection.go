// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentSceneCollectionParams represents the params body for the "GetCurrentSceneCollection" request.
Get the name of the current scene collection.
Since 4.0.0.
*/
type GetCurrentSceneCollectionParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentSceneCollection".
func (o *GetCurrentSceneCollectionParams) GetSelfName() string {
	return "GetCurrentSceneCollection"
}

/*
GetCurrentSceneCollectionResponse represents the response body for the "GetCurrentSceneCollection" request.
Get the name of the current scene collection.
Since v4.0.0.
*/
type GetCurrentSceneCollectionResponse struct {
	requests.ResponseBasic

	// Name of the currently active scene collection.
	ScName string `json:"sc-name,omitempty"`
}

// GetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
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
