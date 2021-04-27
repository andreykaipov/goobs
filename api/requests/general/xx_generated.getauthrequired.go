// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetAuthRequiredParams
func (o *GetAuthRequiredParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetAuthRequiredParams
func (o *GetAuthRequiredParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetAuthRequiredParams
func (o *GetAuthRequiredParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredResponse struct {
	requests.Response

	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

// GetMessageID returns the MessageID of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetError() string {
	return o.Error
}

// GetAuthRequired sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetAuthRequired(
	paramss ...*GetAuthRequiredParams,
) (*GetAuthRequiredResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetAuthRequiredParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetAuthRequired"
	data := &GetAuthRequiredResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
