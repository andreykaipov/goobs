// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

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
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
