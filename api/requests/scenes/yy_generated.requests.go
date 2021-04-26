// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneParams represents the params body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneParams struct {
	requests.Params

	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of SetCurrentSceneParams
func (o *SetCurrentSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentSceneParams
func (o *SetCurrentSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentSceneParams
func (o *SetCurrentSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentSceneResponse represents the response body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentSceneResponse
func (o *SetCurrentSceneResponse) GetError() string {
	return o.Error
}

// SetCurrentScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentScene(params *SetCurrentSceneParams) (*SetCurrentSceneResponse, error) {
	params.RequestType = "SetCurrentScene"
	data := &SetCurrentSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetCurrentSceneParams represents the params body for the "GetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentSceneParams
func (o *GetCurrentSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentSceneParams
func (o *GetCurrentSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentSceneParams
func (o *GetCurrentSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentSceneResponse represents the response body for the "GetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneResponse struct {
	requests.Response

	// Name of the currently active scene.
	Name string `json:"name"`

	// Ordered list of the current scene's source items.
	Sources []map[string]interface{} `json:"sources"`
}

// GetMessageID returns the MessageID of GetCurrentSceneResponse
func (o *GetCurrentSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentSceneResponse
func (o *GetCurrentSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentSceneResponse
func (o *GetCurrentSceneResponse) GetError() string {
	return o.Error
}

// GetCurrentScene sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentScene(
	paramss ...*GetCurrentSceneParams,
) (*GetCurrentSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentScene"
	data := &GetCurrentSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetSceneListParams represents the params body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetSceneListParams
func (o *GetSceneListParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSceneListParams
func (o *GetSceneListParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSceneListParams
func (o *GetSceneListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSceneListResponse represents the response body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListResponse struct {
	requests.Response

	// Name of the currently active scene.
	CurrentScene string `json:"current-scene"`

	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for
	// more information).
	Scenes []map[string]interface{} `json:"scenes"`
}

// GetMessageID returns the MessageID of GetSceneListResponse
func (o *GetSceneListResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSceneListResponse
func (o *GetSceneListResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSceneListResponse
func (o *GetSceneListResponse) GetError() string {
	return o.Error
}

// GetSceneList sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSceneList(paramss ...*GetSceneListParams) (*GetSceneListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneListParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetSceneList"
	data := &GetSceneListResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
ReorderSceneItemsParams represents the params body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsParams struct {
	requests.Params

	Items []struct {
		// Id of a specific scene item. Unique on a scene by scene basis.
		Id int `json:"id"`

		// Name of a scene item. Sufficiently unique if no scene items share sources within the
		// scene.
		Name string `json:"name"`
	} `json:"items"`

	// Name of the scene to reorder (defaults to current).
	Scene string `json:"scene"`
}

// GetRequestType returns the RequestType of ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ReorderSceneItemsParams
func (o *ReorderSceneItemsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ReorderSceneItemsResponse represents the response body for the "ReorderSceneItems" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ReorderSceneItemsResponse
func (o *ReorderSceneItemsResponse) GetError() string {
	return o.Error
}

// ReorderSceneItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSceneItems(
	params *ReorderSceneItemsParams,
) (*ReorderSceneItemsResponse, error) {
	params.RequestType = "ReorderSceneItems"
	data := &ReorderSceneItemsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
