// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

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
