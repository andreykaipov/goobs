// This file has been automatically generated. Don't edit it.

package scenecollections

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ListSceneCollectionsParams represents the params body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ListSceneCollectionsParams
func (o *ListSceneCollectionsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ListSceneCollectionsResponse represents the response body for the "ListSceneCollections" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListSceneCollections.
*/
type ListSceneCollectionsResponse struct {
	requests.Response

	// Scene collections list
	SceneCollections []string `json:"scene-collections"`
}

// GetMessageID returns the MessageID of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ListSceneCollectionsResponse
func (o *ListSceneCollectionsResponse) GetError() string {
	return o.Error
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
	params.RequestType = "ListSceneCollections"
	data := &ListSceneCollectionsResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
