// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionListParams represents the params body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetTransitionListParams
func (o *GetTransitionListParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetTransitionListParams
func (o *GetTransitionListParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetTransitionListParams
func (o *GetTransitionListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTransitionListResponse represents the response body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListResponse struct {
	requests.Response

	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`

	Transitions []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}

// GetMessageID returns the MessageID of GetTransitionListResponse
func (o *GetTransitionListResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetTransitionListResponse
func (o *GetTransitionListResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetTransitionListResponse
func (o *GetTransitionListResponse) GetError() string {
	return o.Error
}

// GetTransitionList sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionList(
	paramss ...*GetTransitionListParams,
) (*GetTransitionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionListParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetTransitionList"
	data := &GetTransitionListResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
