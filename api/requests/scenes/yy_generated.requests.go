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

func (o *SetCurrentSceneParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *SetCurrentSceneResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetCurrentSceneResponse) GetStatus() string {
	return o.Status
}
func (o *SetCurrentSceneResponse) GetError() string {
	return o.Error
}

/*
GetCurrentSceneParams represents the params body for the "GetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentScene.
*/
type GetCurrentSceneParams struct {
	requests.Params
}

func (o *GetCurrentSceneParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *GetCurrentSceneResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetCurrentSceneResponse) GetStatus() string {
	return o.Status
}
func (o *GetCurrentSceneResponse) GetError() string {
	return o.Error
}

/*
GetSceneListParams represents the params body for the "GetSceneList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneList.
*/
type GetSceneListParams struct {
	requests.Params
}

func (o *GetSceneListParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *GetSceneListResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSceneListResponse) GetStatus() string {
	return o.Status
}
func (o *GetSceneListResponse) GetError() string {
	return o.Error
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

func (o *ReorderSceneItemsParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *ReorderSceneItemsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ReorderSceneItemsResponse) GetStatus() string {
	return o.Status
}
func (o *ReorderSceneItemsResponse) GetError() string {
	return o.Error
}
