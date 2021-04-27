// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetPreviewSceneParams
func (o *GetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetPreviewSceneParams
func (o *GetPreviewSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetPreviewSceneParams
func (o *GetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	requests.Response

	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []map[string]interface{} `json:"sources"`
}

// GetMessageID returns the MessageID of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetError() string {
	return o.Error
}

// GetPreviewScene sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetPreviewScene(
	paramss ...*GetPreviewSceneParams,
) (*GetPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetPreviewSceneParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetPreviewScene"
	data := &GetPreviewSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
