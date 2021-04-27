// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetBrowserSourcePropertiesParams represents the params body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of GetBrowserSourcePropertiesParams
func (o *GetBrowserSourcePropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetBrowserSourcePropertiesParams
func (o *GetBrowserSourcePropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetBrowserSourcePropertiesParams
func (o *GetBrowserSourcePropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetBrowserSourcePropertiesResponse represents the response body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesResponse struct {
	requests.Response

	// CSS to inject.
	Css string `json:"css"`

	// Framerate.
	Fps int `json:"fps"`

	// Height.
	Height int `json:"height"`

	// Indicates that a local file is in use.
	IsLocalFile bool `json:"is_local_file"`

	// file path.
	LocalFile string `json:"local_file"`

	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`

	// Source name.
	Source string `json:"source"`

	// Url.
	Url string `json:"url"`

	// Width.
	Width int `json:"width"`
}

// GetMessageID returns the MessageID of GetBrowserSourcePropertiesResponse
func (o *GetBrowserSourcePropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetBrowserSourcePropertiesResponse
func (o *GetBrowserSourcePropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetBrowserSourcePropertiesResponse
func (o *GetBrowserSourcePropertiesResponse) GetError() string {
	return o.Error
}

// GetBrowserSourceProperties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) GetBrowserSourceProperties(
	params *GetBrowserSourcePropertiesParams,
) (*GetBrowserSourcePropertiesResponse, error) {
	params.RequestType = "GetBrowserSourceProperties"
	data := &GetBrowserSourcePropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
